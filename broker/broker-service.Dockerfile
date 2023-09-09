FROM golang:1.20-alpine as builder

WORKDIR /build

COPY . .

RUN CGO_ENABLED=0 go build -o brokerApp ./cmd/api

FROM alpine:latest

WORKDIR /app

COPY --from=builder /build/brokerApp .

CMD ["./brokerApp"]