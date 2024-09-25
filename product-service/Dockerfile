FROM golang:1.21 AS builder

ENV CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /go/src/app

COPY . .

RUN go build -mod vendor -o service ./cmd

FROM scratch

COPY --from=builder /go/src/app/service /service
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt

USER 58446

EXPOSE 8080

ENTRYPOINT ["/service"]