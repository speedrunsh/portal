all: requirements linux darwin

clean:
	rm -rf ../dist

requirements:
	@go get

linux-amd64: requirements
	GOOS=linux GOARCH=amd64 go build -o ../dist/linux/amd64/portal

linux-arm64: requirements
	GOOS=linux GOARCH=arm64 go build -o ../dist/linux/arm64/portal

linux: requirements linux-amd64 linux-arm64

darwin-amd64: requirements
	GOOS=darwin GOARCH=amd64 go build -o ../dist/darwin/amd64/portal

darwin-arm64: requirements
	GOOS=darwin GOARCH=arm64 go build -o ../dist/darwin/arm64/portal

darwin: requirements darwin-amd64 darwin-arm64

all-targets: requirements linux darwin