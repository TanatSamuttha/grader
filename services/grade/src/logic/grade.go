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
	if len(compileError) > 0 {
		return errors.New("Compile error: -> " + compileError);
	}

	inputs, outputs, err := GetTestcases(job.ProblemID);
	log.Println(inputs);
	log.Println(outputs);
	if err != nil {
		return errors.New("Error get test cases -> " + err.Error());
	}

	gradeRes := make([]bool, len(inputs));
	score := 0;
	
	var conn *websocket.Conn;

	for i, input:= range inputs {
		if input[len(input) - 1] != '\n' {
			input += "\n";
		}

		output := outputs[i];

		var gradeResJob models.GradeResJob;
		
		execOutput, execErr,  err := Execute(&input, resp, ctx);
		if err != nil {
			return errors.New("Error execute -> " + err.Error());
		}
		if len(execErr) > 0 {
			gradeRes[i] = false;
			gradeResJob = models.GradeResJob{
				JobID: job.ID,
				Task: i,
				Result: false,
				Error: execErr,
			}
			log.Println("Execution error: " + execErr);
		}
		
		output = strings.TrimRight(output, " \t\r\n");
		execOutput = strings.TrimRight(execOutput, " \t\r\n");
		log.Printf("input      -> %q", input);
		log.Printf("output     -> %q", output);
		log.Printf("execOutput -> %q", execOutput);

		if output == execOutput {
			gradeRes[i] = true;
			score++;
			gradeResJob = models.GradeResJob{
				JobID: job.ID,
				Task: i,
				Result: true,
				Error: "",
			}
			log.Println("correct");
		} else {
			gradeRes[i] = false;
			gradeResJob = models.GradeResJob{
				JobID: job.ID,
				Task: i,
				Result: false,
				Error: "",
			}
			log.Println("wrong");
		}

		
		if i == 0 {
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
		}

		gradeResJob.Conn = conn;

		log.Println(gradeResJob);

		if gradeResJob.Conn != nil {
			GradeResBuffer <- gradeResJob;
		} else {
			log.Println("No WebSocket connection. Skip result sending");
		}
	}
	
	submission := models.Submission{
		UID: job.UID,
		Score: 100 * (score / len(inputs)),
	}

	err = gorm.G[models.Submission](config.DB).Create(ctx, &submission);
	if err != nil {
		return errors.New("Error create new submission history -> " + err.Error());
	}

	return nil;
}