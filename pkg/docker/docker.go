package docker

import (
	"context"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

//Reference: https://docs.docker.com/engine/api/sdk/examples/

func Run() {
}

func ListContainers() {
}

func StopContainerByID(ID string) {
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}
	defer cli.Close()

	containers, err := cli.ContainerList(ctx, types.ContainerListOptions{})
	if err != nil {
		panic(err)
	}

	for _, container := range containers {
		if container.ID == ID {
			if err := cli.ContainerStop(ctx, container.ID, nil); err != nil {
				panic(err)
			}
			break
		}
	}
}
