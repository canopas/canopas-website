# Start from a Debian image with the latest version of Go installed
FROM golang:1.18 as builder

ADD . /go-api-platform

WORKDIR "/go-api-platform"

RUN CGO_ENABLED=0 go build -a -installsuffix cgo -ldflags '-s' -o go-api-platform

FROM alpine

RUN apk update && apk add ca-certificates

COPY --from=builder /go-api-platform/go-api-platform /

# Document that the service listens on port 8080.
EXPOSE 8080

ENTRYPOINT ["/go-api-platform"]