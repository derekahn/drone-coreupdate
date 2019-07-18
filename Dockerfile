FROM golang:1.12.7-alpine AS builder

WORKDIR /src
COPY ./go.mod ./
COPY ./src ./

RUN apk add --no-cache git && \
		go mod download && \
    GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o /go/bin/drone-coreupdate && \
    cd / && go get github.com/coreos/updateservicectl && \


FROM alpine:3.9.4 AS final


RUN mkdir /lib64 && \
    ln -s /lib/libc.musl-x86_64.so.1 /lib64/ld-linux-x86-64.so.2 && \
    apk update && \
    apk --no-cache -Uuv add curl ca-certificates && \
    rm -rf /var/cache/apk/*



ENTRYPOINT ["/bin/drone-coreupdate"]
