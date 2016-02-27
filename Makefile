BIN_NAME="ls-hosts"
VERSION=$(shell cat ./VERSION)
PACKAGES=$(shell go list ./...)
LDFLAGS="-s -w -X main.version=${VERSION}"

all: get-deps build unit

help:
	@echo "make <target>"
	@echo " build"
	@echo " build-all: includes build-osx, build-linux, build-windows"
	@echo " build-osx"
	@echo " build-linux"
	@echo " build-windows"
	@echo " unit"
	@echo " get-deps"

get-deps:
	@go get -v gopkg.in/ini.v1
	@go get -v github.com/aws/aws-sdk-go
	@go get -v github.com/olekukonko/tablewriter

clean:
	@rm -rf ./build/*

build:
	@echo "output to ./build/${BIN_NAME}"
	@go build -ldflags ${LDFLAGS} -o ./build/${BIN_NAME} main.go

build-all: build-osx build-linux build-windows

build-osx:
	@echo "output to ./build/${BIN_NAME}-darwin-amd64"
	@GOOS=darwin GOARCH=amd64 go build -ldflags ${LDFLAGS} -o ./build/${BIN_NAME}-darwin-amd64 main.go

build-linux:
	@echo "output to ./build/${BIN_NAME}-linux-amd64"
	@GOOS=linux GOARCH=amd64 go build -ldflags ${LDFLAGS} -o ./build/${BIN_NAME}-linux-amd64 main.go

build-windows:
	@echo "output to ./build/${BIN_NAME}-windows-amd64"
	@GOOS=windows GOARCH=amd64 go build -ldflags ${LDFLAGS} -o ./build/${BIN_NAME}-windows-amd64 main.go
	@echo "output to ./build/${BIN_NAME}-windows-386"
	@GOOS=windows GOARCH=386 go build -ldflags ${LDFLAGS} -o ./build/${BIN_NAME}-windows-386 main.go

unit:
	@go test ${PACKAGES}
