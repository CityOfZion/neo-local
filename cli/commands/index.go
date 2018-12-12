package commands

import (
	"github.com/urfave/cli"
)

// GenerateCommandsIndex creates a slice of all the commands that are within
// the CLI.
func GenerateCommandsIndex() []cli.Command {
	down := NewDown()
	start := NewStart()
	status := NewStatus()
	stop := NewStop()

	return []cli.Command{
		down.ToCommand(),
		start.ToCommand(),
		status.ToCommand(),
		stop.ToCommand(),
	}
}
