package command

import (
	"github.com/codegangsta/cli"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

// CmdLogin ...
func CmdLogin(c *cli.Context) {
	file := filepath.Join(os.Getenv("HOME"), ".buccalcone")
	token := []byte(c.Args().First())

	err := ioutil.WriteFile(file, token, 0644)

	if err != nil {
		log.Fatal(err)
	}
}
