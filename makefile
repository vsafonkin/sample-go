all: fmt build
	
fmt:
	go fmt ./...

lint:
	go vet ./...

imports:
	gopls imports -w sample.go

build: fmt
	go build -ldflags='-s -w' -o bin/sample sample.go
	go build -gcflags='all=-N -l' -o bin/debug sample.go

run:
	./bin/sample $(ARGS)

debug:
	env GODEBUG=gctrace=1 ./bin/debug

checkrace:
	go run --race sample.go

test:
	go test -v ./...

bench:
	go test -bench=. ./... -benchmem

opt:
	go build -gcflags -m -o ./bin/sample sample.go
