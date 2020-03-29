TOP:=$(realpath ..)

SHELL := /bin/bash
version := $(shell git rev-list --count HEAD)
commit := $(shell git describe --always --long --dirty)
built_at := $(shell date +%FT%T%z)
built_by := $(shell git config --get user.email)

ldflags := -X main.version=v${version} -X main.commit=${commit}
ldflags += -X main.builtAt=${built_at} -X main.builtBy=${built_by}

cmd := ./bin/hmydb

export GO111MODULE:=on

.PHONY: all linux_static

all:
	go build -o $(cmd) -ldflags="$(ldflags) -w" cmd/main.go

test:
	./test/debug.sh

linux_static:
	go build -o $(cmd) -ldflags="$(ldflags) -w -extldflags \"-static\"" cmd/main.go

clean:
	rm -f bin/*
