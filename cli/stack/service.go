package stack

import (
	"fmt"

	"github.com/docker/docker/api/types/container"
)

type (
	// Service defines a Docker container to run within the stack.
	Service struct {
		ContainerConfig container.Config
		Image           string
		Tag             string
	}
)

// Config returns the full launch config for creating the container.
func (s Service) Config() *container.Config {
	config := s.ContainerConfig
	config.Image = s.ImageName()
	return &config
}

// ContainerName is the Docker container name.
func (s Service) ContainerName() string {
	return fmt.Sprintf("coz_neo-local_%s", s.Image)
}

// ImageName is the full Docker image name for the service, including the tag.
func (s Service) ImageName() string {
	return fmt.Sprintf("%s:%s", s.Image, s.Tag)
}

// Services returns all the services within the Docker stack.
func Services() []Service {
	return []Service{
		NewPostgres(),
	}
}
