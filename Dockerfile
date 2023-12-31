FROM golang:alpine

LABEL maintainer="Turdalin Nurassyl"

RUN apk update && apk add --no-cache git && apk add --no-cache bash && apk add build-base

WORKDIR /app

COPY . .
COPY .env ./.env

RUN go get -d -v ./...
RUN go install -v ./...

RUN go install -mod=mod github.com/githubnemo/CompileDaemon
RUN go get -v golang.org/x/tools/gopls

ENTRYPOINT CompileDaemon --build="go build -a -installsuffix cgo -o main ./cmd/api" --command=./main