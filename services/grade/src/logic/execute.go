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

func Execute(input *string, resp *container.CreateResponse, ctx context.Context) (string, string, error) {
	execResp, err := config.DockerClient.ContainerExecCreate(
		ctx,
		resp.ID,
		types.ExecConfig{
			Cmd: []string{
				"/workspace/main",
			},
			AttachStdin:  true,
			AttachStdout: true,
			AttachStderr: true,
		},
	)
	if err != nil {
		return "", "", errors.New("Error create execute -> " + err.Error())
	}

	attachResp, err := config.DockerClient.ContainerExecAttach(
		ctx,
		execResp.ID,
		types.ExecStartCheck{},
	)
	if err != nil {
		return "", "", errors.New("Error attach execute -> " + err.Error())
	}

	defer attachResp.Close()

	_, err = attachResp.Conn.Write([]byte(*input))
	if err != nil {
		return "", "", errors.New("Error write stdin -> " + err.Error())
	}

	// Important: close stdin so program receives EOF
	attachResp.CloseWrite()

	stdout := new(bytes.Buffer)
	stderr := new(bytes.Buffer)

	_, err = stdcopy.StdCopy(stdout, stderr, attachResp.Reader)
	if err != nil {
		return "", "", errors.New("Error read stdout -> " + err.Error())
	}

	if stderr.Len() > 0 {
		return "", stderr.String(), nil;
	}

	return stdout.String(), "", nil;
}