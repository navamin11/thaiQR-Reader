override filename = thaiQRreader
version = v1.0.0

run:
	@echo "Run Go Application"
	go run cmd/main.go

tidy:
	@echo "Install Packages"
	go mod tidy

build:
	@echo "Build Application"
	env GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o ./bin/${filename}${version}-linux cmd/main.go
	env GOOS=windows GOARCH=amd64 go build -ldflags="-s -w" -o ./bin/${filename}${version}-windows.exe cmd/main.go
	env GOOS=darwin GOARCH=amd64 go build -ldflags="-s -w" -o ./bin/${filename}${version}-mac cmd/main.go