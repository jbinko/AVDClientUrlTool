BINARY_NAME=./bin/AVDClientUrlTool
SRC=./cmd/AVDClientUrlTool

build:
	go build -o ${BINARY_NAME} ${SRC}/main.go
	GOARCH=amd64 GOOS=darwin go build -o ${BINARY_NAME}-darwin ${SRC}/main.go
	GOARCH=amd64 GOOS=linux go build -o ${BINARY_NAME}-linux ${SRC}/main.go
	GOARCH=amd64 GOOS=windows go build -o ${BINARY_NAME}-windows ${SRC}/main.go

run:
	go run ${SRC}/main.go

clean:
	go clean
	rm ${BINARY_NAME}
	rm ${BINARY_NAME}-darwin
	rm ${BINARY_NAME}-linux
	rm ${BINARY_NAME}-windows
