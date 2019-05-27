FROM golang:1.11.0-alpine3.8
ADD build/fabric-sdk-server.linux /
ENTRYPOINT ["/fabric-sdk-server.linux"]