package logic

import (
	"context"
	"errors"
	"grade/config"
	"grade/models"
	"log"
	"strconv"
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

	problem, err := gorm.G[models.Problem](config.DB).Where("problem_id = ?", job.ProblemID).First(ctx);
	if err != nil {
		return errors.New("Error load problem meta data -> " + err.Error());
	}

	// Copy submited code to container
	err = CopyCode([]byte(job.Code), resp, ctx);
	if err != nil {
		return errors.New("Error copy code -> " + err.Error());
	}

	// Compile submited code
	compileOutput, compileError, err := Compile(resp, ctx);
	log.Println(compileOutput);
	if err != nil {
		return errors.New("Error compile -> " + err.Error());
	}

	var submission models.Submission;
	var conn *websocket.Conn;

	// Find WebSocket connection
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

	// Check compile error
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
	
	// If compilable
	} else {
		// Load testcases
		inputs, outputs, err := GetTestcases(job.ProblemID);
		log.Println(inputs);
		log.Println(outputs);
		if err != nil {
			return errors.New("Error get test cases -> " + err.Error());
		}

		gradeRes := make([]bool, len(inputs));
		score := 0;

		// Change container's resource values
		_, err = config.DockerClient.ContainerUpdate(
			ctx,
			resp.ID,
			client.ContainerUpdateOptions{
				Resources: &container.Resources{
					Memory: int64(problem.MemoryLimit) * 1024 * 1024,
					MemorySwap: int64(problem.MemoryLimit) * 1024 * 1024,
				},
			},
		);
		if err != nil {
			return errors.New("Error update container resources -> " + err.Error());
		}

		// Loop for each testcase
		for i, input:= range inputs {
			if input[len(input) - 1] != '\n' {
				input += "\n";
			}

			output := outputs[i];
			
			// Execute code
			execOutput, execErr, memory, time, err := Execute(&input, problem.TimeLimit, resp, ctx);
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

			// Check memory limit
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
			
			// Compare output and expected output
			output = strings.TrimRight(output, " \t\r\n");
			execOutput = strings.TrimRight(execOutput, " \t\r\n");
			log.Printf("input      -> %q", input);
			log.Printf("output     -> %q", output);
			log.Printf("execOutput -> %q", execOutput);
			log.Println("Memory    -> " + strconv.Itoa(memory));
			log.Println("Time      -> " + strconv.Itoa(time));
			log.Println("Time      -> " + strconv.Itoa(int(problem.TimeLimit)));

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

	// Save submission to data base
	err = gorm.G[models.Submission](config.DB).Create(ctx, &submission);
	if err != nil {
		return errors.New("Error create new submission history -> " + err.Error());
	}

	return nil;
}

// Send back result to user
func afterGrade(gradeResJob models.GradeResJob, conn *websocket.Conn){
	gradeResJob.Conn = conn;
	log.Println(gradeResJob);

	if gradeResJob.Conn != nil {
		GradeResBuffer <- gradeResJob;
	} else {
		log.Println("No WebSocket connection. Skip result sending");
	}
}