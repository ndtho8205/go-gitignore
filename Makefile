PKGS := $(shell go list ./... | grep -v /vendor)
GOBIN := $(GOPATH)/bin

BINARY := goignore
BUILDDIR := build
PLATFORMS := linux windows darwin
ARCH := 386 amd64

VERSION ?= latest

all: help

$(GOBIN)/%:
	@go get -u $(REPOSITORY)


.PHONY: fmt
fmt: ## Run gofmt
	@echo "Run gofmt"
	@go fmt ${PKGS}


GOLINT = $(GOBIN)/golint
$(GOBIN)/golint: REPOSITORY=golang.org/x/lint/golint
.PHONY: lint
lint: $(GOLINT) fmt ## Run golint
	@echo "Run golint"
	@$(GOLINT) -set_exit_status ${PKGS}

  
.PHONY: test
test: lint ## Test
	@echo "Run unit test"
	@go test $(PKGS)


build = $GOOS=$(1) GOARCH=$(2) go build -ldflags "-s -w" -o $(BUILDDIR)/$(BINARY)-$(VERSION)-$(1)-$(2)$(3) ./cmd/goignore

linux: build/linux
build/linux:
	$(call build,linux,amd64,)

windows: build/windows
build/windows:
	$(call build,windows,amd64,.exe)
	$(call build,windows,386,.exe)

darwin: build/darwin
build/darwin:
	$(call build,darwin,amd64,)

.PHONY: deploy
deploy: windows linux darwin ## Deploy


.PHONY: clean
clean: ## Cleanup everything
	@echo "Cleanup everything"
	@rm -rf test/tests.* test/coverage.*
	@rm -rf $(BUILDDIR)

	
.PHONY: help
help:
	@echo
	@grep -E '^[ a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | \
		awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-15s\033[0m %s\n", $$1, $$2}'
	@echo
