PKG_NAME?=fabric-sdk-server
BUILD_DIR?=build

default: build

build:
	go build -o ${BUILD_DIR}/${PKG_NAME} cmd/server/main.go

build-all: build-linux build-darwin build-windows

build-linux:
	GOOS=linux   GOARCH=amd64 go build -o ${BUILD_DIR}/${PKG_NAME}.linux cmd/server/main.go

build-darwin:
	GOOS=darwin  GOARCH=amd64 go build -o ${BUILD_DIR}/${PKG_NAME}.darwin cmd/server/main.go
	
build-windows:
	GOOS=windows GOARCH=amd64 go build -o ${BUILD_DIR}/${PKG_NAME}.windows cmd/server/main.go

docker: build-linux
	docker build -t fabric-sdk-server:latest .

compose-up: build-linux
	docker-compose up --build

compose-down:
	docker-compose down