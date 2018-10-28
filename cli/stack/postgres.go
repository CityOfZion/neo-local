package stack

import (
	"github.com/docker/docker/api/types/container"
	"github.com/docker/go-connections/nat"
)

// NewPostgres creates a new service for the library/postgres container.
func NewPostgres() Service {
	return Service{
		Author: "library",
		ContainerConfig: &container.Config{
			Env: []string{
				"POSTGRES_DB=neoscan_prodv",
				"POSTGRES_PASSWORD=postgres",
				"POSTGRES_USER=postgres",
			},
			ExposedPorts: map[nat.Port]struct{}{"5432/tcp": {}},
		},
		HostConfig: nil,
		Image:      "postgres",
		Tag:        "10.5",
	}
}
