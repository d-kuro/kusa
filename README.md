# kusa

🌿 kusa is a command line tool that creates a contribution on a specified date.
Can fill GitHub contributions graph with a dummy commit.
 
## What is "kusa"

https://twitter.com/tenderlove/status/907395402336538625

## Install

To install kusa, simply run:

```bash
$ go get github.com/d-kuro/kusa
```

Make sure your PATH includes the $GOPATH/bin directory so your commands can be easily used:

```bash
export PATH=$PATH:$GOPATH/bin
```

## Usage

>TODO

```text
$ kusa -h
Usage:
  kusa [flags]
  kusa [command]

Available Commands:
  create      Create GitHub contribution
  help        Help about any command

Flags:
  -h, --help   help for kusa
```

```text
$ kusa create -h
Create GitHub contribution on date specified by date option

Usage:
  kusa create [flags]

Flags:
  -c, --commit string   commit message (default ":herb: ʕ ◔ϖ◔ʔ :herb:")
  -d, --date string     date [format: yyyy-mm-dd] (default "2018-12-23") // default is today
  -h, --help            help for create
  -m, --mail string     commit author mail address (default "kusa@example.com")
  -n, --name string     commit author name (default "ʕ ◔ϖ◔ʔ")
  -r, --repo string     local directory path for clone GitHub repository (required)
```
