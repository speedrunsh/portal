all: requirements all-targets

requirements:
	@go get

linux-amd64:
	GOOS=linux GOARCH=amd64 go build -o ../dist/linux/amd64/portal

linux-arm64:
	GOOS=linux GOARCH=arm64 go build -o ../dist/linux/arm64/portal

linux: linux-amd64 linux-arm64

darwin-amd64:
	GOOS=darwin GOARCH=amd64 go build -o ../dist/darwin/amd64/portal

darwin-arm64:
	GOOS=darwin GOARCH=arm64 go build -o ../dist/darwin/arm64/portal

darwin: darwin-amd64 darwin-arm64

all-targets: linux darwin