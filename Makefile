build:
	@go build -o bin/go_ecom cmd/main.go

test:
	@go test -v ./...

run: build
	@./bin/go_ecom