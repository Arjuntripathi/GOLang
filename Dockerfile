# syntax=docker/dockerfile:1

FROM golang:latest

WORKDIR /app

RUN git clone https://github.com/Arjuntripathi/GOLang

RUN cd /GOLang && go build

EXPOSE 8080

CMD ['./GOLang']

