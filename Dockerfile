FROM golang:1.10.1-alpine3.7 AS builder

COPY . src/github.com/chadgrant/ethclient/

RUN apk update && apk add git build-base

RUN cd src/github.com/chadgrant/ethclient/ && \
    go get ./... && \
    go build

FROM alpine:3.7
RUN apk update && apk add ca-certificates
COPY --from=builder /go/src/github.com/chadgrant/ethclient/ethclient /
COPY www /www
ENTRYPOINT ["./ethclient"]