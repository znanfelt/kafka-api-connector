package metrics

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	MessagesConsumed = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "messages_consumed_total",
			Help: "Total number of Kafka messages consumed",
		})

	MessagesPosted = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "messages_posted_total",
			Help: "Total number of messages successfully posted to the API",
		})

	MessagesFailed = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "messages_failed_total",
			Help: "Total number of failed message post attempts",
		})

	DedupHits = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "dedup_hits_total",
			Help: "Number of duplicate messages skipped due to deduplication",
		})

	DedupAdded = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "dedup_added_total",
			Help: "Number of unique messages added to the deduplication filter",
		})

	DedupRotated = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "dedup_rotated_total",
			Help: "Number of times the deduplication filter rotated",
		})
)

func init() {
	prometheus.MustRegister(MessagesConsumed, MessagesPosted, MessagesFailed,
		DedupHits, DedupAdded, DedupRotated)
}

func Handler() http.Handler {
	return promhttp.Handler()
}
