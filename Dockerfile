FROM golang:1.16.4-buster

WORKDIR /src

COPY . .

RUN ["go", "mod", "download"]

RUN ["go", "build", "-o", "bot", "-mod", "vendor", "cmd/bot/main.go"]

ENTRYPOINT ["./bot"]