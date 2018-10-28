package services

import (
	"fmt"
	"log"

	"github.com/CityOfZion/neo-local/cli/logger"
	"github.com/CityOfZion/neo-local/cli/stack"
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

	for _, service := range stack.Services() {
		prefix := fmt.Sprintf("↪  %s", service.ImageName())
		s := logger.NewSpinner(prefix)
		s.Start()

		reader, err := cli.ImagePull(
			ctx, service.ImageName(), types.ImagePullOptions{},
		)
		if err != nil {
			return fmt.Errorf(
				"Error when pulling Docker image '%s': %s",
				service.ImageName(),
				err,
			)
		}

		defer reader.Close()
		s.Stop()
	}

	return nil
}