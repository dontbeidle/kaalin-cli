APP_NAME := kaalin
VERSION  := $(shell git describe --tags --always --dirty 2>/dev/null || echo "dev")
COMMIT   := $(shell git rev-parse --short HEAD 2>/dev/null || echo "none")
DATE     := $(shell date -u +"%Y-%m-%dT%H:%M:%SZ")
LDFLAGS  := -s -w -X github.com/dontbeidle/kaalin/cmd.version=$(VERSION) -X github.com/dontbeidle/kaalin/cmd.commit=$(COMMIT) -X github.com/dontbeidle/kaalin/cmd.date=$(DATE)

.PHONY: build install test lint clean release-dry release

build:
	go build -ldflags '$(LDFLAGS)' -o bin/$(APP_NAME) .

install:
	go install -ldflags '$(LDFLAGS)' .

test:
	go test -race -cover ./...

lint:
	golangci-lint run ./...

clean:
	rm -rf bin/ dist/

release-dry:
	goreleaser release --snapshot --clean

release:
	goreleaser release --clean
