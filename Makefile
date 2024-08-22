.PHONY: generate build run

generate:
	go generate ./...

build: generate
	go build .

run: build
	./tenbounce start

test:
	go test ./...

deploy:
	./deploy.sh