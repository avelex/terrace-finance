FROM golang:1.25-alpine3.22 AS builder

WORKDIR /build

COPY . .

RUN go build -o app cmd/main.go

FROM alpine:3.22
COPY --from=builder /build/app /build/app
COPY --from=builder /build/config.yaml /config.yaml
ENTRYPOINT ["/build/app"]