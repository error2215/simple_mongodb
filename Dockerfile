FROM golang:1.9.2 as builder
WORKDIR /go/src/simple_mongodb
#RUN go get -d -v github.com/gorilla/mux \
#	&& go get -d -v gopkg.in/mgo.v2/bson \
#	&& go get -d -v gopkg.in/mgo.v2
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .

FROM alpine:latest
#ARG SOURCE_LOCATION=/
RUN apk --no-cache add curl
EXPOSE 9090
WORKDIR /root/
COPY --from=builder /go/src/simple_mongodb .
CMD ["./app"]
