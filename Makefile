# go option
GO        ?= go
TAGS      :=
LDFLAGS   := -s
GOFLAGS   :=

ifdef GOBIN
    BINDIR:=$(GOBIN)
else
    BINDIR:=$(CURDIR)/bin
endif

ifeq ($(strip $(VERSION)),)
    export VERSION := $(shell scripts/build/get_version.sh)
endif

.PHONY: all
all: clean-all test-all build-all

.PHONY: tools
tools:
	@echo "================="
	@echo "Installing dependancies"
	go get -u github.com/golang/lint/golint
	go get golang.org/x/tools/cmd/cover
	go get github.com/tools/godep
	go get github.com/spf13/cobra

.PHONY: clean
clean:
	rm -rf build

.PHONY: clean-all
clean-all: clean
	rm -rf bin
	rm -rf debug
	rm -rf build
	rm -rf dist

.PHONY: test-all
test-all: test-unit
test-all: test-style

.PHONY: test-unit
test-unit:
	@echo "================="
	@echo "Running unit test"
	go test $(go list ./... | grep -v vendor)


.PHONY: test-style
test-style:
	@echo "================="
	@echo "Running style checks"
	scripts/validate-go.sh

.PHONY: coverage
coverage:
	@echo "================="
	@echo "Running tests with coverage tool"
	./scripts/test-coverage.sh

.PHONY: build
build: clean
	@echo "================="
	./scripts/build/build.sh

.PHONY: build-all
build-all: clean-all
	@echo "================="
	@echo "Building helm-certgen plugin for all the platforms@ $(BINDIR)"
	./scripts/build/build-all-platforms.sh


