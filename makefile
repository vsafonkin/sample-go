.DEFAULT_GOAL := build

fmt:
	go fmt ./...
.PHONY:fmt

build: fmt
	go build -o bin/sample sample.go
	./bin/sample
.PHONY:build

checkrace:
	go run --race sample.go
.PHONY:checkrace

debug:
	go build -gcflags='all=-N -l' -o bin/debug sample.go
.PHONY:debug
