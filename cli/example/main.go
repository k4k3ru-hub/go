//
// main.go
//
package main

import (
	"fmt"

	"github.com/k4k3ru-hub/go/cli"
)


func main() {
    // Initialize CLI.
    myCli := cli.NewCli(mainFunc)
    myCli.SetVersion("1.0.0")
    myCli.Command.SetDefaultConfigOption()

	// Add `list` command.
	listCommand := cli.NewCommand("list")
	listCommand.Usage = "List the configuration."
	listCommand.Action = listFunc
	listCommand.Options["local"] = &cli.Option{
        Alias: "l",
        HasValue: false,
    }
	myCli.Command.Commands = append(myCli.Command.Commands, listCommand)

	// Add `push` command.
	pushCommand := cli.NewCommand("push")
	pushCommand.Usage = "Push the source code."
	myCli.Command.Commands = append(myCli.Command.Commands, pushCommand)

	// Add `push > origin` command.
	pushOriginCommand := cli.NewCommand("origin")
	pushOriginCommand.Usage = "Push the source code to the origin."
	pushOriginCommand.Action = pushOringFunc
	pushOriginCommand.Options["url"] = &cli.Option{
		Alias: "u",
		HasValue: true,
		Value: "https://exmaple.com",
	}
	pushCommand.Commands = append(pushCommand.Commands, pushOriginCommand)

    // Run the CLI.
    myCli.Run()
}


func mainFunc(options map[string]*cli.Option) {
	for _, o := range options {
		fmt.Printf("%v\n", o)
	}
}


func listFunc(options map[string]*cli.Option) {
	fmt.Printf("Started list func.\n")
    for _, o := range options {
        fmt.Printf("%v\n", o)
    }
}


func pushOringFunc(options map[string]*cli.Option) {
    fmt.Printf("Started push origin func.\n")
    for _, o := range options {
        fmt.Printf("%v\n", o)
    }
}
