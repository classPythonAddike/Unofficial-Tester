# syntax=docker/dockerfile:1

FROM golang:1.16-alpine

WORKDIR /app

COPY . .
RUN go mod tidy

CMD [ "go", "run", "." ]