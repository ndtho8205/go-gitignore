PKGS := $(shell go list ./... | grep -v /vendor)
GOBIN := $(GOPATH)/bin

BINARY := ggi
PLATFORMS := windows linux darwin
os = $(word 1, $@)
VERSION ?= latest

all: help

$(GOBIN)/%:
	go get -u $(REPOSITORY)


.PHONY: fmt
fmt: ## Run gofmt
	@echo "Run gofmt"
	@go fmt ${PKGS}


GOLINT = $(GOBIN)/golint
$(GOBIN)/golint: REPOSITORY=golang.org/x/lint/golint
.PHONY: lint
lint: $(GOLINT) fmt ## Run golint
	@echo "Run golint"
	$(GOLINT) -set_exit_status ${PKGS}

  
.PHONY: test
test: lint ## Test
	go test $(PKGS)


.PHONY: $(PLATFORMS)
$(PLATFORMS):
	mkdir -p release
	GOOS=$(os) GOARCH=amd64 go build -o release/$(BINARY)-$(VERSION)-$(os)-amd64


.PHONY: release
release: windows linux darwin ## Release


.PHONY: clean
clean: ## Cleanup everything
	@echo "Cleanup everything"
	@rm -rf test/tests.* test/coverage.*
	@rm -rf release

	
.PHONY: help
help:
	@echo
	@grep -E '^[ a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | \
		awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-15s\033[0m %s\n", $$1, $$2}'
	@echo
