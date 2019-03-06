package stack

import (
	"github.com/docker/docker/api/types/container"
)

// NewAutoheal creates a new service for the autoheal container.
func NewAutoheal() Service {
	return Service{
		Author: "library",
		ContainerConfig: &container.Config{
			Env: []string{
				"AUTOHEAL_CONTAINER_LABEL=autoheal",
				"AUTOHEAL_INTERVAL=5",
				"DOCKER_SOCK=/var/run/docker.sock",
			},
		},
		HostConfig: &container.HostConfig{
			Binds: []string{
				"/var/run/docker.sock:/var/run/docker.sock",
			},
			RestartPolicy: container.RestartPolicy{
				Name:              "always",
				MaximumRetryCount: 3,
			},
		},
		// TODO: Add "DependsOn"
		Image: "willfarrell/autoheal",
		Name:  "autoheal",
		Tag:   "10.5",
	}
}
