.PHONY: install test coverage build

install:
	glide install

test: install
	go test -v $(shell glide novendor)

coverage: install
	go test -cover $(shell glide novendor)

build: install
	go build -v -o ./build/mechanic-store
