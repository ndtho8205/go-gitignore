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

// APIClient is an gitignore.io API client.
type APIClient struct {
	client *http.Client
}

// Client is the global gitignore.io API client variable.
var Client = NewAPIClient()

// NewAPIClient creates new gitignore.io API client.
func NewAPIClient() *APIClient {
	return &APIClient{
		client: &http.Client{
			Timeout: time.Duration(5 * time.Second),
		},
	}
}

// MakeRequest makes a GET request.
func (apiClient *APIClient) MakeRequest(url string) (string, error) {
	response, err := apiClient.client.Get(url)
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

// GetTemplateList makes a GET request to gitignore.io API to get a list of all supported templates.
func (apiClient *APIClient) GetTemplateList() ([]string, error) {
	response, err := apiClient.MakeRequest(createURL("list"))
	if err != nil {
		return nil, err
	}

	templateList := make([]string, 0, 100)
	for _, line := range strings.Split(response, "\n") {
		templateList = append(templateList, strings.Split(line, ",")...)
	}

	return templateList, nil
}

// GetGitignoreContent makes a GET request to gitignore.io API to get a .gitignore file content given template names.
func (apiClient *APIClient) GetGitignoreContent(templates string) (string, error) {
	result, err := apiClient.MakeRequest(createURL(templates))
	return result, err
}

func createURL(path string) string {
	return defaultAPIURL + url.PathEscape(path)
}
