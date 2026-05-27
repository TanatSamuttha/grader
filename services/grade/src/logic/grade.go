package logic

import (
	"context"
	"errors"
	"grade/models"
	"log"
	"strings"

	"github.com/docker/docker/api/types/container"
)

func Grade(job models.Job, resp *container.CreateResponse, ctx context.Context) error {
	log.Println(job);
	log.Println(*resp);
	err := CopyCode([]byte(job.Code), resp, ctx);
	if err != nil {
		return errors.New("Error copy code -> " + err.Error());
	}

	compileOutput, err := Compile(resp, ctx);
	log.Println(compileOutput);

	inputs, outputs, err := GetTestcases(job.ProblemID);
	log.Println(inputs);
	log.Println(outputs);
	if err != nil {
		return errors.New("Error get test cases -> " + err.Error());
	}

	for i, input:= range inputs {
		if input[len(input) - 1] != '\n' {
			input += "\n";
		}

		output := outputs[i];

		
		execOutput, err := Execute(&input, resp, ctx);
		if err != nil {
			return errors.New("Error execute -> " + err.Error());
		}

		execString := execOutput.String();
		
		output = strings.ReplaceAll(output, "\r\n", "\n");
		execString = strings.ReplaceAll(execString, "\r\n", "\n");
		
		output = strings.TrimRight(output, " \t\r\n");
		execString = strings.TrimRight(execString, " \t\r\n");
		log.Printf("input      -> %q", input);
		log.Printf("output     -> %q", output);
		log.Printf("execOutput -> %q", execString);

		if output == execString {
			log.Println("correct");
		} else {
			log.Println("wrong");
		}
	}

	return nil;
}