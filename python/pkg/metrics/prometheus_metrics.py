from prometheus_client import Counter, start_http_server

# Kafka Metrics
messages_consumed_total = Counter("messages_consumed_total", "Total messages consumed from Kafka")
messages_posted_total = Counter("messages_posted_total", "Messages successfully posted to API")
messages_failed_total = Counter("messages_failed_total", "Messages failed to post to API")

# Deduplication Metrics
dedup_hits_total = Counter("dedup_hits_total", "Duplicate messages skipped")
dedup_added_total = Counter("dedup_added_total", "New messages added to dedup filter")
dedup_rotated_total = Counter("dedup_rotated_total", "Bloom filters rotated")

# DLQ Limiting
dlq_limiter_drops_total = Counter("dlq_limiter_drops_total", "Messages dropped due to DLQ rate limiting")

def start_metrics_server(port: int = 2112):
    start_http_server(port)
    print(f"Prometheus metrics server running on :{port}/metrics")
