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
	listSupported bool
	listCustom    bool
	pattern       string
}

func (flags *listCommandFlags) Apply() *flag.FlagSet {
	fs := ListCommand.NewFlags()
	fs.BoolVar(&flags.listSupported, "supported", false, "List supported templates only.")
	fs.BoolVar(&flags.listCustom, "custom", false, "List user custom templates.")
	return fs
}

func (flags *listCommandFlags) Handle() {
	fs := flags.Apply()
	if err := fs.Parse(flag.Args()[1:]); err != nil {
		flags.Usage()
		return
	}

	if len(fs.Args()) > 0 {
		flags.pattern = fs.Args()[0]
	}

	if !flags.listSupported && !flags.listCustom {
		listSupportedTemplates(flags.pattern)
		listCustomTemplates(flags.pattern)
		return
	}

	if flags.listSupported {
		listSupportedTemplates(flags.pattern)
	}

	if flags.listCustom {
		listCustomTemplates(flags.pattern)
	}

}

func (flags *listCommandFlags) Usage() {
	fmt.Printf("usage: goignore %s [%s flags] [patterns]:\n", ListCommand.Name, ListCommand.Name)
	flags.Apply().PrintDefaults()
	os.Exit(0)
}

func listSupportedTemplates(pattern string) {
	if len(goignore.Config.Templates.SupportedTemplates) == 0 {
		supportedTemplates, err := goignore.Client.GetTemplateList()
		if err != nil {
			log.Fatal(err)
		}
		goignore.Config.Templates.SupportedTemplates = supportedTemplates
	}

	outputSupportedTemplates, _ := goignore.Config.Templates.FilterPattern(pattern)

	fmt.Println("Supported templates by gitignore.io:")
	fmt.Println(formatColumn(outputSupportedTemplates...))
}

func listCustomTemplates(pattern string) {
	_, customTemplates := goignore.Config.Templates.FilterPattern(pattern)

	fmt.Println("Custom templates:")
	fmt.Println(formatColumn(customTemplates...))
}

func formatColumn(list ...string) string {
	var nList = len(list)
	var output = "  "

	if nList == 0 {
		return output + "<empty>"
	}

	for i := 0; i < nList; i++ {
		output += fmt.Sprintf("%-28s", list[i])
		if (i+1)%3 == 0 {
			output += "\n  "
		}
	}

	return output
}
