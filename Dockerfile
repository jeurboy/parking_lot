FROM golang:1.13
RUN apt-get update && apt-get install -y vim
WORKDIR /usr/src/parking_lot

ENV GOPATH=/usr

RUN go get github.com/oxequa/realize

ENV GO111MODULE=on