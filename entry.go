package main

import (
	"fmt"

	"github.com/codegangsta/cli"
)

func goof(c *cli.Context) {
	fmt.Println("hello")
}

var goofSubcommands = []cli.Command{
	make,
	delete,
}

var (
	make = cli.Command{
		Name:    "make",
		Aliases: []string{"m"},
		Action:  makeAction,
		Flags: []cli.Flag{
			cli.BoolFlag{
				Name:  "gae",
				Usage: "enable GAE workspace generation",
			},
		},
	}
	delete = cli.Command{
		Name:    "delete",
		Aliases: []string{"d"},
		Action:  deleteAction,
	}
)

func makeAction(c *cli.Context) {
	fmt.Println("hoge", c.Args(), c.FlagNames())
}

func deleteAction(c *cli.Context) {
	fmt.Println("foo", c.Args())
}
