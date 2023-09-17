.DEFAULT_GOAL := build

fmt:
	go fmt ./...
.PHONY:fmt

build: fmt
	go build -o dist/sample sample.go
	./dist/sample
.PHONY:build

checkrace:
	go run --race sample.go
.PHONY:checkrace
