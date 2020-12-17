package validate

import (
	"flag"
	"fmt"
	"sync"

	"github.com/mitchellh/cli"
)

type Command struct {
	UI cli.Ui

	flags *flag.FlagSet
	//hello             string

	once sync.Once
	help string
}

func (c *Command) init() {
	// Commented out till I have flags
	c.flags = flag.NewFlagSet("", flag.ContinueOnError)
	// c.flags.StringVar(&c.flagHello, "Hello", "",
	// "Help!")

	c.help = c.Help()
}

func (c *Command) Run(args []string) int {
	c.once.Do(c.init)
	if err := c.flags.Parse(args); err != nil {
		return 1
	}

	fmt.Println("validate good to go!")
	return 0
}

func (c *Command) Help() string {
	c.once.Do(c.init)
	return c.help
}

func (c *Command) Synopsis() string { return synopsis }

const synopsis = "Any workflows that validate expected inputs."
