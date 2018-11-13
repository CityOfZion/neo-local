package stack

import (
	"github.com/docker/docker/api/types/container"
	"github.com/docker/go-connections/nat"
)

// NewPrivateNet creates a new service for the cityofzion/neo-privatenet image.
func NewPrivateNet() Service {
	return Service{
		Author: "cityofzion",
		ContainerConfig: &container.Config{
			ExposedPorts: map[nat.Port]struct{}{
				"20333/tcp": {},
				"20334/tcp": {},
				"20335/tcp": {},
				"20336/tcp": {},
				"30333/tcp": {},
				"30334/tcp": {},
				"30335/tcp": {},
				"30336/tcp": {},
			},
		},
		HostConfig: &container.HostConfig{
			PortBindings: map[nat.Port][]nat.PortBinding{
				"20333/tcp": {
					{
						HostIP:   "0.0.0.0",
						HostPort: "20333",
					},
				},
				"20334/tcp": {
					{
						HostIP:   "0.0.0.0",
						HostPort: "20334",
					},
				},
				"20335/tcp": {
					{
						HostIP:   "0.0.0.0",
						HostPort: "20335",
					},
				},
				"20336/tcp": {
					{
						HostIP:   "0.0.0.0",
						HostPort: "20336",
					},
				},
				"30333/tcp": {
					{
						HostIP:   "0.0.0.0",
						HostPort: "30333",
					},
				},
				"30334/tcp": {
					{
						HostIP:   "0.0.0.0",
						HostPort: "30334",
					},
				},
				"30335/tcp": {
					{
						HostIP:   "0.0.0.0",
						HostPort: "30335",
					},
				},
				"30336/tcp": {
					{
						HostIP:   "0.0.0.0",
						HostPort: "30336",
					},
				},
			},
			Privileged:      false,
			PublishAllPorts: true,
		},
		Image: "neo-privatenet",
		Tag:   "2.8.0",
	}
}
