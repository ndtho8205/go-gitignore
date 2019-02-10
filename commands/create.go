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

	inputTemplates := make([]string, 0)
	for _, template := range fs.Args() {
		inputTemplates = append(inputTemplates, strings.Split(strings.TrimSpace(template), ",")...)
	}
	inputTemplates = preprocessInputTemplates(inputTemplates...)

	content, err := goignore.Config.Templates.GetTemplate(inputTemplates...)
	if err != nil {
		log.Fatal(err)
	}

	var filename string
	if goignore.IsProductionEnvironment() {
		filename = ".gitignore"
	} else {
		filename = "gitignore_dev"
	}

	if err := ioutil.WriteFile(filename, []byte(content), 0644); err != nil {
		log.Fatal(err)
	}

	if flags.customTemplateName != "" {
		err := goignore.Config.Templates.SaveCustomTemplate(flags.customTemplateName, &content, inputTemplates...)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func (flags *createCommandFlags) Usage() {
	fmt.Printf("usage: goignore %s [list of using templates]:\n", CreateCommand.Name)
	os.Exit(0)
}

// PreprocessInputTemplates preprocess the input.
func preprocessInputTemplates(inputTemplates ...string) []string {
	if len(inputTemplates) <= 1 {
		return inputTemplates
	}

	templates := &goignore.Config.Templates

	duplicate := make(map[string]bool)
	checkedInputTemplates := make([]string, 0, len(inputTemplates))

	for _, inputTemplate := range inputTemplates {
		if templates.IsCustomTemplate(inputTemplate) == nil {
			basedTemplates := strings.Split(templates.CustomTemplates[inputTemplate[1:]], ",")
			for _, basedTemplate := range basedTemplates {
				duplicate[basedTemplate] = true
			}
		} else {
			duplicate[inputTemplate] = true
		}
	}
	for k := range duplicate {
		checkedInputTemplates = append(checkedInputTemplates, k)
	}

	return checkedInputTemplates
}
