package logic

import (
	"grade/config"
	"bytes"
	"context"
	"errors"

	"github.com/moby/moby/api/pkg/stdcopy"
	"github.com/moby/moby/client"
)

func Compile(resp *client.ContainerCreateResult, ctx context.Context) (string, string, error) {
	execResp, err := config.DockerClient.ExecCreate(
		ctx,
		(*resp).ID,
		client.ExecCreateOptions{
			Cmd: []string{
				"g++",
				"/workspace/main.cpp",
				"-o",
				"/workspace/main",
			},
			AttachStdin:  false,
			AttachStdout: true,
			AttachStderr: true,
		},
	)
	if err != nil {
		return "", "", errors.New("Error create execute -> " + err.Error())
	}

	attachResp, err := config.DockerClient.ExecAttach(
		ctx,
		execResp.ID,
		client.ExecAttachOptions{},
	)
	if err != nil {
		return "", "", errors.New("Error attach execute -> " + err.Error())
	}

	defer attachResp.Close()
	stdout := new(bytes.Buffer)
	stderr := new(bytes.Buffer)

	_, err = stdcopy.StdCopy(stdout, stderr, attachResp.Reader)
	if err != nil {
		return "", "", errors.New("Error read compile stdout -> " + err.Error())
	}

	if stderr.Len() > 0 {
		return "", stderr.String(), nil
	}

	return stdout.String(), "", nil
}
