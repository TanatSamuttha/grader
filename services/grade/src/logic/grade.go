package logic

import (
	"context"
	"errors"
	"grade/models"
	"log"
	"strings"

	"github.com/moby/moby/client"
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

	for i, input:= range inputs {
		if input[len(input) - 1] != '\n' {
			input += "\n";
		}

		output := outputs[i];

		var gradeResDTO models.GradeResDTO;
		
		execOutput, execErr,  err := Execute(&input, resp, ctx);
		if err != nil {
			return errors.New("Error execute -> " + err.Error());
		}
		if len(execErr) > 0 {
			gradeRes[i] = false;
			gradeResDTO = models.GradeResDTO{
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
			gradeResDTO = models.GradeResDTO{
				JobID: job.ID,
				Task: i,
				Result: true,
				Error: "",
			}
			log.Println("correct");
		} else {
			gradeRes[i] = false;
			gradeResDTO = models.GradeResDTO{
				JobID: job.ID,
				Task: i,
				Result: false,
				Error: "",
			}
			log.Println("wrong");
		}

		GradeResBuffer <- gradeResDTO;
	}

	return nil;
}