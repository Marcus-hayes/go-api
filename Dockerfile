# syntax=docker/dockerfile:1
##
## Build
##
FROM golang:1.16-alpine AS builder

WORKDIR /app

ENV HTTP_PORT=:8080
ENV JOKE_API=https://v2.jokeapi.dev/joke/
ENV MOVIE_API=http://www.omdbapi.com/?

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go ./

RUN go build -o /go-api
##
## Deploy
##
FROM gcr.io/distroless/base-debian10

WORKDIR /go-api

COPY --from=builder /go-api /go-api

EXPOSE 8080

USER nonroot:nonroot

ENTRYPOINT [ "/go-api" ]
