# Makefile for use with MinGW make (will add git build sha to compiled exe)
.ONESHELL:
SHELL = ps
export GOOS=windows
export GOARCH=amd64
export CGO_ENABLED=0

BUILD_FLAGS_SERVER=-o ./bin/bug_server.exe
BUILD_FLAGS_WASM=-o ./web/app.wasm

# targets
all: ui

# force compilation
.PHONY: all ui

# build targets
ui:
	GOARCH=wasm GOOS=js go build $(BUILD_FLAGS_WASM) ./cmd/wasm
	go build $(BUILD_FLAGS_SERVER) ./cmd/server
