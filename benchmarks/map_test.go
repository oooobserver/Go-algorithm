package benchmarks

import (
	"sync"
	"testing"
	"time"
)

//
// This benchmark is intended to test the speed of
// Go std concurrent map and normal map with the RWMlock.
//

type CMap struct {
	m  map[int]int
	mu sync.RWMutex
}

func (c *CMap) put(k, v int) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.m[k] = v
}

func (c *CMap) get(k int) (int, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	v, ok := c.m[k]
	return v, ok
}

func BenchmarkCMapWrite(b *testing.B) {
	cm := CMap{m: make(map[int]int)}

	t := time.Now()
	for i := 0; i < 10_000; i++ {
		cm.put(i, i)
	}
	dur := time.Since(t)
	b.Log("Time consume: ", dur.Seconds())
}

func BenchmarkSyncMapWrite(b *testing.B) {
	var m sync.Map

	t := time.Now()
	for i := 0; i < 10_000; i++ {
		m.Store(i, i)
	}
	dur := time.Since(t)
	b.Log("Time consume: ", dur.Seconds())
}
