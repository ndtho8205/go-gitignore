package goignore

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
)

const (
	// ConfigDirName is the default config directory.
	ConfigDirName = ".goignore"

	// ConfigFileName is the default config filename.
	ConfigFileName = "config.json"

	// CustomTemplateDirName is the default directory name where custom template will be cached.
	CustomTemplateDirName = "customs"
)

// Configuration defines app configuration.
type Configuration struct {
	LastUpdated string
	Templates   Templates
}

// Config is the global config variable.
var Config = Configuration{}

// Read reads latest app configuration saved.
func (config *Configuration) Read() error {
	configFilepath, err := config.GetConfigFilePath()
	if err != nil {
		return err
	}

	if isExist(configFilepath) {
		configContent, _ := ioutil.ReadFile(configFilepath) // nolint: gosec
		if err := json.Unmarshal(configContent, config); err != nil {
			return err
		}
	}

	return nil
}

// Save saves current configuration to file.
func (config *Configuration) Save() error {
	configFilepath, err := config.GetConfigFilePath()
	if err != nil {
		return err
	}

	configNewContent, err := json.Marshal(config)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(configFilepath, configNewContent, 0644)

	return err
}

// GetConfigFilePath gets the default config file path.
func (config *Configuration) GetConfigFilePath() (string, error) {
	configDir, err := config.GetConfigDir()
	if err != nil {
		return "", err
	}

	configFilePath := filepath.Join(configDir, ConfigFileName)

	return filepath.Clean(configFilePath), err
}

// GetConfigDir gets the default config directory.
func (config *Configuration) GetConfigDir() (string, error) {
	var configDir string
	if IsProductionEnvironment() {
		home, err := getHomeDir()
		if err != nil {
			return "", err
		}
		configDir = filepath.Join(home, ConfigDirName)
	} else {
		currentDirectory, err := os.Getwd()
		if err != nil {
			return "", err
		}
		configDir = filepath.Join(currentDirectory, ConfigDirName)
	}

	if !isExist(configDir) {
		if err := os.Mkdir(configDir, os.ModePerm); err != nil {
			return "", err
		}
	}

	return filepath.Clean(configDir), nil
}
