package benchmarks

import (
	"sync"
	"testing"
	"time"
)

type TestMap struct {
	items map[int]any
	mu    sync.RWMutex
}

func (t *TestMap) SetDefer(a int, b any) {
	t.mu.Lock()
	defer t.mu.Unlock()
	t.items[a] = b
}

func (t *TestMap) Set(a int, b any) {
	t.mu.Lock()
	t.items[a] = b
	t.mu.Unlock()
}

func BenchmarkTestDefer(b *testing.B) {
	v := "value"
	t := time.Now()
	m := TestMap{items: make(map[int]any)}
	for i := 0; i < 10_000; i++ {
		m.SetDefer(i, v)
	}
	dur := time.Since(t)
	b.Log(dur.Abs().Seconds())
}

func BenchmarkTestSet(b *testing.B) {
	v := "value"
	t := time.Now()
	m := TestMap{items: make(map[int]any)}
	for i := 0; i < 10_000; i++ {
		m.Set(i, v)
	}
	dur := time.Since(t)
	b.Log(dur.Abs().Seconds())
}
