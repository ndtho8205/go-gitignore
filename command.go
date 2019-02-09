package goignore

import (
	"flag"
)

// Flags defines subcommand's flags.
type Flags interface {
	Apply() *flag.FlagSet
	Handle()
	Usage()
}

// Command defines subcommand.
type Command struct {
	Name        string
	Description string
	Flags
}

// NewFlags creates subcommand flags.
func (command *Command) NewFlags() *flag.FlagSet {
	flags := flag.NewFlagSet(command.Name, flag.ExitOnError)
	flags.Usage = command.Flags.Usage
	return flags
}

// Usage prints the usage.
func (command *Command) Usage() {
	command.NewFlags().Usage()
}
