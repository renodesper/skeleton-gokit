.DEFAULT_GOAL	:= run
BUILD_DIR   	:= build
SHELL 		  	:= /bin/bash
BINARY 	      := skeletond
CGO_ENABLED   := 1
GOPKGS				:= $(shell go list ./... | grep -v /vendor/)
GOTEST				:= go test -v -count=1

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
	go run skeletond/main.go

.PHONY: build
build:
	CGO_ENABLED=$(CGO_ENABLED) GOOS=$(GOOS) GOARCH=$(GOARCH) go build -v -o $(BUILD_DIR)/$(BINARY) ./skeletond

.PHONY: install
install:
	CGO_ENABLED=$(CGO_ENABLED) GOOS=$(GOOS) GOARCH=$(GOARCH) go install -a -tags netgo -ldflags '-w -extldflags "-static"' ./skeletond

.PHONY: test
test:
	$(GOTEST) -cover ./...