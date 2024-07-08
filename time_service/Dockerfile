# syntax=docker/dockerfile:1

## Build
FROM golang:1.19-alpine AS build

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

COPY . ./

RUN go mod download

RUN go build -o /time-service ./time_service/server/server.go

## Deploy
FROM alpine:latest

WORKDIR /

COPY --from=build /time-service /time-service
EXPOSE 8081
ENTRYPOINT ["/time-service"]