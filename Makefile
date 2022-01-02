PROJECT_NAME := "new-newt"
PKG := "github.com/rwxd/$(PROJECT_NAME)"
PKG_LIST := $(shell go list ${PKG}/...)
GO_FILES := $(shell find . -name '*.go' | grep -v _test.go)

help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'


all: build

build-docker: ## build docker image
	docker build -t new-newt .

run-docker: build-docker ## run project
	docker-compose up -d
	docker-compose logs -f

test: ## test go code
	go test -race ./...

run: ##
	go run cli/main.go

dep: ## get the dependencies
	@go get -v -d ./...

build: dep ## Build the binary file
	@go build -v $(PKG)

clean: ## Remove previous build
	@rm -f $(PROJECT_NAME)