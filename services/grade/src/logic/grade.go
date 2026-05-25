package logic

import (
	"context"
	"errors"
	"grade/models"
	"log"

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

	input, output, err := GetTestcases(job.ProblemID);
	log.Println(input);
	log.Println(output);

	execOutput, err := Execute(resp, ctx);
	log.Println(execOutput);

	return nil;
}