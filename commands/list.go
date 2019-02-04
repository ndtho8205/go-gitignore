package commands

import (
	"flag"
	"fmt"
	"github.com/ndtho8205/goignore"
	"log"
	"os"
)

// ListCommand defines `list` subcommand
var ListCommand = goignore.Command{
	Name:        "list",
	Description: "List supported templates",
	Flags:       &listCommandFlags{},
}

type listCommandFlags struct {
	listAll       bool
	listSupported bool
	listSaved     bool
	pattern       string
}

func (flags *listCommandFlags) Apply() *flag.FlagSet {
	fs := ListCommand.NewFlags()
	fs.BoolVar(&flags.listAll, "all", true, "List all supported and saved templates.")
	fs.BoolVar(&flags.listSupported, "t", false, "List supported templates only.")
	fs.BoolVar(&flags.listSaved, "s", false, "List saved templates.")
	return fs
}

func (flags *listCommandFlags) Handle() {
	fs := flags.Apply()
	if err := fs.Parse(flag.Args()[1:]); err != nil {
		flags.Usage()
		return
	}
	templateList, err := goignore.GetTemplateList()
	if err != nil {
		//TODO: Handle error
		log.Fatal(err)
	}
	for _, name := range templateList {
		fmt.Println(name)
	}
}

func (flags *listCommandFlags) Usage() {
	fmt.Printf("usage: goignore %s [%s flags] [patterns]:\n", ListCommand.Name, ListCommand.Name)
	flags.Apply().PrintDefaults()
	os.Exit(0)
}
