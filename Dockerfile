FROM golang:latest

WORKDIR /GOLang
COPY . .


RUN go build .


EXPOSE 8080

CMD [ "./GOLang" ]
