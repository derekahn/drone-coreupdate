FROM golang:1.12.7-alpine AS builder

WORKDIR /src
COPY ./go.mod ./
COPY ./src ./

RUN apk add --no-cache git && \
		go mod download && \
    GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o /go/bin/drone-coreupdate && \
    cd / && go get github.com/coreos/updateservicectl && \

    # Create the user and group files that will be used in the running container to
    # run the process an unprivileged user.
    mkdir /user && \
		echo 'nobody:x:65534:65534:nobody:/:' > /user/passwd && \
		echo 'nobody:x:65534:' > /user/group

FROM alpine:3.9.4 AS final

# Import the user and group files from the first stage.
COPY --from=builder /user/group /user/passwd /etc/
COPY --chown=nobody --from=builder /go/bin/ /bin/

RUN mkdir /lib64 && \
    ln -s /lib/libc.musl-x86_64.so.1 /lib64/ld-linux-x86-64.so.2 && \
    apk update && \
    apk --no-cache -Uuv add curl ca-certificates && \
    rm -rf /var/cache/apk/*


USER nobody:nobody

ENTRYPOINT ["/bin/drone-coreupdate"]
