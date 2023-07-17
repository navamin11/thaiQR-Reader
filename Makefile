run:
	@echo "Run Go Application"
	go run cmd/main.go

tidy:
	@echo "Install Packages"
	go mod tidy

build:
	@echo "Build Application"
	env GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o ./bin/ThaiQRreader-linux cmd/main.go
	env GOOS=windows GOARCH=amd64 go build -ldflags="-s -w" -o ./bin/ThaiQRreader-windows.exe cmd/main.go
	env GOOS=darwin GOARCH=amd64 go build -ldflags="-s -w" -o ./bin/ThaiQRreader-mac cmd/main.go