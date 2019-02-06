package goignore

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

const (
	defaultAPIURL = "https://www.gitignore.io/api/"
)

func createUrl(path string) string {
	return defaultAPIURL + url.PathEscape(path)
}

// MakeRequest makes a GET request
func MakeRequest(url string) (string, error) {
	client := http.Client{
		Timeout: time.Duration(5 * time.Second),
	}
	response, err := client.Get(url)
	if err != nil {
		return "", err
	}
	if response.StatusCode != http.StatusOK {
		return "", fmt.Errorf("%s", response.Status)
	}

	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	return string(body), err
}

// GetTemplateList makes a GET request to gitignore.io API to get a list of all supported templates
func GetTemplateList() ([]string, error) {
	fmt.Println("get list")
	response, err := MakeRequest(createUrl("list"))
	if err != nil {
		return nil, err
	}
	templateList := make([]string, 0, 100)
	for _, line := range strings.Split(response, "\n") {
		templateList = append(templateList, strings.Split(line, ",")...)
	}
	return templateList, err
}

// GetGitignoreContent makes a GET request to gitignore.io API to get a .gitignore file content given template names
func GetGitignoreContent(templates string) (string, error) {
	result, err := MakeRequest(createUrl(templates))
	return result, err
}
