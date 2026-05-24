package logic

import (
	"archive/tar"
	"bytes"
	"context"
	"errors"
	"grade/config"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
)

func CopyCode(code []byte, resp *container.CreateResponse, ctx context.Context) error {
	var buffer bytes.Buffer;
	tw := tar.NewWriter(&buffer);

	header := &tar.Header{
		Name: "main.cpp",
		Mode: 0644,
		Size: int64(len(code)),
	}

	tw.WriteHeader(header);
	tw.Write(code);
	tw.Close();

	err := config.DockerClient.CopyToContainer(
		ctx,
		(*resp).ID,
		"/workspace",
		&buffer,
		types.CopyToContainerOptions{},
	);
	if err != nil {
		return errors.New("Error copy code to container -> " + err.Error());
	}
	
	return nil;
}