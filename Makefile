export grpc_address := 1
export grpc_port := 2

GOBIN=./cmd/main

## build: Build go binary
build:
	go build -o $(GOBIN)

## run: Run go server
.PHONY: run
run:
	$(GOBIN)

## get: Run go get missing dependencies
get:
	go get ./...

## deploy: Run commands to deploy app to container
deploy:
	make get
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build \
          -ldflags='-w -s -extldflags "-static"' -a \
          -o /go/bin/main .

.PHONY: help
all: help
help: Makefile
	@echo
	@echo " Choose a command"
	@echo
	@sed -n 's/^##//p' $< | column -t -s ':' |  sed -e 's/^/ /'
	@echo

.DEFAULT_GOAL := help