help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

build-docker: ## build docker image
	docker build -t new-newt .

run-docker: build-docker ## run project
	docker-compose up -d
	docker-compose logs -f

test: ## test go code
	go test -race ./...

run: ##
	go run cli/main.go