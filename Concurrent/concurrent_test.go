package concurrent

import (
	"sync"
	"sync/atomic"
	"testing"
)

func corr_mutex(n int) {
	var counter int64
	var mu sync.Mutex

	var wg sync.WaitGroup
	wg.Add(n)

	for j := 0; j < n; j++ {
		go func() {
			defer wg.Done()
			for i := 0; i < 10000; i++ {
				mu.Lock()
				counter++
				mu.Unlock()
			}
		}()
	}
	wg.Wait()
}

func corr_atomic(n int) {
	var counter atomic.Int64

	var wg sync.WaitGroup
	wg.Add(n)

	for j := 0; j < n; j++ {
		go func() {
			defer wg.Done()
			for i := 0; i < 10000; i++ {
				counter.Add(1)
			}
		}()
	}
	wg.Wait()
}

func BenchmarkMutex(t *testing.B) {
	for i := 0; i < t.N; i++ {
		corr_mutex(5)
	}
}

func BenchmarkAtomic(t *testing.B) {
	for i := 0; i < t.N; i++ {
		corr_atomic(5)
	}
}
