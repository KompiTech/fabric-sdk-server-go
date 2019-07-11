@ECHO OFF

pushd %~dp0

set PKGNAME=fabric-sdk-server
set BUILDDIR=build

if "%1" == "build-linux" goto build-linux
if "%1" == "build-darwin" goto build-darwin
if "%1" == "build-windows" goto build-windows
if "%1" == "docker" goto docker
if "%1" == "compose-up" goto compose-up
if "%1" == "compose-down" goto compose-down else goto build

:build
go build -o %BUILD_DIR%/%PKG_NAME% cmd/server/main.go
goto end

:build-linux
GOOS=linux GOARCH=amd64 go build -o %BUILD_DIR%/%PKG_NAME%.linux cmd/server/main.go
goto end

:build-darwin
GOOS=darwin GOARCH=amd64 go build -o %BUILD_DIR%/%PKG_NAME%.darwin cmd/server/main.go
goto end

:build-windows
GOOS=windows GOARCH=amd64 go build -o %BUILD_DIR%/%PKG_NAME%.windows cmd/server/main.go
goto end

:docker
GOOS=linux GOARCH=amd64 go build -o %BUILD_DIR%/%PKG_NAME%.linux cmd/server/main.go
docker build -t fabric-sdk-server:latest .
goto end

:compose-up
GOOS=linux GOARCH=amd64 go build -o %BUILD_DIR%/%PKG_NAME%.linux cmd/server/main.go
docker-compose up --build
goto end

:compose-down
docker-compose down

:end
popd
