FROM golang:1.23.0-alpine

RUN mkdir /app

WORKDIR /app

ADD . /app

RUN go build -o main .

CMD ["/app/main"]
