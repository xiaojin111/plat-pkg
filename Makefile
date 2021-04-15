SOURCE_FILES?=./...

export PATH := ./bin:$(PATH)
export GO111MODULE := on
export GOPATH := $(shell go env GOPATH)
export GOPROXY := https://goproxy.io,direct
export GOPRIVATE := gitee.com/jt-heath/*
export GOVERSION := $(shell go version | awk '{print $$3}')
# GORELEASER is the path to the goreleaser binary.
export GORELEASER := $(shell which goreleaser)

# all is the default target
all: release
.PHONY: all

# Install all the build and lint dependencies
setup:
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s
	#curl -L https://git.io/misspell | sh
.PHONY: setup

# Update go packages
go-update:
	@echo "Updating go packages..."
	@go get -u -t ./...
	@echo "go mod tidy..."
	@$(MAKE) go-mod-tidy
.PHONY: go-update

# Clean go.mod
go-mod-tidy:
	@go mod tidy -v
	# @git --no-pager diff HEAD
	# @git --no-pager diff-index --quiet HEAD
.PHONY: go-mod-tidy

# Reset go.mod
go-mod-reset:
	@rm -f go.sum
	@sed -i '' -e '/^require/,/^)/d' go.mod
	@go mod tidy -v
	# @git --no-pager diff HEAD
	# @git --no-pager diff-index --quiet HEAD
.PHONY: go-mod-tidy

generate:
	@go generate ./...
.PHONY: generate

# Format go files
format:
	@goimports -w ./
.PHONY: format

# Run all the linters
lint:
	@golangci-lint run --allow-parallel-runners
.PHONY: lint

# Go build all
build:
	@go build ./... > /dev/null
.PHONY: build

# Go test all
test:
	@go test -v ./...
.PHONY: test

# Run all code checks
ci: generate format lint build test
.PHONY: ci

# Release wia goreleaser
release:
	@[ -x "$(GORELEASER)" ] || ( echo "goreleaser not installed"; exit 1)
	@goreleaser --snapshot --skip-publish --rm-dist
.PHONY: release

.DEFAULT_GOAL := all
