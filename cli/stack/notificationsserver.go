package stack

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/docker/docker/api/types/container"
	"github.com/docker/go-connections/nat"
)

// NewNotificationsServer creates a new service for the cityofzion/neo-python image.
func NewNotificationsServer() (Service, error) {
	pwd, err := os.Getwd()
	if err != nil {
		return Service{}, err
	}

	binds := []string{
		fmt.Sprintf("%s:/configs", filepath.Dir(pwd)),
	}

	return Service{
		Author: "cityofzion",
		ContainerConfig: &container.Config{
			Cmd: []string{
				"/bin/sh",
				"-c",
				"/bin/cp /configs/notifications-server.config.json /neo-python/custom-config.json && /usr/bin/python3 /neo-python/neo/bin/api_server.py --config /neo-python/custom-config.json --port-rest 8080",
			},
			Env: []string{
				"NOTIFICATIONS_SERVER=notifications-server",
			},
			ExposedPorts: map[nat.Port]struct{}{
				"8080/tcp": {},
			},
		},
		HostConfig: &container.HostConfig{
			Binds: binds,
			PortBindings: map[nat.Port][]nat.PortBinding{
				"8080/tcp": {
					{
						HostIP:   "0.0.0.0",
						HostPort: "8080",
					},
				},
			},
			Privileged:      false,
			PublishAllPorts: false,
		},
		Image: "neo-python",
		Name:  "notifications-server",
		Tag:   "v0.8.1",
	}, nil
}
