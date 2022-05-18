BINARY_LINUX_NAME=nfcgate_server_go
BUILD_PATH=builds
GOARCH=amd64
OS=linux

all: build

build:
	OS=${OS} GOARCH=${GOARCH} go build -o ./${BUILD_PATH}/${BINARY_LINUX_NAME}

run:
	./${BUILD_PATH}/${BINARY_LINUX_NAME}

deps:
	go mod download

clean:
	rm ./${BUILD_PATH}/${BINARY_LINUX_NAME}