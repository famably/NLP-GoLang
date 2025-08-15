.PHONY: run test build docker

run:
	go run main.go

test:
	go test -v ./...

build:
	go build -o bin/chip-api main.go

docker:
	docker build -t chip-api .