package commands

import (
	"errors"

	"log"

	"github.com/CityOfZion/neo-local/cli/services"
	"github.com/CityOfZion/neo-local/cli/stack"
	"github.com/urfave/cli"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"golang.org/x/net/context"
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
		Aliases: []string{"go", "up"},
		Flags:   s.flags(),
		Name:    "start",
		Usage:   "Launch the development environment",
	}
}

func (s Start) action() func(c *cli.Context) error {
	return func(c *cli.Context) error {
		// verbose := c.Bool("v")
		// if verbose {
		// 	log.Println("Verbose logging is enabled")
		// }

		// saveState := c.Bool("ss")
		// if saveState {
		// 	log.Println("Save state is enabled, existing environment will not be destroyed")
		// } else {
		// 	log.Printf(
		// 		"Save state is %s, existing environment will be %s",
		// 		color.RedString("disabled"),
		// 		color.RedString("destroyed"),
		// 	)
		// }

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

		serv, _ := stack.Services()

		containerStarted := []string{}
		loopLimit := 0
		for len(containerStarted) < len(serv) {
			if err != nil {
				return err
			}

			if loopLimit == 30 {
				log.Printf("Infinite loop sanity triggered")
				return errors.New("Possible wrong container dependencies")
			}

		serviceIteration:
			for _, service := range serv {

				for _, c := range containerStarted {
					if c == service.Name {
						continue serviceIteration
					}
				}

				if service.DependsOn != nil {
					if sliceSubsetCheck(service.DependsOn, containerStarted) != true {
						continue
					}
				}

				resp, err := cli.ContainerCreate(
					ctx,
					service.Config(),
					service.HostConfig,
					nil,
					service.ContainerName(),
				)
				if err != nil {
					return err
				}

				err = cli.ContainerStart(ctx, resp.ID, types.ContainerStartOptions{})
				if err != nil {
					return err
				}
				log.Printf("Started %s", service.Name)

				containerStarted = append(containerStarted, service.Name)
			}

			loopLimit++
		}
		return nil
	}
}

func (s Start) flags() []cli.Flag {
	return []cli.Flag{
		// cli.BoolFlag{
		// 	Name:  "pull-images, pi",
		// 	Usage: "Pull the latest Docker images (default: true)",
		// },
		// cli.BoolFlag{
		// 	Name:  "save-state, ss",
		// 	Usage: "Any state in the existing environment will be saved (default: false)",
		// },
		// cli.BoolFlag{
		// 	Name:  "verbose, v",
		// 	Usage: "Enable verbose logging (default: false)",
		// },
	}
}

func sliceSubsetCheck(first, second []string) bool {
	set := make(map[string]int)
	for _, value := range second {
		set[value]++
	}

	for _, value := range first {
		if count, found := set[value]; !found {
			return false
		} else if count < 1 {
			return false
		} else {
			set[value] = count - 1
		}
	}

	return true
}
