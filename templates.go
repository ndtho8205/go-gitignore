package goignore

import (
	"github.com/pkg/errors"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

// Templates defines list of suppored templates and saved custom templates.
type Templates struct {
	SupportedTemplates []string
	CustomTemplates    map[string]string
}

// IsSupportedTemplates checks if input template names is supported by gitignore.io.
func (templates *Templates) IsSupportedTemplates(inputTemplates ...string) error {
	if len(templates.SupportedTemplates) == 0 {
		supportedTemplates, err := Client.GetTemplateList()
		if err != nil {
			return err
		}
		templates.SupportedTemplates = supportedTemplates
	}

	for _, inputTemplate := range inputTemplates {
		isSupported := false
		for _, supportedTemplate := range templates.SupportedTemplates {
			if supportedTemplate == inputTemplate {
				isSupported = true
				break
			}
		}
		if !isSupported {
			return errors.New("Template " + inputTemplate + " is not supported.")
		}
	}

	return nil
}

// IsCustomTemplate checks if input template name is exist.
func (templates *Templates) IsCustomTemplate(inputTemplate string) error {
	if len(inputTemplate) > 1 && inputTemplate[0] == '@' {
		if _, exist := templates.CustomTemplates[inputTemplate[1:]]; exist {
			return nil
		}
	}
	return errors.New("custom templates not found")
}

// GetTemplate gets supported and custom templates
func (templates *Templates) GetTemplate(inputTemplates ...string) (string, error) {
	if len(inputTemplates) <= 1 && templates.IsCustomTemplate(inputTemplates[0]) == nil {
		return templates.GetCustomTemplate(inputTemplates[0])
	}

	return templates.GetSupportedTemplate(inputTemplates...)
}

// GetSupportedTemplate uses Client to get .gitignore content given input template names.
func (templates *Templates) GetSupportedTemplate(inputTemplates ...string) (string, error) {
	if err := templates.IsSupportedTemplates(inputTemplates...); err != nil {
		return "", err
	}

	content, err := Client.GetGitignoreContent(strings.Join(inputTemplates, ","))
	return content, err
}

// GetCustomTemplate get saved .gitignore content given input template name.
func (templates *Templates) GetCustomTemplate(inputTemplate string) (string, error) {
	if err := templates.IsCustomTemplate(inputTemplate); err != nil {
		return "", err
	}

	inputTemplate = inputTemplate[1:]

	customTemplateFilepath, err := getCustomTemplateFilePath(inputTemplate)
	if err != nil || !isExist(customTemplateFilepath) {
		inputTemplates := strings.Split(templates.CustomTemplates[inputTemplate], ",")
		return templates.GetSupportedTemplate(inputTemplates...)
	}
	content, err := ioutil.ReadFile(customTemplateFilepath) // nolint: gosec

	return string(content), err
}

// SaveCustomTemplate saves user custom templates.
func (templates *Templates) SaveCustomTemplate(templateName string, content *string, basedTemplates ...string) error {
	if templates.CustomTemplates == nil {
		templates.CustomTemplates = make(map[string]string)
	}
	templates.CustomTemplates[templateName] = strings.Join(basedTemplates, ",")
	templateFilepath, err := getCustomTemplateFilePath(templateName)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(templateFilepath, []byte(*content), 0644)
	return err
}

// FilterPattern searches in supported and custom templates given pattern.
func (templates *Templates) FilterPattern(pattern string) ([]string, []string) {
	if pattern == "" {
		return templates.SupportedTemplates, templates.getCustomTemplatesList()
	}

	filteredSupportedTemplates := make([]string, 0, 0)
	if pattern[0] != '@' {
		for _, template := range templates.SupportedTemplates {
			if strings.HasPrefix(template, pattern) {
				filteredSupportedTemplates = append(filteredSupportedTemplates, template)
			}
		}
	}

	filteredCustomTemplates := make([]string, 0, 0)
	if pattern[0] == '@' {
		pattern = pattern[1:]
	}
	for template := range templates.CustomTemplates {
		if strings.HasPrefix(template, pattern) {
			filteredCustomTemplates = append(filteredCustomTemplates, "@"+template)
		}
	}
	return filteredSupportedTemplates, filteredCustomTemplates
}

func (templates *Templates) getCustomTemplatesList() []string {
	customTemplates := make([]string, 0, len(templates.CustomTemplates))
	for k := range templates.CustomTemplates {
		customTemplates = append(customTemplates, k)
	}
	return customTemplates
}

// GetCustomTemplateFilePath gets the custom template file path.
func getCustomTemplateFilePath(templateFilename string) (string, error) {
	configDir, err := Config.GetConfigDir()
	if err != nil {
		return "", err
	}
	customTemplateDir := filepath.Join(configDir, CustomTemplateDirName)
	if !isExist(customTemplateDir) {
		if errMkdir := os.Mkdir(customTemplateDir, os.ModePerm); err != nil {
			return "", errMkdir
		}
	}

	templateFilepath := filepath.Clean(filepath.Join(customTemplateDir, templateFilename))

	return templateFilepath, err
}
