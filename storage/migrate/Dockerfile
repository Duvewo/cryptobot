FROM golang:1.16.4-buster

WORKDIR /migrate

RUN ["go", "get", "-u", "github.com/pressly/goose/cmd/goose"]

ENTRYPOINT ["goose", "up"]