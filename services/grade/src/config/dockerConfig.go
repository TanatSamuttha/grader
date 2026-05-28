package config

import (
	"context"
	"io"

	"github.com/moby/moby/client"
)

var DockerClient *client.Client;

func InitDockerClient() error {
	var err error;
	DockerClient, err = client.NewClientWithOpts(
		client.FromEnv,
		client.WithAPIVersionNegotiation(),
	)
	if err != nil {
		return err;
	}
	ctx := context.Background();
	reader, err := DockerClient.ImagePull(ctx, "gcc:latest", client.ImagePullOptions{})
	if err != nil {
		return err;
	}
	io.Copy(io.Discard, reader)
	reader.Close()
	return nil;
}