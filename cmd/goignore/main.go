package main

import (
	"flag"
	"fmt"
	"github.com/ndtho8205/goignore"
	"github.com/ndtho8205/goignore/commands"
	"os"
	"sort"
)

var cmds = map[string]*goignore.Command{
	"create": &commands.CreateCommand,
	"list":   &commands.ListCommand,
}

func main() {
	flag.Usage = printUsage
	flag.Parse()
	if flag.NArg() < 1 {
		//TODO: goignore CLI
		printUsage()
		os.Exit(1)
	}

	if cmd, ok := cmds[os.Args[1]]; ok {
		cmd.Handle()
	} else {
		fmt.Printf("go %s: unknown command\n", os.Args[1])
		fmt.Printf("Run 'goignore --help' for usage.\n")
		os.Exit(1)
	}
}

func printUsage() {
	//FIXME: print usage
	fmt.Printf("goignore is a tool for generating a .gitignore file.\n\n")
	fmt.Printf("Usage: \n\n\tgoignore <command> [<arguments>]\n\n")
	fmt.Printf("The commands are:\n\n")

	names := make([]string, 0, len(cmds))
	for name := range cmds {
		if name != "help" {
			names = append(names, name)
		}
	}
	sort.Strings(names)
	for _, name := range names {
		fmt.Printf("\t%-8s %s\n", name, cmds[name].Description)
	}

	fmt.Printf("\nUse 'goignore <command> --help' for more information about a command.\n")
	os.Exit(0)
}
