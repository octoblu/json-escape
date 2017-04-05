FROM golang:1.6
MAINTAINER Octoblu, Inc. <docker@octoblu.com>

WORKDIR /go/src/github.com/octoblu/json-escape
COPY . /go/src/github.com/octoblu/json-escape

RUN env CGO_ENABLED=0 go build -o json-escape -a -ldflags '-s' .

CMD ["./json-escape"]
