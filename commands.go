package main

import (
	"fmt"
	"github.com/Layzie/buccalcone/command"
	"github.com/codegangsta/cli"
	"os"
)

// GlobalFlags ...
var GlobalFlags = []cli.Flag{
	cli.StringFlag{
		EnvVar: "ENV_T",
		Name:   "team, T",
		Value:  "",
		Usage:  "Set team to post.",
	},
	cli.StringFlag{
		Name:  "category, c",
		Value: "",
		Usage: "Set category to post.",
	},
	cli.BoolFlag{
		Name:  "wip, w",
		Usage: "Set WIP setting.",
	},
	cli.StringFlag{
		Name:  "title, t",
		Value: "",
		Usage: "Set title to post.",
	},
}

// Commands ...
var Commands = []cli.Command{
	{
		Name:   "login",
		Usage:  "Login to esa.io using API Token.",
		Action: command.CmdLogin,
		Flags:  []cli.Flag{},
	},
}

// CommandNotFound ...
func CommandNotFound(c *cli.Context, command string) {
	fmt.Fprintf(os.Stderr, "%s: '%s' is not a %s command. See '%s --help'.", c.App.Name, command, c.App.Name, c.App.Name)
	os.Exit(2)
}
