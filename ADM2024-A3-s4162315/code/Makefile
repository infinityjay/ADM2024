# Binary name
BINARY=ADM2024-A3
# Builds the project
build:
		go build -o "./cmd/${BINARY}"
release:
		GOARCH=amd64 GOOS=darwin go build -o "./cmd/${BINARY}-darwin" main.go
		GOARCH=amd64 GOOS=linux go build -o "./cmd/${BINARY}-linux" main.go
		GOARCH=amd64 GOOS=windows go build -o "./cmd/${BINARY}-windows" main.go
# Cleans our projects: deletes binaries
clean:
		rm -rf cmd