package commands

import (
	"errors"
	"log"

	"github.com/CityOfZion/neo-local/cli/services"
	"github.com/urfave/cli"

	"github.com/docker/docker/client"
	"golang.org/x/net/context"
)

type (
	// Stop is the CLI command for stopping within the neo-local stack.
	Stop struct{}
)

// NewStop creates a new Stop.
func NewStop() Stop {
	return Stop{}
}

// ToCommand generates the CLI command struct.
func (s Stop) ToCommand() cli.Command {
	return cli.Command{
		Action:  s.action(),
		Aliases: []string{""},
		Flags:   s.flags(),
		Name:    "stop",
		Usage:   "Stop containers",
	}
}

func (s Stop) action() func(c *cli.Context) error {
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

		log.Println("Stopping containers")

		containerReferences, err := services.FetchContainerReferences(ctx, cli)
		if err != nil {
			return err
		}

		for containerName, containerID := range containerReferences {
			if err := cli.ContainerStop(ctx, containerID, nil); err != nil {
				return err
			}
			log.Printf("'%s' container stopped (%s)", containerName, containerID[:10])
		}

		return nil
	}
}

func (s Stop) flags() []cli.Flag {
	return []cli.Flag{}
}
