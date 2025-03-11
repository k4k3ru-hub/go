# CLI (Command Line Interface) Utility for Go

This project provides a customizable CLI (Command Line Interface) tool written in Go. It supports multiple commands, options, and subcommands with versioning and configuration capabilities.


## Features

- Custom command execution
- Subcommand support
- Version flag (--version | -h)
- Help flag (--help | -h)


## Installation

Importing this module.
```console
import "github.com/k4k3ru-hub/go/cli"
```


## Usage

1. Run as default

There are reserved flags:
- --version | -h: Show the version of the CLI tool.
- --help | -h: Display a list of available commands and options.

When you run like:
```console
go run main.go --version
```

It would be output:
```output
Version: 1.0.0
```

2. Initialize CLI

```go
myCli := cli.NewCli(mainFunc)

func mainFunc(options map[string]*cli.Option) {
	// Here is default function.
}
```

3. Set options

```go
// Version.
myCli.SetVersion("1.0.0")

// Default config option.
myCli.Command.SetDefaultConfigOption()

// Customized option.
myCli.Command.Options["local"] = &cli.Option{
    Alias: "l",
    HasValue: false,
}
```

4. Run the CLI

```go
myCli.Run()
```


## Support me
I am a Japanese developer, and your support is a great encouragement for my work!
In addition to support, feel free to reach out with comments, feature requests, or development inquiries!

Thank you for your support😊

[![Support on Ko-fi](https://img.shields.io/badge/Ko--fi-Support%20Me-blue?style=flat-square&logo=ko-fi)](https://ko-fi.com/k4k3ru)
[![Support on Buy Me a Coffee](https://img.shields.io/badge/Buy%20Me%20a%20Coffee-Support%20Me-yellow?style=flat-square&logo=buy-me-a-coffee)](https://buymeacoffee.com/k4k3ru)


## License
This repository is open-source and distributed under the MIT License.
