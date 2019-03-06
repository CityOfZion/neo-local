package stack

import (
	"fmt"

	"github.com/docker/docker/api/types/container"
)

const (
	// ContainerNamePrefix is the prefix added to all neo-local containers.
	ContainerNamePrefix = "coz_neo-local_"
)

type (
	// Service defines a Docker container to run within the stack.
	Service struct {
		Author          string
		ContainerConfig *container.Config
		DependsOn       []string
		HostConfig      *container.HostConfig
		Image           string
		Name            string
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
	return fmt.Sprintf("%s%s", ContainerNamePrefix, s.Name)
}

// ImageName is the full Docker image name for the service, including the tag.
func (s Service) ImageName() string {
	return fmt.Sprintf("%s/%s:%s", s.Author, s.Image, s.Tag)
}
