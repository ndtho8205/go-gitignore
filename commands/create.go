package commands

import (
	"flag"
	"fmt"
	"github.com/ndtho8205/goignore"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

// CreateCommand defines `create` subcommand
var CreateCommand = goignore.Command{
	Name:        "create",
	Description: "Create .gitignore file given template names.",
	Flags:       &createCommandFlags{},
}

type createCommandFlags struct {
	templatesFlag string
}

func (flags *createCommandFlags) Apply() *flag.FlagSet {
	return CreateCommand.NewFlags()
}

func (flags *createCommandFlags) Handle() {
	fs := flags.Apply()
	if err := fs.Parse(flag.Args()[1:]); err != nil {
		flags.Usage()
		return
	}
	templates := strings.TrimSpace(strings.Join(fs.Args(), ","))
	if templates == "" {
		log.Fatal("Template names not provided.")
		flags.Usage()
		return
	}
	content, err := goignore.GetGitignoreContent(templates)
	if err != nil {
		//TODO: Handle error
		log.Fatal(err)
	}
	if err := ioutil.WriteFile(".gitignore", []byte(content), 0644); err != nil {
		log.Fatal(err)
	}
}

func (flags *createCommandFlags) Usage() {
	fmt.Printf("usage: goignore %s [list of using templates]:\n", CreateCommand.Name)
	os.Exit(0)
}
