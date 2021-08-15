.PHONY: build all

build:
	go build -o template7-cli

build-docker:
	docker-compose -f build/docker-compose.yaml build

build-linux:
	GOOS=linux GOARCH=amd64 go build
