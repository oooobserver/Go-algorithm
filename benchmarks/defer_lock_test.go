package benchmarks

import "sync"

type TestMap struct {
	items map[string]any
	mu    sync.RWMutex
}

func (t *TestMap) SetDefer(a string, b any) {
	t.mu.Lock()
	t.items[a] = b
	t.mu.Unlock()
}

func BenchmarkTestDefer() {

}
