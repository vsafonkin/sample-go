.DEFAULT_GOAL := build

fmt:
	go fmt ./...
.PHONY:fmt

imports:
	gopls imports -w *.go
.PHONY:imports

build: fmt
	go build -ldflags='-s -w' -o bin/sample sample.go
	./bin/sample
.PHONY:build

checkrace:
	go run --race sample.go
.PHONY:checkrace

debug:
	go build -gcflags='all=-N -l' -o bin/debug sample.go
.PHONY:debug

test:
	go test -v ./...
.PHONY:test

bench:
	go test -bench=. ./...
