package commands

import (
	"errors"
	"log"

	"github.com/CityOfZion/neo-local/cli/services"
	"github.com/urfave/cli"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"golang.org/x/net/context"
)

type (
	// Down is the CLI command for stopping and removing containers within the
	// neo-local stack.
	Down struct{}
)

// NewDown creates a new Down.
func NewDown() Down {
	return Down{}
}

// ToCommand generates the CLI command struct.
func (d Down) ToCommand() cli.Command {
	return cli.Command{
		Action:  d.action(),
		Aliases: []string{"destroy"},
		Flags:   d.flags(),
		Name:    "down",
		Usage:   "Stop and destroy all containers",
	}
}

func (d Down) action() func(c *cli.Context) error {
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

		log.Println("Removing containers")

		containerReferences, err := services.FetchContainerReferences(ctx, cli)
		if err != nil {
			return err
		}

		options := types.ContainerRemoveOptions{
			Force: true,
		}

		for containerName, containerID := range containerReferences {
			if containerID == "" {
				log.Printf("'%s' container already removed", containerName)
				continue
			}

			log.Printf("'%s' container removed (%s)", containerName, containerID[:10])
			if err := cli.ContainerRemove(ctx, containerID, options); err != nil {
				return err
			}
		}

		return nil
	}
}

func (d Down) flags() []cli.Flag {
	return []cli.Flag{}
}
