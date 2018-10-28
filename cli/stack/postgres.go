package stack

import (
	"github.com/docker/docker/api/types/container"
	"github.com/docker/go-connections/nat"
)

// NewPostgres creates a new service for the Postgres container.
func NewPostgres() Service {
	return Service{
		ContainerConfig: container.Config{
			Env: []string{
				"POSTGRES_DB=neoscan_prodv",
				"POSTGRES_PASSWORD=postgres",
				"POSTGRES_USER=postgres",
			},
			ExposedPorts: map[nat.Port]struct{}{"5432/tcp": {}},
		},
		Image: "library/postgres",
		Tag:   "10.5",
	}
}
