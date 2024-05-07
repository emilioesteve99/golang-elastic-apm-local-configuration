FROM golang:1.20-alpine

RUN mkdir /app

ADD . /app

WORKDIR /app

RUN go build -o main pkg/main.go

CMD ["/app/main"]