package services

import (
	"fmt"
	"io"
	"os"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"golang.org/x/net/context"
)

// CheckDockerRunning pings the Docker client to check that it is running.
func CheckDockerRunning(ctx context.Context, cli *client.Client) bool {
	response, err := cli.Ping(ctx)
	if err != nil {
		return false
	}

	return response.APIVersion != ""
}

// PullDockerImages pulls each Docker image required for the stack
// sequentially.
func PullDockerImages(ctx context.Context, cli *client.Client) error {
	for _, imageName := range dockerImageNames() {
		reader, err := cli.ImagePull(ctx, imageName, types.ImagePullOptions{})
		if err != nil {
			return fmt.Errorf("Error when pulling Docker image '%s': %s", imageName, err)
		}

		defer reader.Close()
		io.Copy(os.Stdout, reader)
	}

	return nil
}

func dockerImageNames() []string {
	return []string{
		"cityofzion/neo-local-faucet",
		"cityofzion/neo-privatenet:2.7.6",
		"cityofzion/neo-python:v0.7.7",
		"postgres:10.1",
		"registry.gitlab.com/cityofzion/neo-scan",
	}
}
