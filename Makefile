build:
	@go build -o bin/health-record src/main.go

test:
	@go test -v ./...
	
run: build
	@./bin/health-record
