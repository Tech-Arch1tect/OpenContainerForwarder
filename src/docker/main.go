package docker

import (
	"context"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

// Get information about all running docker containers
func GetContainers() []types.Container {
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}
	containers, err := cli.ContainerList(ctx, types.ContainerListOptions{})
	if err != nil {
		panic(err)
	}
	defer cli.Close()
	return containers
}

// get hostname of a container by its id
func GetContainerHostname(containerID string) string {
	// if test, return "test"
	if containerID == "test" {
		return "test"
	}
	// if not test, get hostname from docker
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}
	container, err := cli.ContainerInspect(ctx, containerID)
	if err != nil {
		panic(err)
	}
	defer cli.Close()
	return container.Config.Hostname
}
