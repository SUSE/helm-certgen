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


.PHONY: all
all: clean test-all build

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
	rm -rf ${GOBIN}/helm-certgen
	rm -rf helm-certgen
	rm -rf debug
	rm -rf build
	rm -rf bin/helm-certgen
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
build: 
	@echo "================="
	@echo "Building helm-certgen plugin @ $(BINDIR)"
	$(GO) build $(GOFLAGS) -tags '$(TAGS)' -ldflags '$(LDFLAGS)' -o $(BINDIR)/helm-certgen main.go

.PHONY: build
build-all:
	@echo "================="
	@echo "Building helm-certgen plugin for all the platforms@ $(BINDIR)"
	GOOS=linux GOARCH=amd64 $(GO) build $(GOFLAGS) -tags '$(TAGS)' -ldflags '$(LDFLAGS)' -o dist/linux-amd64/certgen main.go
	GOOS=windows GOARCH=amd64 $(GO) build $(GOFLAGS) -tags '$(TAGS)' -ldflags '$(LDFLAGS)' -o dist/windows-amd64/certgen main.go
	GOOS=darwin GOARCH=amd64 $(GO) build $(GOFLAGS) -tags '$(TAGS)' -ldflags '$(LDFLAGS)' -o dist/darwin-amd64/certgen main.go

