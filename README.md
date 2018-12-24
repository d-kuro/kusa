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

## Quick Start

Use `kusa create` command to create a GitHub contribution.

```bash
$ kusa create --repo ./kusa --mail kusa@example.com --date 2018-12-24
```

Option:

* `--repo (-r)`
  * local directory path for clone GitHub repository (required)
* `--mail (-m)`
  * Your e-mail address.
  * Notes: please use the email address registered on the [GitHub setting page](https://github.com/settings/emails).
* `--date (-d)`
  * Date to create GitHub contribution.
  * Default is Today
  * Format: yyyy-mm-dd

Please input as you are prompted to enter user name and password.
  
```bash
$ kusa create --repo ./kusa --mail kusa@example.com --date 2018-12-24
...
user name:
password:
```

When input is completed push is done and GitHub contribution is created.

### Help

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

## Todo

* [ ] Testing
* [ ] Create multiple GitHub contributions at once

## License

MIT License
