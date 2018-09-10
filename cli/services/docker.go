package services

import (
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
