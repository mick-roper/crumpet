.PHONY: buildAll

buildAll: clean install test build

install:
	go get -v ./...

test:
	go test -v ./...

build:
	go build -o bin/crumpet src/main.go

clean:
	rm -rf ./bin/