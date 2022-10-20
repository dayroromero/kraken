SHELL := /bin/bash
GO := GO15VENDOREXPERIMENT=1 go
NAME := kraken
OS := $(shell uname)
MAIN_GO := main.go
ROOT_PACKAGE := $(GIT_PROVIDER)/$(ORG)/$(NAME)
GO_VERSION := $(shell $(GO) version | sed -e 's/^[^0-9.]*\([0-9.]*\).*/\1/')
PACKAGE_DIRS := $(shell $(GO) list ./... | grep -v /vendor/)
PKGS := $(shell go list ./... | grep -v /vendor | grep -v generated)
PKGS := $(subst  :,_,$(PKGS))
BUILDFLAGS := ''
CGO_ENABLED = 0
VENDOR_DIR=vendor

proto:
	protoc pkg/grpc/*.proto --go_out=. --go-grpc_out=require_unimplemented_servers=false:.

server:
	go run cmd/main.go

.PHONY: build
build:
	CGO_ENABLED=$(CGO_ENABLED) $(GO) build -ldflags $(BUILDFLAGS) -o bin/$(NAME) cmd/$(MAIN_GO)

linux:
	CGO_ENABLED=$(CGO_ENABLED) GOOS=linux GOARCH=amd64 $(GO) build -ldflags $(BUILDFLAGS) -o bin/$(NAME) cmd/$(MAIN_GO)
