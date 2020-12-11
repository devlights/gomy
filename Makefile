GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOINSTALL=$(GOCMD) install
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GORUN=$(GOCMD) run
GOGENERATE=$(GOCMD) generate

PRJ_NAME=gomy
GITHUB_USER=devlights
PKG_NAME=github.com/$(GITHUB_USER)/$(PRJ_NAME)

.PHONY: all
all: clean build test

.PHONY: build
build:
	$(GOBUILD) ./...

.PHONY: test
test:
	$(GOTEST) ./...

.PHONY: clean
clean:
	$(GOCLEAN)
	$(GOCLEAN) --testcache ./...
