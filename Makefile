build:
	@go build -o bin/cache 

test:
	@go test ./... -v

run: build
	@./bin/cache
