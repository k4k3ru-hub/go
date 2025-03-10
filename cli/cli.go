//
// cli.go
//
package cli

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)


const (
	OptConfigName	= "config"
	OptConfigAlias	= "c"
	OptConfigDesc	= "Specify the configuration file to use. Supported formats: JSON, YAML, TOML."

	OptHelpName		= "help"
	OptHelpAlias	= "h"
	OptHelpDesc		= "Display a list of available commands and global options."

	OptVersionName	= "version"
	OptVersionAlias	= "v"
	OptVersionDesc	= "Show the version of the CLI tool."
)


type Cli struct {
	Command		*Command
	Version		string
	execName	string
}
type Command struct {
	Action		CommandFunc
	Commands	[]*Command
	Name		string
	Usage		string
	Options		map[string]*Option
}
type CommandFunc func(map[string]*Option)
type Option struct {
	Alias		string
	Value		string
	Description	string
	HasValue	bool
	IsFlagSet	bool
}


//
// New CLI
//
func NewCli(defaultFunc func(map[string]*Option)) *Cli {
	// Set reserved options.
	options := make(map[string]*Option)
	options[OptHelpName] = &Option{
		Alias: OptHelpAlias,
		HasValue: false,
	}
	options[OptVersionName] = &Option{
		Alias: OptVersionAlias,
		HasValue: false,
	}

	// Create a root command.
	rootCommand := &Command{
		Action: defaultFunc,
		Name: filepath.Base(os.Args[0]),
		Options: options,
	}

	return &Cli{
		Command: rootCommand,
	}
}


//
// New Command
//
func NewCommand(name string) *Command {
	return &Command{
		Name: name,
		Options: make(map[string]*Option),
	}
}


//
// Run CLI
//
func (cli *Cli) Run() {
	args := os.Args[1:]

	// If there is no arguments provided, run the root command.
	if len(args) == 0 {
		cli.Command.Action(cli.Command.Options)
		return
	}

	// Check the help flag.
	if isHelpFlagSet(args) {
		cli.Command.showUsage()
		return
	}

	// Check the version flag.
	if isVersionFlagSet(args) {
		fmt.Printf("Version: %s\n", cli.Version)
	}
	

	// Run command.
	cli.Command.run(cli.Command.Options, args)
}


//
// Set default config option.
//
func (cmd *Command) SetDefaultConfigOption() {
	cmd.Options[OptConfigName] = &Option{
		Alias: OptConfigAlias,
		Description: OptConfigDesc,
		HasValue: true,
	}
}


//
// Set version option
//
func (cli *Cli) SetVersion(version string) {
	cli.Version = version
}


//
// Get option by the argument.
//
func getOptionByArgument(arg string, options map[string]*Option) *Option {
	if strings.HasPrefix(arg, "--") {
		optionName := strings.SplitN(arg[2:], "=", 2)[0]
		if optionName == "" {
			return nil
		}
		for name, option := range options {
			if name == optionName {
				return option
			}
		}
	} else if strings.HasPrefix(arg, "-") {
		optionName := strings.SplitN(arg[1:], "=", 2)[0]
		if optionName == "" {
			return nil
		}
		for _, option := range options {
			if option.Alias == optionName {
				return option
			}
		}
	}
	return nil
}


//
// Check if help flag (--help or -h) is set in os.Args.
//
func isHelpFlagSet(args []string) bool {
	for _, arg := range args {
		if arg == "--" + OptHelpName || arg == "-" + OptHelpAlias {
			return true
		}
	}
	return false
}


//
// Check if version flag (--version or -v) is set in os.Args.
//
func isVersionFlagSet(args []string) bool {
	for _, arg := range args {
		if arg == "--" + OptVersionName || arg == "-" + OptVersionAlias {
			return true
		}
	}
	return false
}


//
// Run command
//
func (cmd *Command) run(options map[string]*Option, args []string) {
	// Check the arguments.
	for i := 0; i < len(args); i++ {
		arg := args[i]

		if strings.HasPrefix(arg, "--") || strings.HasPrefix(arg, "-") {
			foundOption := getOptionByArgument(arg, cmd.Options)
			if foundOption == nil {
				fmt.Printf("Unknown option: %s\n\n", arg)
				cmd.showUsage()
				return
			}

			// Check if the option has a value or not.
			if foundOption.HasValue {
				// Override to the option value if the arg has `=`.
				if strings.Count(arg, "=") == 1 {
					parts := strings.SplitN(arg, "=", 2)
					if len(parts) == 2 {
						foundOption.Value = parts[1]
					}
				} else {
					if i+1 < len(args) {
						if !strings.HasPrefix(args[i+1], "-") {
							foundOption.Value = args[i+1]
						}
						i++
					}
				}
			} else {
				// Set the `IsFlagSet` flag.
				foundOption.IsFlagSet = true
			}
		} else {
			// Check if the sub command has been registered or not.
			for _, subCommand := range cmd.Commands {
				if subCommand.Name == arg {
					// Migrate options.
					migratedOptions := options
					for subCommandOptionName, subCommandOption := range subCommand.Options {
						// Append / Override an option.
						migratedOptions[subCommandOptionName] = subCommandOption
					}

					// Run recursively.
					subCommand.run(migratedOptions, args[i+1:])
					return
				}
			}

			// Unsupported sub command.
			fmt.Printf("Unknown sub command: %s\n\n", arg)
			cmd.showUsage()
			return
		}
	}

	// Run the command action.
	if cmd.Action != nil {
		cmd.Action(options)
	} else {
		cmd.showUsage()
	}
}


//
// Show usage of the command.
//
func (cmd *Command) showUsage() {
	var usage strings.Builder

	// Usage section
	usage.WriteString("Usage: " + cmd.Name)
	for optionName, option := range cmd.Options {
		if optionName != "" && option.Alias != "" {
			usage.WriteString(" [--" + optionName + "|-" + option.Alias + "]")
		} else if optionName != "" {
			usage.WriteString(" [--" + optionName + "]")
		} else if option.Alias != "" {
			usage.WriteString(" [-" + option.Alias + "]")
		} else {
			continue
		}
	}
	if len(cmd.Commands) > 0 {
		usage.WriteString(" [")
		var commandNames []string
		for _, command := range cmd.Commands {
			if command.Name == "" {
				continue
			}
			commandNames = append(commandNames, command.Name)
		}
		usage.WriteString(strings.Join(commandNames, "|") + "]")
	}


    fmt.Println(usage.String())

}
