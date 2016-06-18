# buccalcone

## Description

`buccalcone` is the command for post markdown file to [esa.io](https://esa.io/).

## Usage

First, you need to login esa.io using [Personal access token](https://medley.esa.io/user/applications).

```bash
$ buccalcone login ESA_TOKEN
```

Then you'll post local markdown file to your esa team. For example,

```bash
buccalcone -T "YOUR"_TEAM -c "CATEGORY/TO/POST" -t "TITLE" -w README.md
```

Help is below.

```bash
NAME:
   buccalcone - Post markdown file to esa.io

USAGE:
   buccalcone [global options] command [command options] [arguments...]

VERSION:
   0.1.0

AUTHOR(S):
   Layzie <HIRAKI Satoru> <saruko313@gmail.com>

COMMANDS:
   login        Login to esa.io using API Token.
   help, h      Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --team, -T           Set team to post. [$ENV_T]
   --category, -c       Set category to post.
   --wip, -w            Set WIP setting.
   --title, -t          Set title to post.
   --help, -h           show help
   --version, -v        print the version
```

## Install

To install, use `go get`:

```bash
$ go get -d github.com/Layzie/buccalcone
```

or

Download the binary file from this [link](https://github.com/Layzie/buccalcone/releases).
Copy to the binary file in your `$PATH`.

## Contribution

1. Fork ([https://github.com/Layzie/buccalcone/fork](https://github.com/Layzie/buccalcone/fork))
1. Create a feature branch
1. Commit your changes
1. Rebase your local changes against the master branch
1. Run test suite with the `go test ./...` command and confirm that it passes
1. Run `gofmt -s`
1. Create a new Pull Request

## Author

[Layzie](https://github.com/Layzie)
