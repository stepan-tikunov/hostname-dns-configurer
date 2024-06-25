FROM golang:latest AS base

WORKDIR /app

RUN go install github.com/air-verse/air@latest && go install github.com/go-delve/delve/cmd/dlv@master

ENTRYPOINT ["air"]
