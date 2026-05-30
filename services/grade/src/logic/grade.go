package logic

import (
	"context"
	"errors"
	"grade/config"
	"grade/models"
	"log"
	"strings"
	"time"

	"github.com/gofiber/contrib/v3/websocket"
	"github.com/moby/moby/api/types/container"
	"github.com/moby/moby/client"
	"gorm.io/gorm"
)

func Grade(job models.Job, resp *client.ContainerCreateResult, ctx context.Context) error {
	log.Println(job);
	log.Println(*resp);
	err := CopyCode([]byte(job.Code), resp, ctx);
	if err != nil {
		return errors.New("Error copy code -> " + err.Error());
	}

	compileOutput, compileError, err := Compile(resp, ctx);
	log.Println(compileOutput);
	if err != nil {
		return errors.New("Error compile -> " + err.Error());
	}

	var submission models.Submission;
	var conn *websocket.Conn;

	maxRetry := 10;
	for j := 0; j < maxRetry; j++ {
		SocketMutex.RLock();
		conn = SocketMap[job.ID];
		SocketMutex.RUnlock();
					
		if conn != nil {
			break;
		}

		time.Sleep(1000 * time.Millisecond);
	}

	if len(compileError) > 0 {
		log.Println(errors.New("Compile error: -> " + compileError));
		afterGrade(
			models.GradeResJob{
				JobID: job.ID,
				Task: 0,
				Score: false,
				Compile: true,
				Error: compileError,
			},
			conn,
		)
		submission = models.Submission{
			UID: job.UID,
			Score: 0,
			Error: compileError,
		}
		
	} else {

		inputs, outputs, err := GetTestcases(job.ProblemID);
		log.Println(inputs);
		log.Println(outputs);
		if err != nil {
			return errors.New("Error get test cases -> " + err.Error());
		}

		gradeRes := make([]bool, len(inputs));
		score := 0;

		_, err = config.DockerClient.ContainerUpdate(
			ctx,
			resp.ID,
			client.ContainerUpdateOptions{
				Resources: &container.Resources{
					Memory: 6 * 1024 * 1024,
					MemorySwap: 6 * 1024 * 1024,
				},
			},
		);
		if err != nil {
			return errors.New("Error update container resources -> " + err.Error());
		}

		for i, input:= range inputs {
			if input[len(input) - 1] != '\n' {
				input += "\n";
			}

			output := outputs[i];
			
			execOutput, execErr, err := Execute(&input, resp, ctx);
			if err != nil {
				return errors.New("Error execute -> " + err.Error());
			}
			if len(execErr) > 0 {
				gradeRes[i] = false;
				afterGrade(
					models.GradeResJob{
						JobID: job.ID,
						Task: i,
						Score: false,
						Compile: true,
						Error: execErr,
					},
					conn,
				)
				log.Println("Execution error: " + execErr);
				continue;
			}

			inspect, err := config.DockerClient.ContainerInspect(
				ctx, 
				resp.ID,
				client.ContainerInspectOptions{},
			);
			if err != nil {
				return errors.New("Error container inspect -> " + err.Error());
			}
			if inspect.Container.State != nil && inspect.Container.State.OOMKilled {
				afterGrade(
					models.GradeResJob{
						JobID: job.ID,
						Task: i,
						Score: false,
						Compile: true,
						Error: "Memory limit exceed",
					},
					conn,
				)
				log.Println("Memmory limit exceed");
				continue;
			}
			
			output = strings.TrimRight(output, " \t\r\n");
			execOutput = strings.TrimRight(execOutput, " \t\r\n");
			log.Printf("input      -> %q", input);
			log.Printf("output     -> %q", output);
			log.Printf("execOutput -> %q", execOutput);

			if output == execOutput {
				gradeRes[i] = true;
				score++;
				afterGrade(
					models.GradeResJob{
						JobID: job.ID,
						Task: i,
						Score: true,
						Compile: true,
						Error: "",
					},
					conn,
				)
				log.Println("correct");
				continue;
			} else {
				gradeRes[i] = false;
				afterGrade(
					models.GradeResJob{
						JobID: job.ID,
						Task: i,
						Score: false,
						Compile: true,
						Error: "",
					},
					conn,
				)
				log.Println("wrong");
				continue;
			}
		}
		
		submission = models.Submission{
			UID: job.UID,
			Score: 100 * (score / len(inputs)),
			Error: "",
		}
	}

	err = gorm.G[models.Submission](config.DB).Create(ctx, &submission);
	if err != nil {
		return errors.New("Error create new submission history -> " + err.Error());
	}

	return nil;
}

func afterGrade(gradeResJob models.GradeResJob, conn *websocket.Conn){
	gradeResJob.Conn = conn;
	log.Println(gradeResJob);

	if gradeResJob.Conn != nil {
		GradeResBuffer <- gradeResJob;
	} else {
		log.Println("No WebSocket connection. Skip result sending");
	}
}