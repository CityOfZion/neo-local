package services

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"strings"

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

		var b bytes.Buffer
		_, err = io.Copy(&b, reader)
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

// FetchContainerReferences finds the container ID for each service within the
// stack.
func FetchContainerReferences(ctx context.Context, cli *client.Client) (map[string]string, error) {
	containerReferences := map[string]string{}
	serviceContainerNames := stack.ServiceContainerNames()

	for _, serviceContainerName := range serviceContainerNames {
		containerReferences[serviceContainerName] = ""
	}

	containers, err := cli.ContainerList(
		ctx,
		types.ContainerListOptions{
			All: true,
		},
	)
	if err != nil {
		return nil, err
	}

	for _, container := range containers {
		var containerName string
		for _, name := range container.Names {
			name = strings.Replace(name, "/", "", -1)
			if strings.HasPrefix(name, stack.ContainerNamePrefix) {
				containerName = name
				break
			}
		}

		if containerName == "" {
			continue
		}

		for _, serviceContainerName := range serviceContainerNames {
			if containerName == serviceContainerName {
				containerReferences[containerName] = container.ID
				break
			}
		}
	}

	return containerReferences, nil
}

