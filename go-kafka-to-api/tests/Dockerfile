FROM golang:1.21 as builder
WORKDIR /app
COPY . .
RUN go build -o kafka-api-consumer ./cmd/consumer

FROM debian:bullseye
COPY --from=builder /app/kafka-api-consumer /usr/local/bin/
ENTRYPOINT ["/usr/local/bin/kafka-api-consumer"]
