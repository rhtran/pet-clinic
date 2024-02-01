# build executable binary
FROM golang:1.15.1-alpine as builder

COPY . $GOPATH/src/golang-petclinic-rest
WORKDIR $GOPATH/src/golang-petclinic-rest

RUN mkdir -p /config
#ADD config /config

# git
RUN apk add --no-cache git

# Install library dependencies
RUN go mod tidy

# build the binary
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o /go/bin/golang-petclinic-rest

# build a small image
FROM scratch

# copy static executable
COPY --from=builder /go/bin/golang-petclinic-rest  /go/bin/golang-petclinic-rest
COPY config/. /config


ENV PORT 8090
EXPOSE 8090

ENTRYPOINT ["/go/bin/golang-petclinic-rest"]
