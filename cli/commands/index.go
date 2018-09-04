package commands

import (
	"github.com/urfave/cli"
)

// GenerateCommandsIndex creates a slice of all the commands that are within
// the CLI.
func GenerateCommandsIndex() []cli.Command {
	start := NewStart()

	return []cli.Command{
		start.ToCommand(),
	}
}
