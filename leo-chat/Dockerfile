# syntax=docker/dockerfile:1

FROM golang:1.17-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY . .

ENV CHATPORT=2020

RUN go build -o /docker-leo

EXPOSE 8080

CMD [ "/docker-leo" ]