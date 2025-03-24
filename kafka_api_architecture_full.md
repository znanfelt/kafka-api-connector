
# Kafka to API Streaming Pipeline — Full Architecture Specification

This document describes the implementation-ready architecture for a production-grade Kafka-to-API pipeline written in Go. It includes precise technical instructions, interfaces, dependencies, and component interactions so that the system can be implemented from scratch or extended without ambiguity.

---

## High-Level Objective

Consume messages from a Kafka topic, transform and validate them, deduplicate in-memory, and POST them to an HTTP API with full observability, error handling, and enterprise-grade configuration and security.

---

## Components Overview

### 1. Kafka Consumer (franz-go)
- Library: [github.com/twmb/franz-go](https://github.com/twmb/franz-go)
- Secure connection via TLS and SASL/PLAIN
- Kafka config from `viper` (env/config.yaml)

**Settings Required**:
```yaml
kafka_brokers: localhost:9092
kafka_topic: events
kafka_group_id: api-consumer-group
kafka_username: file://secrets/kafka_user
kafka_password: file://secrets/kafka_pass
```

---

### 2. JSON Schema Validator
- Library: [github.com/xeipuuv/gojsonschema](https://pkg.go.dev/github.com/xeipuuv/gojsonschema)
- Schema path configured via `schema_path`
- Validate decoded JSON before transformation

---

### 3. Message Transformation
- Remove sensitive fields (e.g., `password`)
- Rename keys (e.g., `old_key` → `new_key`)
- Inject static tags (e.g., `"source": "redpanda"`)

**Transform Struct:**
```go
type Config struct {
  RenameFields   map[string]string
  RemoveFields   []string
  StaticTags     map[string]interface{}
}
```

---

### 4. Deduplication (Rotating Bloom Filter)
- Library: [github.com/bits-and-blooms/bloom/v3](https://github.com/bits-and-blooms/bloom)
- Structure: `current`, `prev`, with TTL and rotation
- Called every 30s via ticker
- Keys are Kafka message keys

**Implementation:**
```go
type RotatingBloomFilter struct {
  current *bloom.BloomFilter
  prev    *bloom.BloomFilter
  ttl     time.Duration
  ...
}
```

---

### 5. HTTP API Posting
- Use `http.Client` with 2s timeout
- Headers:
  - `Content-Type: application/json`
  - `Authorization: Bearer <token>`
- Token loaded from env or `file://`

**API Retry / Failover (Optional):**
```yaml
api_endpoint_primary: https://api1.example.com
api_endpoint_secondary: https://api2.example.com
```

---

### 6. DLQ (Dead Letter Queue)
- Kafka topic for messages that failed API POST
- Optional: Rate limiter (token bucket)
- Drop metric if limit exceeded

**DLQ Settings**:
```yaml
dlq_topic: failed-events
dlq_rate_limit: 100  # per second
dlq_capacity: 1000   # max burst
```

---

### 7. Prometheus Metrics
- Metrics exposed at `:2112/metrics`
- Library: [prometheus/client_golang](https://github.com/prometheus/client_golang)
- Counters:
  - `messages_consumed_total`
  - `messages_posted_total`
  - `messages_failed_total`
  - `dedup_hits_total`
  - `dedup_added_total`
  - `dedup_rotated_total`
  - `dlq_dropped_total`

---

### 8. Configuration Loader
- Library: [github.com/spf13/viper](https://github.com/spf13/viper)
- Fallback to env vars
- Supports `file://` secret paths

**Example**:
```yaml
api_auth_token: file://secrets/api_token
```

---

### 9. CLI Entrypoint
- Library: [github.com/spf13/cobra](https://github.com/spf13/cobra)
- Entrypoint: `cmd/consumer/main.go`
- Flags:
  - `--schema ./schema.json`
  - `--config ./config.yaml`

---

## Additional Operational Features

### Helm Deployment
- `values.yaml` for all config options
- Probe, replicaCount, resource requests/limits

### GitHub Actions
- Test, lint, coverage badge, Docker build, helm lint

### GitHub Pages Docs
- Markdown in `docs/`
- Mermaid diagram + architecture

---

## Testing Strategy

| Type        | Tool                 | Notes |
|-------------|----------------------|-------|
| Unit Tests  | Go test              | All packages |
| Integration | testcontainers-go    | Kafka, API, DLQ |
| Chaos       | Timeout simulation   | API latency |
| Shadow      | Mirror Kafka traffic | Secondary pipeline |
| Snapshot    | JSON payloads        | Validate transform |
| Contract    | JSON Schema          | Validation layer |
| Load        | `k6`                 | 20k+ TPS |

---

## Output Directory Structure

```
kafka_api_enterprise/
├── cmd/
├── internal/
├── pkg/
│   ├── api/
│   ├── kafka/
│   ├── dedup/
│   └── metrics/
├── docs/
├── charts/
├── config.yaml
├── Makefile
├── docker-compose.yaml
└── README.md
```

---

## Summary

This design is:
- Safe: Secure auth, DR ready, and observable
- Performant: Scales past 20k TPS
- Maintainable: Modular and CI-ready
- OSS-ready: GitHub Pages, Helm, Docs, Contribution guide



---

## Deduplication: Local Bloom Filter vs Redis

### Design Decision: Use In-Memory Rotating Bloom Filter

We opted for an in-memory rotating Bloom filter instead of Redis for message deduplication due to the following reasons:

---

### Local Bloom Filter (Chosen)

#### ✅ Pros:
- **No external dependency**: Simplifies deployment and local dev
- **Fast**: In-memory checks/adds are nanosecond-level
- **Low resource usage**: Compact memory footprint, configurable false-positive rate
- **Stateless horizontally**: Each instance independently filters keys
- **Well-suited for time-bound deduplication**

#### ⚠️ Cons:
- Not shared between instances (risk of duplicate delivery in multi-instance mode)
- Limited visibility into dedup state (not queryable externally)

---

### Redis-based Deduplication (Not Chosen)

#### ✅ Pros:
- Global key visibility across replicas
- Expiration-based dedup window
- Good for strict once-only delivery requirements

#### ⚠️ Cons:
- **Extra infra**: Requires provisioning, scaling, securing Redis
- **Latency**: Redis round trip can be 1-5ms vs local ~100ns
- **Throughput bottleneck**: 20k TPS could cause Redis saturation under load
- **Complexity**: Adds operational overhead for DR, TLS, HA

---

### Summary

For this pipeline, **fast, cost-effective, stateless deduplication** was prioritized over strict global enforcement.

Since dedup is bounded to short time windows (e.g. minutes), and occasional duplicates are tolerable, the rotating Bloom filter was ideal.

If global deduplication is required (e.g. for financial transactions or ledger systems), Redis or Kafka Streams may be more appropriate.

