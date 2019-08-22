SOURCE_FILES?=./...

export PATH := ./bin:$(PATH)
export GO111MODULE := on
export GOPROXY := https://goproxy.cn


# Install all the build and lint dependencies
setup:
	curl -sfL https://install.goreleaser.com/github.com/golangci/golangci-lint.sh | sh
	curl -L https://git.io/misspell | sh
	go mod download
.PHONY: setup

# Update go packages
go-update:
	@echo "Updating go packages..."
	@go get -u all
	@echo "go mod tidy..."
	@$(MAKE) go-mod-tidy
.PHONY: go-update

# Clean go.mod
go-mod-tidy:
	@go mod tidy -v
	# @git --no-pager diff HEAD
	# @git --no-pager diff-index --quiet HEAD
.PHONY: go-mod-tidy

# Format go files
format:
	@goimports -w ./
.PHONY: format

# Run all the linters
lint:
	@./bin/golangci-lint run
.PHONY: lint

# Go build all
build:
	@go build ./... > /dev/null
.PHONY: build

# Go test all
test:
	@go test ./...
.PHONY: test

# Run all code checks
ci: go-update lint build test
.PHONY: ci

.DEFAULT_GOAL := ci
