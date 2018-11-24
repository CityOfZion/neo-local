package stack

import (
	"github.com/docker/docker/api/types/container"
	"github.com/docker/go-connections/nat"
)

// NewFaucet creates a new service for the cityofzion/neo-local-faucet image.
func NewFaucet() Service {
	return Service{
		Author: "cityofzion",
		ContainerConfig: &container.Config{
			Env: []string{
				"NEOSCAN=neo-scan-api:4000",
			},
			ExposedPorts: map[nat.Port]struct{}{
				"4002/tcp": {},
			},
		},
		HostConfig: &container.HostConfig{
			// Links: []string{"neo-scan-api:4000"}, TODO: Add when container exists for linking.
			PortBindings: map[nat.Port][]nat.PortBinding{
				"4002/tcp": {
					{
						HostIP:   "0.0.0.0",
						HostPort: "4002",
					},
				},
			},
			Privileged:      false,
			PublishAllPorts: true,
		},
		Image: "neo-local-faucet",
		Name:  "neo-local-faucet",
		Tag:   "latest",
	}
}
