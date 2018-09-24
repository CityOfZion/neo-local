package commands

import (
	"errors"

	"log"

	"github.com/CityOfZion/neo-local/cli/services"
	"github.com/fatih/color"
	"github.com/urfave/cli"

	"github.com/docker/docker/client"
	"golang.org/x/net/context"
)

const (
	name  = "start"
	usage = "Launch the development environment"
)

var (
	alias = string(name[0])
)

type (
	// Start is the CLI command for launching the development environment.
	Start struct{}
)

// NewStart creates a new Start.
func NewStart() Start {
	return Start{}
}

// ToCommand generates the CLI command struct.
func (s Start) ToCommand() cli.Command {
	return cli.Command{
		Action:  s.action(),
		Aliases: []string{alias},
		Flags:   s.flags(),
		Name:    name,
		Usage:   usage,
	}
}

func (s Start) action() func(c *cli.Context) error {
	return func(c *cli.Context) error {
		verbose := c.Bool("v")
		if verbose {
			log.Println("Verbose logging is enabled")
		}

		saveState := c.Bool("ss")
		if saveState {
			log.Println("Save state is enabled, existing environment will not be destroyed")
		} else {
			log.Printf(
				"Save state is %s, existing environment will be %s",
				color.RedString("disabled"),
				color.RedString("destroyed"),
			)
		}

		ctx := context.Background()
		cli, err := client.NewEnvClient()
		if err != nil {
			return errors.New("Unable to create Docker client")
		}

		ok := services.CheckDockerRunning(ctx, cli)
		if !ok {
			return errors.New("Docker is not running")
		}

		err = services.PullDockerImages(ctx, cli)
		if err != nil {
			return err
		}

		// TODO

		return nil
	}
}

func (s Start) flags() []cli.Flag {
	return []cli.Flag{
		cli.BoolFlag{
			Name:  "save-state, ss",
			Usage: "Any state in the existing environment will be saved (default: false)",
		},
		cli.BoolFlag{
			Name:  "verbose, v",
			Usage: "Enable verbose logging (default: false)",
		},
	}
}
