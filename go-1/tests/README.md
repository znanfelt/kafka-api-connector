```

██████╗ ██╗   ██╗ █████╗  ██████╗██╗  ██╗ █████╗ ██████╗ 
██╔══██╗██║   ██║██╔══██╗██╔════╝██║ ██╔╝██╔══██╗██╔══██╗
██████╔╝██║   ██║███████║██║     █████╔╝ ███████║██████╔╝
██╔═══╝ ██║   ██║██╔══██║██║     ██╔═██╗ ██╔══██║██╔═══╝ 
██║     ╚██████╔╝██║  ██║╚██████╗██║  ██╗██║  ██║██║     
╚═╝      ╚═════╝ ╚═╝  ╚═╝ ╚═════╝╚═╝  ╚═╝╚═╝  ╚═╝╚═╝     
                                                        
        Kafka → Dedup → Transform → REST API

```

# Kafka API Enterprise

![Test](https://img.shields.io/badge/tests-passing-brightgreen)
![Coverage](https://img.shields.io/badge/coverage-100%25-blue)
![Go Version](https://img.shields.io/badge/go-1.21-brightgreen)

A production-ready, enterprise-grade Kafka consumer pipeline in Go. It consumes JSON messages from Kafka, applies transformations and filtering, deduplicates using a rotating Bloom filter, and POSTs them to a downstream REST API. Includes full test coverage, observability, secure secret handling, and integration testing with Testcontainers.

---

## 🚀 Features

- ✅ High-throughput Kafka consumer (`franz-go`)
- ✅ JSON transformation, filtering, and schema validation
- ✅ Rotating Bloom filter deduplication
- ✅ DLQ support for failed API calls
- ✅ Prometheus metrics (`/metrics`)
- ✅ Secure secret handling (file-based)
- ✅ Testcontainers integration tests
- ✅ Chaos tests and contract validation
- ✅ Helm chart and Docker support

---

## 🧪 Running Tests

```bash
go mod tidy
go test ./... -coverprofile=coverage.out
go tool cover -html=coverage.out
```

For integration tests (requires Docker):
```bash
go test -tags=integration ./cmd/consumer
```

---

## 🐳 Running Locally with Docker Compose

```bash
docker-compose up --build
```

---

## ⚙️ Configuration

| Env Var | Description |
|---------|-------------|
| `KAFKA_BROKERS` | Comma-separated list of Kafka brokers |
| `KAFKA_TOPIC`   | Kafka topic to consume |
| `KAFKA_GROUP`   | Consumer group ID |
| `API_URL`       | Downstream API URL |
| `API_AUTH_TOKEN`| API bearer token (or `file://...`) |
| `KAFKA_USERNAME`| Kafka SASL username (or `file://...`) |
| `KAFKA_PASSWORD`| Kafka SASL password (or `file://...`) |

---

## 📦 Building

```bash
go build -o kafka-api-consumer ./cmd/consumer
```

---

## 📈 Observability

Visit [http://localhost:2112/metrics](http://localhost:2112/metrics) for Prometheus metrics.

---

## 🧪 Test Coverage Goals

| Module     | Coverage |
|------------|----------|
| API Client | ✅ 100% |
| Dedup      | ✅ 100% |
| Kafka      | ✅ DLQ fallback, secure client |
| Config     | ✅ Dynamic secret loading |
| Schema     | ✅ Snapshot validation |
| Chaos      | ✅ Timeout simulation |

---

## 📄 License

MIT © Your Company
