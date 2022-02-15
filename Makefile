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

.PHONY: prepare
prepare:
	$(GOCMD) mod download
	$(GOINSTALL) honnef.co/go/tools/cmd/staticcheck@latest

.PHONY: build
build: prepare
	$(GOBUILD) -race ./...

.PHONY: test
test:
	$(GOTEST) -race ./...

.PHONY: clean
clean:
	$(GOCLEAN)
	$(GOCLEAN) --testcache ./...

build-cmds:
	go build github.com/devlights/gomy/cmd/splitbin
	go build github.com/devlights/gomy/cmd/splitrec
	go build github.com/devlights/gomy/cmd/disphex
	go build github.com/devlights/gomy/cmd/extract

clean-cmds:
	rm -f ./splitbin ./splitrec ./disphex ./extract
