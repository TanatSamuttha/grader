package logic

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"grade/config"
	"log"
	"strconv"

	"github.com/moby/moby/api/pkg/stdcopy"
	"github.com/moby/moby/client"
)

func Execute(input *string, timeLimit uint64, resp *client.ContainerCreateResult, ctx context.Context) (string, string, int, int, error) {
	seconds := fmt.Sprintf("%.3fs", float64(timeLimit)/1000.0)
	log.Println(seconds)

	execResp, err := config.DockerClient.ExecCreate(
		ctx,
		resp.ID,
		client.ExecCreateOptions{
			Cmd: []string{
				"timeout",
				"-s", "KILL",
				seconds,
				"/workspace/main",
			},
			AttachStdin:  true,
			AttachStdout: true,
			AttachStderr: true,
		},
	)
	if err != nil {
		return "", "", 0, 0, errors.New("Error create execute -> " + err.Error())
	}

	attachResp, err := config.DockerClient.ExecAttach(
		ctx,
		execResp.ID,
		client.ExecAttachOptions{},
	)
	if err != nil {
		return "", "", 0, 0, errors.New("Error attach execute -> " + err.Error())
	}
	defer attachResp.Close()

	_, err = attachResp.Conn.Write([]byte(*input))
	if err != nil {
		return "", "", 0, 0, errors.New("Error write stdin -> " + err.Error())
	}
	attachResp.CloseWrite()

	stdout := new(bytes.Buffer)
	stderr := new(bytes.Buffer)
	_, err = stdcopy.StdCopy(stdout, stderr, attachResp.Reader)
	if err != nil {
		return "", "", 0, 0, errors.New("Error read stdout -> " + err.Error())
	}

	execInspect, err := config.DockerClient.ExecInspect(
		context.Background(),
		execResp.ID,
		client.ExecInspectOptions{},
	)
	log.Println("Exit code -> " + strconv.Itoa(execInspect.ExitCode));
	if err != nil {
		return "", "", 0, 0, errors.New("Error inspect execute -> " + err.Error())
	}

	if execInspect.ExitCode == 137 {
		return "", "Time limit exceeded", 0, 0, nil
	}

	if stderr.Len() > 0 {
		return "", stderr.String(), 0, 0, nil
	}

	return stdout.String(), "", 0, 0, nil
}