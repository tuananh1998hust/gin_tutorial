FROM golang:1.12.3-alpine as build-env
WORKDIR /go/src/github.com/tuananh1998hust/gin_tutorial
COPY . .
RUN apk add git build-base && \
    go get -u -f -v . && \
    go build main.go

FROM alpine:3.10
WORKDIR /app
COPY --from=build-env /go/src/github.com/tuananh1998hust/gin_tutorial/main ./
EXPOSE 8080
CMD ["./main"]