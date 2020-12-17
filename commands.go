package main

import (
	"os"

	cmdCompare "github.com/jsosulska/janit/cmd/compare"
	cmdLabels "github.com/jsosulska/janit/cmd/labels"
	cmdValidate "github.com/jsosulska/janit/cmd/validate"
	"github.com/mitchellh/cli"
)

// Commands is the mapping of all available Janit commands.
var Commands map[string]cli.CommandFactory

func init() {
	ui := &cli.BasicUi{Writer: os.Stdout, ErrorWriter: os.Stderr}

	Commands = map[string]cli.CommandFactory{

		"compare": func() (cli.Command, error) {
			return &cmdCompare.Command{UI: ui}, nil
		},

		"labels": func() (cli.Command, error) {
			return &cmdLabels.Command{UI: ui}, nil
		},

		"validate": func() (cli.Command, error) {
			return &cmdValidate.Command{UI: ui}, nil
		},
	}
}

func helpFunc() cli.HelpFunc {
	// This should be updated for any commands we want to hide for any reason.
	// Hidden commands can still be executed if you know the command, but
	// aren't shown in any help output. We use this for prerelease functionality
	// or advanced features.
	hidden := map[string]struct{}{
		"inject-connect": struct{}{},
	}

	var include []string
	for k := range Commands {
		if _, ok := hidden[k]; !ok {
			include = append(include, k)
		}
	}

	return cli.FilteredHelpFunc(include, cli.BasicHelpFunc("janit"))
}
