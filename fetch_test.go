package goignore

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestMakeRequest_Ok(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		assertEqual(t, request.Method, "GET")
		_, _ = fmt.Fprint(writer, "goignore")
	})
	server := httptest.NewServer(mux)

	response, err := NewAPIClient().MakeRequest(server.URL)
	if err != nil {
		t.Fatalf("MakeRequest(): %v", err)
	}

	expected := "goignore"
	assertEqual(t, response, expected)
}

func TestMakeRequest_WrongUrl(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		assertEqual(t, request.Method, "GET")
		_, _ = fmt.Fprint(writer, "goignore")
	})
	server := httptest.NewServer(mux)

	_, err := NewAPIClient().MakeRequest(server.URL + "/goignore")
	if err != nil {
		t.Fatalf("MakeRequest(): %v", err)
	}
}

func TestGetTemplateList(t *testing.T) {
	templateList, err := NewAPIClient().GetTemplateList()
	if err != nil {
		t.Fatalf("GetTemplateList(): %v", err)
	}

	assertEqual(t, templateList[0], "1c")
	assertEqual(t, templateList[len(templateList)-1], "zukencr8000")
}

func TestGetGitignoreContent_Ok(t *testing.T) {
	content, err := NewAPIClient().GetGitignoreContent("go,android")
	if err != nil {
		t.Fatalf("GetGitignoreContent(): %v", err)
	}

	if len(content) <= 0 {
		t.Fatalf("Expect (Len: > 0) (Type: string) - Got (Len: %v) (Type: %T)", len(content), content)
	}
}

func TestGetGitignoreContent_WrongTemplates(t *testing.T) {
	_, err := NewAPIClient().GetGitignoreContent("go ignore")
	if err == nil {
		t.Fatalf("GetGitignoreContent(): %v", err)
	}
}

func assertEqual(t *testing.T, result interface{}, expect interface{}) {
	if result != expect {
		t.Fatalf("Expect (Value: %v) (Type: %T) - Got (Value: %v) (Type: %T)", expect, expect, result, result)
	}
}
