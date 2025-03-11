//
// main.go
//
package main

import (
	"fmt"

	"github.com/k4k3ru-hub/go/cli"
)


const (
	RestCommandName = "rest"
	RestCommandUsage = "REST API "
)


//
// Main
//
func main() {
	// Initialize CLI.
    myCli := cli.NewCli(run)
    myCli.SetVersion("1.0.0")
    myCli.Command.SetDefaultConfigOption()

	// Add `rest` command.
	restCommand := cli.NewCommand(RestCommandName)
	myCli.Command.Commands = append(myCli.Command.Commands, restCommand)

	// Run the CLI.
    myCli.Run()
}


//
// Run
//
func run(options map[string]*cli.Option) {
	fmt.Printf("Started run function.\n")
}
