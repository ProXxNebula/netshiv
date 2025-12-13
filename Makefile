# Makefile for NetShiv

VERSION := v1.0.0
COMMIT := $(shell git rev-parse --short HEAD 2>/dev/null || echo "none")
DATE := $(shell date -u +%Y-%m-%dT%H:%M:%SZ)

LDFLAGS := -ldflags="-s -w -X 'main.Version=$(VERSION)' -X 'main.Commit=$(COMMIT)' -X 'main.Date=$(DATE)'"

BINARY := netshiv
BUILD_DIR := build

.PHONY: all build clean release test install

all: build

build:
	@echo "Building $(BINARY)..."
	go build $(LDFLAGS) -o $(BINARY)
	@echo "Build complete: ./$(BINARY)"

release: clean
	@echo "Building releases..."
	@mkdir -p $(BUILD_DIR)
	GOOS=linux GOARCH=amd64 go build $(LDFLAGS) -o $(BUILD_DIR)/$(BINARY)-linux-amd64
	GOOS=linux GOARCH=arm64 go build $(LDFLAGS) -o $(BUILD_DIR)/$(BINARY)-linux-arm64
	GOOS=windows GOARCH=amd64 go build $(LDFLAGS) -o $(BUILD_DIR)/$(BINARY)-windows-amd64.exe
	GOOS=darwin GOARCH=amd64 go build $(LDFLAGS) -o $(BUILD_DIR)/$(BINARY)-darwin-amd64
	GOOS=darwin GOARCH=arm64 go build $(LDFLAGS) -o $(BUILD_DIR)/$(BINARY)-darwin-arm64
	@echo "Release builds complete in $(BUILD_DIR)/"

clean:
	@echo "Cleaning..."
	rm -f $(BINARY)
	rm -rf $(BUILD_DIR)
	@echo "Clean complete"

test:
	go test -v ./...

install: build
	@echo "Installing to /usr/local/bin..."
	sudo cp $(BINARY) /usr/local/bin/
	@echo "Install complete"
