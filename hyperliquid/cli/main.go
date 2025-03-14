//
// main.go
//
package main

import (
	"fmt"

	myCliRestMetaAndAssetCtxs "github.com/k4k3ru-hub/go/hyperliquid/cli/rest/meta_and_asset_ctxs"

	"github.com/k4k3ru-hub/go/cli"
)


const (
	RestCommandName = "rest"
	RestCommandUsage = "REST API commands."
)


//
// Main.
//
func main() {
	// Initialize CLI.
	myCli := cli.NewCli(nil)
	myCli.SetVersion("1.0.0")
	myCli.Command.SetDefaultConfigOption()

	// Add `rest` command.
	restCommand := cli.NewCommand(RestCommandName)
	restCommand.Usage = RestCommandUsage
	myCli.Command.Commands = append(myCli.Command.Commands, restCommand)
	myCliRestMetaAndAssetCtxs.SetCommand(restCommand)

	// Run the CLI.
    myCli.Run()
}


//
// Run.
//
func run(options map[string]*cli.Option) {
	// [TODO] Here is not supported yet.
	fmt.Printf("Started run function.\n")
}
