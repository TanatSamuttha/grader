package config

import "github.com/docker/docker/client"

var DockerClient *client.Client;

func InitDockerClient() error {
	var err error;
	DockerClient, err = client.NewClientWithOpts(
		client.FromEnv,
		client.WithAPIVersionNegotiation(),
	)
	return err;
}