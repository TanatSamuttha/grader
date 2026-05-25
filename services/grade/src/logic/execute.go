package logic

import (
	"bytes"
	"context"
	"errors"
	"grade/config"
	"io"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
)

func Execute(resp *container.CreateResponse, ctx context.Context) (*bytes.Buffer, error) {
	execResp, err := config.DockerClient.ContainerExecCreate(
		ctx,
		(*resp).ID,
		types.ExecConfig{
			Cmd: []string{
				"/workspace/main",
			},
			AttachStdin: true,
			AttachStdout: true,
			AttachStderr: true,
		},
	);
	if err != nil {
		return nil, errors.New("Error create execute -> " + err.Error());
	}

	attachResp, err := config.DockerClient.ContainerExecAttach(
		ctx,
		execResp.ID,
		types.ExecStartCheck{},
	);

	defer attachResp.Close();

	attachResp.Conn.Write([]byte("5 7\n"));

	output := new(bytes.Buffer);
	io.Copy(output, attachResp.Reader);

	return output, nil;
}