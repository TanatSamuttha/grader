package logic

import (
	"archive/tar"
	"bytes"
	"context"
	"errors"
	"grade/config"

	"github.com/moby/moby/client"
)

func CopyCode(code []byte, resp *client.ContainerCreateResult, ctx context.Context) error {
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

	_, err := config.DockerClient.CopyToContainer(
		ctx,
		(*resp).ID,
		client.CopyToContainerOptions{
			DestinationPath: "/workspace",
			Content: &buffer,
		},
	);
	if err != nil {
		return errors.New("Error copy code to container -> " + err.Error());
	}
	
	return nil;
}