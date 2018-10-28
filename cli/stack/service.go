package stack

import (
	"fmt"

	"github.com/docker/docker/api/types/container"
)

type (
	// Service defines a Docker container to run within the stack.
	Service struct {
		Author          string
		ContainerConfig *container.Config
		HostConfig      *container.HostConfig
		Image           string
		Tag             string
	}
)

// Config returns the full launch config for creating the container.
func (s Service) Config() *container.Config {
	config := s.ContainerConfig
	config.Image = s.ImageName()
	return config
}

// ContainerName is the Docker container name.
func (s Service) ContainerName() string {
	return fmt.Sprintf("coz_neo-local_%s", s.Image)
}

// ImageName is the full Docker image name for the service, including the tag.
func (s Service) ImageName() string {
	return fmt.Sprintf("%s/%s:%s", s.Author, s.Image, s.Tag)
}
