.PHONY: buildAll

buildAll: clean install build

install:
	go get -v ./...

build:
	go build -o bin/crumpet src/main.go

clean:
	rm -rf ./bin/