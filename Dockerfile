FROM golang:1.11.2 as builder
WORKDIR /go/src/github.com/error2215/simple_mongodb
COPY . .

RUN go build -o main .

EXPOSE 3034

CMD ["./main"]
