BINARY_NAME=./bin/AVDClientUrlTool
SRC=./cmd/AVDClientUrlTool

build:
	GOARCH=amd64 GOOS=darwin go build -o ${BINARY_NAME}-darwin-amd64 ${SRC}/main.go
	GOARCH=amd64 GOOS=linux go build -o ${BINARY_NAME}-linux-amd64 ${SRC}/main.go
	GOARCH=amd64 GOOS=windows go build -o ${BINARY_NAME}-windows-amd64 ${SRC}/main.go

run:
	go run ${SRC}/main.go

clean:
	go clean
	rm ${BINARY_NAME}-darwin-amd64
	rm ${BINARY_NAME}-linux-amd64
	rm ${BINARY_NAME}-windows-amd64
