PKG_NAME?=fabric-sdk-server
BUILD_DIR?=build

default: build

build:
	go build -o ${BUILD_DIR}/${PKG_NAME} cmd/server/main.go

build-all: build-linux build-darwin build-windows

build-linux:
	GO111MODULE=on GOOS=linux GOARCH=amd64 go build -o ${BUILD_DIR}/${PKG_NAME}.linux cmd/server/main.go

build-darwin:
	GO111MODULE=on GOOS=darwin GOARCH=amd64 go build -o ${BUILD_DIR}/${PKG_NAME}.darwin cmd/server/main.go
	
build-windows:
	GO111MODULE=on GOOS=windows GOARCH=amd64 go build -o ${BUILD_DIR}/${PKG_NAME}.windows cmd/server/main.go

docker: build-linux
	docker build -t fabric-sdk-server:latest .

compose-up: build-linux
	docker-compose up --build

compose-down:
	docker-compose down

example-one: build-linux
	docker-compose -f examples/one-org/docker-compose.yml up --build 

example-one-down:
	docker-compose -f examples/one-org/docker-compose.yml down

example-two: build-linux
	docker-compose -f examples-two-orgs/docker-compose.yml up --build 

example-two-down:
	docker-compose -f examples-two-orgs/docker-compose.yml down
