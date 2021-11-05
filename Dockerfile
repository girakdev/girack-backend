FROM golang:1.15-alpine3.12

RUN apk update && \
    apk upgrade && \
    apk add git

ENV CGO_ENABLED=0

WORKDIR /go/src/app
COPY go.mod go.sum main.go ./

RUN go mod download

EXPOSE 8080
