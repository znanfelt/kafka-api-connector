package dedup

import (
	"testing"
	"time"

	"github.com/bits-and-blooms/bloom/v3"
)

func TestRotatingBloomFilter(t *testing.T) {
	ttl := 10 * time.Millisecond
	f := NewRotatingBloomFilter(ttl, 1000, 5.0)

	key := []byte("duplicate-key")

	if f.Exists(key) {
		t.Error("Expected key to not exist initially")
	}

	f.Add(key)

	if !f.Exists(key) {
		t.Error("Expected key to exist after adding")
	}

	f.created = time.Now().Add(-20 * time.Millisecond)
	f.Rotate()

	if f.current.Test(key) {
		t.Error("Key should not exist in current after rotation")
	}

	f.prev = bloom.NewWithEstimates(1000, 0.01)

	if f.Exists(key) {
		t.Error("Expected key to not exist after rotation and clearing prev")
	}
}

func TestMultipleAddExist(t *testing.T) {
	f := NewRotatingBloomFilter(1*time.Minute, 10000, 0.01)

	keys := [][]byte{
		[]byte("key1"),
		[]byte("key2"),
		[]byte("key3"),
	}

	for _, k := range keys {
		if f.Exists(k) {
			t.Errorf("Key %s should not exist", k)
		}
		f.Add(k)
		if !f.Exists(k) {
			t.Errorf("Key %s should now exist", k)
		}
	}
}

func TestEmptyKeyHandling(t *testing.T) {
	f := NewRotatingBloomFilter(1*time.Minute, 1000, 0.01)

	empty := []byte("")
	if f.Exists(empty) {
		t.Error("Expected empty key to not exist")
	}

	f.Add(empty)
	if !f.Exists(empty) {
		t.Error("Expected empty key to exist after add")
	}
}

func TestFalsePositiveRate(t *testing.T) {
	f := NewRotatingBloomFilter(1*time.Minute, 10000, 0.01)

	added := make(map[string]struct{})
	for i := 0; i < 5000; i++ {
		k := []byte("real-" + string(rune(i)))
		f.Add(k)
		added[string(k)] = struct{}{}
	}

	falsePositives := 0
	for i := 0; i < 5000; i++ {
		k := []byte("fake-" + string(rune(i)))
		if _, exists := added[string(k)]; !exists && f.Exists(k) {
			falsePositives++
		}
	}

	if falsePositives > 100 { // tolerate some
		t.Errorf("Too many false positives: %d", falsePositives)
	}
}

func TestMultipleRotations(t *testing.T) {
	f := NewRotatingBloomFilter(50*time.Millisecond, 1000, 0.01)

	key := []byte("rotate-me")

	f.Add(key)
	if !f.Exists(key) {
		t.Error("Key should exist after add")
	}

	time.Sleep(60 * time.Millisecond)
	f.Rotate()
	time.Sleep(60 * time.Millisecond)
	f.Rotate()

	if f.Exists(key) {
		t.Error("Key should not exist after two rotations")
	}
}
