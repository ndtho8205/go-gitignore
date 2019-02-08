package goignore

import (
	"errors"
	"os"
	"runtime"
)

// IsProductionEnvironment checks if current environment is production.
func IsProductionEnvironment() bool {
	env := os.Getenv("APP_ENV")
	return env != "dev"
}

func isExist(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	}
	return true
}

func getHomeDir() (string, error) {
	switch runtime.GOOS {
	case "windows":
		return getWindowsHomeDir()
	case "linux":
		return getLinuxHomeDir()
	default:
		return "", errors.New("cannot find the home directory")
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

	return "", errors.New("cannot find the home directory")
}

func getLinuxHomeDir() (string, error) {
	// FIXME: Test on Linux  & Mac
	if home := os.Getenv("HOME"); home != "" {
		return home, nil
	}
	return "", errors.New("cannot find the home directory")
}
