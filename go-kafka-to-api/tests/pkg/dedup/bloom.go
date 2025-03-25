package dedup

import (
    "sync"
    "time"

    "github.com/bits-and-blooms/bloom/v3"
)

type RotatingBloomFilter struct {
    current *bloom.BloomFilter
    prev    *bloom.BloomFilter
    ttl     time.Duration
    lock    sync.Mutex
    created time.Time
    n       uint
    fpRate  float64
}

func NewRotatingBloomFilter(ttl time.Duration, n uint, fpRate float64) *RotatingBloomFilter {
    return &RotatingBloomFilter{
        current: bloom.NewWithEstimates(n, fpRate),
        prev:    bloom.NewWithEstimates(n, fpRate),
        ttl:     ttl,
        created: time.Now(),
        n:       n,
        fpRate:  fpRate,
    }
}

func (r *RotatingBloomFilter) Add(data []byte) {
    r.lock.Lock()
    defer r.lock.Unlock()
    r.current.Add(data)
}

func (r *RotatingBloomFilter) Exists(data []byte) bool {
    r.lock.Lock()
    defer r.lock.Unlock()
    return r.current.Test(data) || r.prev.Test(data)
}

func (r *RotatingBloomFilter) Rotate() {
    r.lock.Lock()
    defer r.lock.Unlock()

    if time.Since(r.created) > r.ttl {
        r.prev = r.current
        r.current = bloom.NewWithEstimates(r.n, r.fpRate)
        r.created = time.Now()
    }
}
