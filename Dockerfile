FROM golang:1.16

COPY ./app /go/src/app
WORKDIR /go/src/app

ENV GO111MODULE=on

RUN go mod download github.com/gin-gonic/gin

EXPOSE 8080
