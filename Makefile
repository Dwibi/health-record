build:
	@go build -o bin/health-record src/main.go

test:
	@go test -v ./...
	
run: build
	@./bin/health-record

migration:
	@migrate create -ext sql -dir db/migrations $(filter-out $@,$(MAKECMDGOALS))

migrate-up:
	@go run cmd/migrate/main.go up

migrate-down:
	@go run cmd/migrate/main.go down
