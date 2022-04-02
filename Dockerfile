# syntax=docker/dockerfile:1
##
## Build
##
FROM golang:1.16-alpine AS builder

WORKDIR /app

ENV HTTP_PORT=:8080
ENV JOKE_API=https://v2.jokeapi.dev/joke/
ENV MOVIE_API=http://www.omdbapi.com/?
ENV GO111MODULE=on

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o go-api

##
## Deploy
##
FROM alpine:3.14 as production
COPY --from=builder app .

EXPOSE 8080

ENTRYPOINT [ "/go-api" ]
