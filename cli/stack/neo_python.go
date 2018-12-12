package stack

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/docker/docker/api/types/container"
)

// NewNeoPython creates a new service for the cityofzion/neo-python image.
func NewNeoPython() (Service, error) {
	pwd, err := os.Getwd()
	if err != nil {
		return Service{}, err
	}

	binds := []string{
		fmt.Sprintf("%s/smart-contracts:/smart-contracts", filepath.Dir(pwd)),
		fmt.Sprintf("%s/wallets:/wallets", filepath.Dir(pwd)),
	}

	return Service{
		Author: "cityofzion",
		ContainerConfig: &container.Config{
			Tty: true,
		},
		HostConfig: &container.HostConfig{
			Binds:           binds,
			Privileged:      false,
			PublishAllPorts: false,
		},
		Image: "neo-python",
		Name:  "neo-python",
		Tag:   "v0.8.2",
	}, nil
}
