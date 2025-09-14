APP_NAME := slow-lang
PKG := ./...

all: build

build:
	go build -o $(APP_NAME)

test:
	go test $(PKG)

clean:
	go clean -testcache
	rm -f $(APP_NAME)

