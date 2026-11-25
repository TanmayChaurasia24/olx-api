.PHONY: build run

build:
	go build -o bin/olx-api-server main.go

run:
	./bin/olx-api-server

