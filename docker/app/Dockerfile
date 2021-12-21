FROM golang:1.16

COPY app/go.mod /go/src/app/go.mod
COPY app/go.sum /go/src/app/go.sum
WORKDIR /go/src/app
RUN go mod download
COPY ./app /go/src/app
EXPOSE 8080
