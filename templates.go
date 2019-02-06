package goignore

import "github.com/pkg/errors"

func IsSupportedTemplates(templates []string) error {
	if len(Config.SupportedTemplates) == 0 {
		supportedTemplates, err := GetTemplateList()
		if err != nil {
			return err
		}
		Config.SupportedTemplates = supportedTemplates
	}

	exist := true
	for _, template := range templates {
		exist = false
		for _, supportedTemplate := range Config.SupportedTemplates {
			if supportedTemplate == template {
				exist = true
				break
			}
		}
		if !exist {
			return errors.New("Template " + template + " not supported.")
		}
	}

	return nil
}

func GetCachedTemplate(template string) (string, error) {
	return "", nil
}

func CacheTemplate(template string) error {
	return nil
}

func IsCustomTemplates(template string) (bool, error) {
	if !Config.IsRead {
		return false, errors.New("Cannot load saved custom templates.")
	}
	if _, err := Config.CustomTemplates[template]; err {
		return true, nil
	}
	return false, errors.New("Custom templates not found.")
}

func GetCustomTemplate() (string, error) {
	return "", nil
}

func SaveCustomTemplate(templateName string, content *string, basedTemplates *string) error {
	return nil
}
