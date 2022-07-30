FROM golang:1.18-alpine

RUN mkdir /app

WORKDIR /app

ADD . /app

RUN go build -o main .

CMD ["/app/main"]