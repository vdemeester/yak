PACKAGES=$(shell go list ./... | grep -v /vendor/)
REVISION=$(shell git rev-parse HEAD)
GO_LDFLAGS=-s -w -X github.com/vdemeester/yak/version.Version=$(REVISION)

export GO111MODULE=on

all:
	go build -v -ldflags '${GO_LDFLAGS}' ./cmd/yak

static:
	CGO_ENALBED=0 go build -v -ldflags '${GO_LDFLAGS} -extldflags "-static"' ./cmd/yak

install:
	@install bee $HOME/bin/bee
