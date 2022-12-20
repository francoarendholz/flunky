VERSION := $(shell cat cfg/VERSION)
COMMIT := $(shell git rev-parse HEAD)

build-local:
	go build -ldflags="-X 'github.com/francoarendholz/flunky/base.FlunkyVersion=$(VERSION)' -X 'github.com/francoarendholz/flunky/base.FlunkyCommit=$(COMMIT)'" .

test:
	go clean -testcache && go test ./... -v -cover