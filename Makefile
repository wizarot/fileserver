.DEFAULT_GOAL := help

GO = go
V = 0
Q = $(if $(filter 1,$V),,@)
M = $(shell printf "\033[34;1m▶\033[0m")
RICE = rice append --exec

PACKAGE = fileserver
CMD = $(CURDIR)
DATE ?= $(shell date +%FT%T%z)
BIN = $(CURDIR)/bin
VERSION ?= $(shell git describe --tags --always --dirty --match=v* 2> /dev/null || \
			cat $(CURDIR)/.version 2> /dev/null || echo v0)

export GO111MODULE=on

all: fmt build build-linux

.PHONY: run
run: build ## run
	./bin/$(PACKAGE)

.PHONY: fmt
fmt: ; $(info $(M) running gofmt…) @ ## Run gofmt on all source files
	$Q $(GO) fmt ./...

.PHONY: build
build: ; $(info $(M) building executable…) @ ## Build program binary
	$Q $(GO) build \
		-tags release \
		-ldflags '-X $(PACKAGE)/cmd.Version=$(VERSION) -X $(PACKAGE)/cmd.BuildDate=$(DATE)' \
		-o $(BIN)/$(PACKAGE) $(CMD)/*.go && $(RICE) $(BIN)/$(PACKAGE)

.PHONY: build-osx
build-osx: ; $(info $(M) building osx executable…) @ ## Build program binary for osx
	$Q CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 $(GO) build \
		-tags release \
		-ldflags '-X $(PACKAGE)/cmd.Version=$(VERSION) -X $(PACKAGE)/cmd.BuildDate=$(DATE)' \
		-o $(BIN)/$(PACKAGE) $(CMD)/*.go && $(RICE) $(BIN)/$(PACKAGE)

.PHONY: build-linux
build-linux: ; $(info $(M) building linux executable…) @ ## Build program binary for linux
	$Q CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GO) build \
		-tags release \
		-ldflags '-X $(PACKAGE)/cmd.Version=$(VERSION) -X $(PACKAGE)/cmd.BuildDate=$(DATE)' \
		-o $(BIN)/$(PACKAGE)_unix $(CMD)/*.go && $(RICE) $(BIN)/$(PACKAGE)_unix

.PHONY: build-win
build-win: ; $(info $(M) building win executable…) @ ## Build program binary for windows
	$Q CGO_ENABLED=0 GOOS=windows GOARCH=amd64 $(GO) build \
		-tags release \
		-ldflags '-X $(PACKAGE)/cmd.Version=$(VERSION) -X $(PACKAGE)/cmd.BuildDate=$(DATE)' \
		-o $(BIN)/$(PACKAGE)_win.exe $(CMD)/*.go && $(RICE) $(BIN)/$(PACKAGE)_win.exe

.PHONY: build-osx build-linux build-win

# Tools
$(BIN):
	@mkdir -p $@

# Misc
.PHONY: clean
clean: ; $(info $(M) cleaning…)	@ ## Cleanup everything
	@rm -rf $(BIN)
	@rm -rf test/tests.* test/coverage.*

.PHONY: help
help:
	@grep -E '^[ a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | \
		awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-15s\033[0m %s\n", $$1, $$2}'

.PHONY: version
version:
	@echo $(VERSION)%