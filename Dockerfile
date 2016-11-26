FROM golang:1.7

ADD . /go/src/github.com/zeusproject/zeus-server

WORKDIR /go/src/github.com/zeusproject/zeus-server

RUN go build

