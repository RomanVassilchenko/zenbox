package internal

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"time"
	"zenbox/pkg/docker"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

func getenv(key string, fallback int) int64 {
	value := os.Getenv(key)
	if len(value) == 0 {
		return int64(fallback)
	}
	res, err := strconv.Atoi(value)
	if err != nil {
		res = fallback
	}
	return int64(res)
}

func Cleaner() {
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}
	defer cli.Close()

	TIMEOUT := getenv("TIMEOUT", 3600)

	for true {
		containers, err := cli.ContainerList(ctx, types.ContainerListOptions{})
		// TODO: Add
		// label='zenbox.created=true'
		if err != nil {
			panic(err)
		}

		now := time.Now().Unix()
		for _, container := range containers {
			created := container.Created
			seconds := created - now
			if seconds > TIMEOUT {
				fmt.Println("Deleted ", container.ID, " since it's too old. (", seconds, ", ", container.Status, ")")
				docker.StopContainerByID(container.ID)
			}

		}
		time.Sleep(5 * time.Second)
	}
}
