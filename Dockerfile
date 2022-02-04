FROM golang:1.16.13

RUN mkdir -p /go/src/github.com/beemensameh/
WORKDIR /go/src/github.com/beemensameh/

ADD . /go/src/github.com/beemensameh/
