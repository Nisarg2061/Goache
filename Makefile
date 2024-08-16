build:
	@go build -o bin/cache 

test:
	@go test ./... -v

run: test build
	@./bin/cache
