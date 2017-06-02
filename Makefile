# go option
GO        ?= go
TAGS      :=
LDFLAGS   := -s
GOFLAGS   :=
BINDIR    := $(CURDIR)/bin

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

.PHONY: build
build: 
	@echo "================="
	@echo "Building helm-certgen plugin @ $(BINDIR)"
	$(GO) build $(GOFLAGS) -tags '$(TAGS)' -ldflags '$(LDFLAGS)' -o $(BINDIR)/helm-certgen main.go
