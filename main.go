package main

import (
	"github.com/codegangsta/cli"
	"os"
)

const (
	xclip   = "xclip -o"
	xsel    = "xsel -o"
	pbcopy  = "pbpaste"
	putclip = "getclip"
)

// ESA_TOKEN
var EsaToken string

func main() {
	app := cli.NewApp()
	app.Name = Name
	app.Version = Version
	app.Author = "Layzie <HIRAKI Satoru>"
	app.Email = "saruko313@gmail.com"
	app.Usage = "Post markdown file to esa.io"
	app.Flags = GlobalFlags
	app.Commands = Commands
	app.CommandNotFound = CommandNotFound
	app.Action = Action
	app.Run(os.Args)
}
