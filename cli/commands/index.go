package commands

import (
	"github.com/urfave/cli"
)

// GenerateCommandsIndex creates a slice of all the commands that are within
// the CLI.
func GenerateCommandsIndex() []cli.Command {
	start := NewStart()
	status := NewStatus()

	return []cli.Command{
		start.ToCommand(),
		status.ToCommand(),
	}
}
