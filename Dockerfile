FROM golang:1.15.0-alpine3.12 AS builder

WORKDIR /app

COPY . .

RUN apk add --no-cache git && \
    go mod download && \
    GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o /go/bin/drone-coreupdate && \
    cd / && go get github.com/coreos/updateservicectl


FROM alpine:3.12 AS final

RUN mkdir /lib64 && \
    ln -s /lib/libc.musl-x86_64.so.1 /lib64/ld-linux-x86-64.so.2 && \
    apk update && \
    apk --no-cache -Uuv add curl ca-certificates && \
    rm -rf /var/cache/apk/*

COPY --from=builder /go/bin/ /bin/

HEALTHCHECK --interval=5s --timeout=10s --retries=3 \
  CMD [[ $(/bin/updateservicectl --version) =~ "1.4.0" ]] || exit 1

ENTRYPOINT ["/bin/drone-coreupdate"]
