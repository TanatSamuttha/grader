package logic

import (
	"bytes"
	"context"
	"errors"
	"grade/config"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/pkg/stdcopy"
)

func Compile(resp *container.CreateResponse, ctx context.Context) (string, string, error) {
	execResp, err := config.DockerClient.ContainerExecCreate(
		ctx,
		(*resp).ID,
		types.ExecConfig{
			Cmd: []string{
				"g++",
				"/workspace/main.cpp",
				"-o",
				"/workspace/main",
			},
			AttachStdin: false,
			AttachStdout: true,
			AttachStderr: true,
		},
	);
	if err != nil {
		return "", "", errors.New("Error create execute -> " + err.Error());
	}

	attachResp, err := config.DockerClient.ContainerExecAttach(
		ctx,
		execResp.ID,
		types.ExecStartCheck{},
	);
	if err != nil {
		return "", "", errors.New("Error attach execute -> " + err.Error());
	}

	defer attachResp.Close();
	stdout := new(bytes.Buffer);
	stderr := new(bytes.Buffer);

	_, err = stdcopy.StdCopy(stdout, stderr, attachResp.Reader)
	if err != nil {
		return "", "", errors.New("Error read compile stdout -> " + err.Error())
	}

	if stderr.Len() > 0 {
		return "", stderr.String(), nil;
	}

	return stdout.String(), "", nil;
}