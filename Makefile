BINARY      := green
CMD_PATH    := ./cmd/server
GOLINT      := $(shell go env GOPATH)/bin/golangci-lint

.PHONY: build run lint install-lint

build:
	go build -o bin/$(BINARY) $(CMD_PATH)

run:
	go run $(CMD_PATH)

lint: $(GOLINT)
	$(GOLINT) run ./...

$(GOLINT):
	go install github.com/golangci/golangci-lint/v2/cmd/golangci-lint@latest
