package main

import (
	"os"

	"github.com/codegangsta/cli"
)

const (
	appName    = "goof"
	appVersion = "0.1.0"
)

func main() {
	app := cli.NewApp()
	app.Name = appName
	app.Usage = "handy army knife tool for goofy gophers"
	app.Action = goof
	app.Commands = goofSubcommands
	app.Run(os.Args)
}
