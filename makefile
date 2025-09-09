APP_NAME := slow-lang
PKG := ./...

.PHONY: all build run test clean fmt

all: build

build:
	go build -o $(APP_NAME)

run: build
	./$(APP_NAME)

test:
	go test $(PKG)

clean:
	go clean -testcache
	rm -f $(APP_NAME)

fmt:
	go fmt $(PKG)

