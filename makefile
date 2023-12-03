BINARY_NAME=ginCRUD

build:
	@go build -o bin/${BINARY_NAME}.exe

run: build
	@./bin/${BINARY_NAME}.exe

migrate: 
	@go run migrate/migrate.go

test: 
	@go test -v ./...