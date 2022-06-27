FROM golang:1.18.0-alpine3.15

WORKDIR /app
RUN apk update && apk add --no-cache ca-certificates curl && rm -rf /var/cache/apk/*

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY ./cmd ./cmd
COPY ./internal ./internal
RUN go build -o powerflex ./cmd

CMD ./powerflex