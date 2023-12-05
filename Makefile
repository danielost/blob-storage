c=run service

build:
	@go build -o bin/blob-storage

run: build
	@./bin/blob-storage $(c)

test:
	@go test -v ./...