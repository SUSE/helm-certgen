# go option
GO        ?= go
TAGS      :=
LDFLAGS   := -s
GOFLAGS   :=
BINDIR    := $(CURDIR)/bin

.PHONY: tools
tools:
	go get -u github.com/golang/lint/golint
	go get golang.org/x/tools/cmd/cover
	go get github.com/tools/godep
	go get github.com/spf13/cobra

.PHONY: all
all: clean test build

.PHONY: clean
clean:
	rm -rf ${GOBIN}/helm-certgen
	rm -rf helm-certgen
	rm -rf debug
	rm -rf build

test:
	@echo "Test"
	go test -v ./...

.PHONY: build
build: 
	@echo "Building helm-certgen plugin @ $(BINDIR)"
	GOBIN=$(BINDIR) $(GO) install $(GOFLAGS) -tags '$(TAGS)' -ldflags '$(LDFLAGS)' github.com/saurabhsurana/helm-certgen/cmd/...

