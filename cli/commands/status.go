package commands

import (
	"errors"
	"fmt"
	"log"
	"strings"

	"github.com/CityOfZion/neo-local/cli/services"
	"github.com/CityOfZion/neo-local/cli/stack"
	"github.com/urfave/cli"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"golang.org/x/net/context"
)

type (
	// Status is the CLI command for checking the status of containers in the
	// development environment.
	Status struct{}
)

// NewStatus creates a new Status.
func NewStatus() Status {
	return Status{}
}

// ToCommand generates the CLI command struct.
func (s Status) ToCommand() cli.Command {
	return cli.Command{
		Action:  s.action(),
		Aliases: []string{"ps"},
		Flags:   s.flags(),
		Name:    "status",
		Usage:   "Output overall health of network",
	}
}

func (s Status) action() func(c *cli.Context) error {
	return func(c *cli.Context) error {
		ctx := context.Background()
		cli, err := client.NewEnvClient()
		if err != nil {
			return errors.New("Unable to create Docker client")
		}

		ok := services.CheckDockerRunning(ctx, cli)
		if !ok {
			return errors.New("Docker is not running")
		}

		services := stack.Services()
		runningContainers, err := cli.ContainerList(ctx, types.ContainerListOptions{})
		if err != nil {
			return err
		}

		for _, container := range runningContainers {
			containerName := ""
			for _, name := range container.Names {
				if strings.Contains(name, "coz_neo-local_") { // TODO: this string should be a const.
					fmt.Println(name)
					containerName = name
					break
				}
			}

			if containerName == "" {
				continue
			}

			for _, service := range services {
				if strings.Contains(containerName, service.ContainerName()) {
					log.Printf(
						"%s in '%s' state", service.Image, container.State,
					)
					break
				}
			}
		}

		return nil
	}
}

func (s Status) flags() []cli.Flag {
	return []cli.Flag{}
}
