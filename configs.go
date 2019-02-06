package goignore

import (
	"encoding/json"
	"github.com/pkg/errors"
	"io/ioutil"
	"os"
	"path/filepath"
	"runtime"
)

const (
	ConfigDirName         = ".goignore"
	ConfigFileName        = "config.json"
	CustomTemplateDirName = "customs"
	CachedTemplateDirName = "cached"
)

type Configuration struct {
	FirstTimeRun       bool `json:"-"`
	IsRead             bool `json:"-"`
	LastUpdated        string
	SupportedTemplates []string
	CustomTemplates    map[string]string
}

var Config = Configuration{}

func (config *Configuration) Read() error {
	config.IsRead = false
	filePath, err := getConfigFilePath()
	if err != nil {
		return err
	}
	if config.FirstTimeRun {
		return nil
	}
	configContent, _ := ioutil.ReadFile(filePath)
	if err := json.Unmarshal(configContent, config); err != nil {
		return err
	}
	config.IsRead = true
	return nil
}

func (config *Configuration) Save() error {
	filePath, err := getConfigFilePath()
	if err != nil {
		return err
	}
	configNewContent, err := json.Marshal(config)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(filePath, configNewContent, 0644)
	return err
}

func getConfigFilePath() (string, error) {
	configDir, err := getConfigDir()
	if err != nil {
		return "", err
	}
	configFile := filepath.Join(configDir, ConfigFileName)
	return configFile, err
}

func getConfigDir() (string, error) {
	home, err := getHomeDir()
	if err != nil {
		return "", err
	}
	configDir := filepath.Join(home, ConfigDirName)
	if _, err := os.Stat(configDir); os.IsNotExist(err) {
		Config.FirstTimeRun = true
		if err := os.Mkdir(configDir, os.ModePerm); err != nil {
			return "", err
		}
	}
	return configDir, nil
}

func getHomeDir() (string, error) {
	switch runtime.GOOS {
	case "windows":
		return getWindowsHomeDir()
	case "linux":
		return getLinuxHomeDir()
	default:
		return "", errors.New("Cannot find the home directory")
	}
}

func getWindowsHomeDir() (string, error) {
	if home := os.Getenv("HOME"); home != "" {
		return home, nil
	}
	if home := os.Getenv("USERPROFILE"); home != "" {
		return home, nil
	}

	if home := os.Getenv("HOMEDRIVE") + os.Getenv("HOMEPATH"); home != "" {
		return home, nil
	}

	return "", errors.New("Cannot find the home directory.")
}

func getLinuxHomeDir() (string, error) {
	// FIXME: Test on Linux  & Mac
	if home := os.Getenv("HOME"); home != "" {
		return home, nil
	}
	return "", errors.New("Cannot find the home directory.")
}
