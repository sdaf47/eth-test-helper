.PHONY:  build run fmt clean lint vet tools

.DEFAULT_GOAL: all

PROGRAMM=eth-test-helper

all: fmt build

build: clean get fmt vet
	go build -o build/${PROGRAMM} ./cmd/${PROGRAMM}/main.go
	chmod +x build/${PROGRAMM}
	cp build/${PROGRAMM} ~/bin/${PROGRAMM}

clean:
	rm -rf ./build/*

fmt:
	go fmt ./...

lint:
	golint ./...

get:
	go get ./...

vet:
	go vet ./...

run: build
	./build/${PROGRAMM}
