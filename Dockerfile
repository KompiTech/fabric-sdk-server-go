FROM golang:1.12.6-stretch
ADD build/fabric-sdk-server.linux /
CMD ["/fabric-sdk-server.linux"]