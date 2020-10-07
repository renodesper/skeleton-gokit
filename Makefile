.DEFAULT_GOAL	:= run
BUILD_DIR   	:= build
SHELL 		  	:= /bin/bash
GOPKGS				:= $(shell go list ./... | grep -v /vendor/)
BINARY 	      := skeleton

ifndef GOOS
  GOOS := $(shell go env GOHOSTOS)
endif

ifndef GOARCH
	GOARCH := $(shell go env GOHOSTARCH)
endif

ifneq ($(filter $(TEST_VERBOSE),$(IS_OK)),)
  GOTEST += -v
endif

.PHONY: run
run:
	go run cmd/main.go

.PHONY: build
build:
	CGO_ENABLED=$(CGO_ENABLED) GOOS=$(GOOS) GOARCH=$(GOARCH) go build -v -o $(BUILD_DIR)/$(BINARY) ./cmd