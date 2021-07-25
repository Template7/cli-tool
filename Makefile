.PHONY: build all

build:
	go build -o template7-cli

build-linux:
	GOOS=linux GOARCH=amd64 go build
