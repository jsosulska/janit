package main

import (
	"log"
	"os"

	"github.com/mitchellh/cli"
)

// Version constant
const version = "0.0.1"

// Main function

func main() {
	c := cli.NewCLI("janit", version)
	c.Args = os.Args[1:]
	c.Commands = Commands
	c.HelpFunc = helpFunc()

	// error handling
	exitStatus, err := c.Run()
	if err != nil {
		log.Println(err)
	}
	// exit with the status.
	os.Exit(exitStatus)
}
