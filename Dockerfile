FROM golang:1.21 as builder

WORKDIR /app

COPY . .
RUN CGO_ENABLED=0 go build -o /app/tdarr_exporter cmd/tdarr_exporter/*.go

FROM alpine:3.6 as alpine

RUN apk add -U --no-cache ca-certificates

FROM scratch as app

WORKDIR /app

COPY --from=builder /app/tdarr_exporter /app/tdarr_exporter
COPY --from=alpine /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

ENTRYPOINT [ "/app/tdarr_exporter" ]