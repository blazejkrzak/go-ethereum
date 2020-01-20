# Build Geth in a stock Go builder container
FROM golang:1.13-alpine as builder

RUN apk add --no-cache make gcc musl-dev linux-headers git

ADD . /go-ethereum
RUN cd /go-ethereum && make boot-and-geth

# Pull Geth into a second stage deploy alpine container
FROM alpine:latest

RUN apk add --no-cache ca-certificates
COPY --from=builder /go-ethereum/build/bin/geth /usr/local/bin/ssc
COPY --from=builder /go-ethereum/build/bin/bootnode /usr/local/bin/bootnode

EXPOSE 8545 8546 8547 30303 30303/udp
CMD ["ssc"]
