package main

import (
	"github.com/codegangsta/cli"
	"github.com/upamune/go-esa/esa"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

// Action ...
var Action = func(c *cli.Context) {
	if !c.Args().Present() {
		log.Fatal("Pass the file to post.")
	}

	filepath := c.Args().First()

	readToken()

	createPost(
		readFile(filepath),
		c.String("team"),
		c.String("title"),
		c.String("category"),
		c.Bool("wip"),
	)
}

// readToken read ESA_TOKEN from file
func readToken() string {
	file := filepath.Join(os.Getenv("HOME"), ".buccalcone")
	token, err := ioutil.ReadFile(file)

	if err != nil {
		log.Fatal("No esa token, Use `buccalcone login ESA_TOKEN`")
	}

	EsaToken = strings.TrimSpace(string(token))

	return EsaToken
}

// readFile read markdown file
func readFile(filepath string) string {
	file, err := ioutil.ReadFile(filepath)

	if err != nil {
		log.Fatal(err)
	}

	return string(file)
}

// Client make esa client
func client() *esa.Client {
	client := esa.NewClient(EsaToken)

	return client
}

// createPost post markdown to esa
func createPost(file string, team string, title string, category string, wip bool) *esa.PostResponse {
	post := esa.Post{
		BodyMd:   file,
		Category: category,
		Name:     title,
		Wip:      wip,
	}

	res, err := client().Post.Create(team, post)

	if err != nil {
		log.Fatal(err)
	}

	return res
}
