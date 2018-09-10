package main

import (
	"log"
	"os"

	"github.com/CityOfZion/neo-local/cli/commands"
	"github.com/CityOfZion/neo-local/cli/logger"
	"github.com/fatih/color"
	"github.com/urfave/cli"
)

const (
	copyright   = "MIT"
	description = "Quickly setup a local environment for NEO smart contract development"
	version     = "0.1.1"
)

var (
	author = cli.Author{
		Name: "City of Zion - https://github.com/cityofzion",
	}
	name = color.GreenString("neo-local")
)

func main() {
	logWriter := logger.NewWriter(name, version)
	log.SetFlags(0)
	log.SetOutput(logWriter)

	app := cli.NewApp()
	app.Authors = []cli.Author{author}
	app.Commands = commands.GenerateCommandsIndex()
	app.Copyright = copyright
	app.Name = name
	app.Usage = description
	app.Version = version

	err := app.Run(os.Args)
	if err != nil {
		log.Fatalf(
			"%s %s. Please check the FAQ: https://github.com/CityOfZion/neo-local/wiki/FAQ",
			color.RedString("ERROR:"),
			color.RedString(err.Error()),
		)
	}
}
