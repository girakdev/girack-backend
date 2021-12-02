FROM golang:1.16

COPY ./app /go/src/app
WORKDIR /go/src/app

ENV GO111MODULE=on

EXPOSE 8080
