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
	templatesFlag      string
	customTemplateName string
}

func (flags *createCommandFlags) Apply() *flag.FlagSet {
	fs := CreateCommand.NewFlags()
	fs.StringVar(&flags.customTemplateName, "save", "", "Save this as a template for future.")
	return fs
}

func (flags *createCommandFlags) Handle() {
	fs := flags.Apply()
	if err := fs.Parse(flag.Args()[1:]); err != nil {
		flags.Usage()
		return
	}
	if len(fs.Args()) == 0 {
		log.Fatal("Template names not provided.")
		flags.Usage()
		return
	}
	templates := make([]string, 0)
	for _, template := range fs.Args() {
		templates = append(templates, strings.Split(strings.TrimSpace(template), ",")...)
	}

	if err := goignore.IsSupportedTemplates(templates); err != nil {
		log.Fatal(err)
	}

	content, err := goignore.GetGitignoreContent(strings.Join(templates, ","))

	if err != nil {
		//TODO: Handle error
		log.Fatal(err)
	}
	if err := ioutil.WriteFile("gitignore", []byte(content), 0644); err != nil {
		log.Fatal(err)
	}
}

func (flags *createCommandFlags) Usage() {
	fmt.Printf("usage: goignore %s [list of using templates]:\n", CreateCommand.Name)
	os.Exit(0)
}

func isSupportedTemplates(templates []string) {

}
