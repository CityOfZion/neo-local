package services

import (
	"fmt"
	"log"

	"github.com/CityOfZion/neo-local/cli/logger"
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
	log.Println("Pulling Docker images")

	for _, imageName := range dockerImageNames() {
		prefix := fmt.Sprintf("↪  %s", imageName)
		s := logger.NewSpinner(prefix)
		s.Start()

		reader, err := cli.ImagePull(ctx, imageName, types.ImagePullOptions{})
		if err != nil {
			return fmt.Errorf("Error when pulling Docker image '%s': %s", imageName, err)
		}

		defer reader.Close()
		s.Stop()
	}

	return nil
}

func dockerImageNames() []string {
	return []string{
		"cityofzion/neo-local-faucet:latest",
		"cityofzion/neo-privatenet:2.7.6",
		"cityofzion/neo-python:v0.8.1",
		"postgres:10.5",
		"registry.gitlab.com/cityofzion/neo-scan/api:latest",
		"registry.gitlab.com/cityofzion/neo-scan/sync:latest",
	}
}
