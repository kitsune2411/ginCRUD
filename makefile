build:
	@go build -o bin/ginCRUD.exe

run: build
	@./bin/ginCRUD.exe

test: 
	@go test -v ./...