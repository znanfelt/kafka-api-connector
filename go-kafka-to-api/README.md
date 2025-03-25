
// README.md
# Go Kafka Consumer

## Setup

1. Install dependencies:
```sh
go mod tidy
```

2. Run Kafka:
```sh
docker-compose up -d
```

3. Start consumer:
```sh
go run cmd/main.go
```

## Testing

Run unit tests:
```sh
go test ./...
```

Run tests with coverage:
```sh
go test ./... -cover
```
