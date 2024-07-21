package benchmarks

import (
	"sync"
	"testing"
	"time"

	skiplist "shenye.com/skip_list"
)

func addKV(base int, num int, group int, l *skiplist.SCSkipList) {
	for i := 0; i < num; i++ {
		val := base + group*i
		l.Put(val, val)
	}
}

func BenchmarkSCSkipListWrite(b *testing.B) {
	l := skiplist.NewSCSkipList()
	group := 10
	num := 10000

	t := time.Now()
	var wg sync.WaitGroup
	wg.Add(group)
	for i := 1; i <= group; i++ {
		go func(i int) {
			defer wg.Done()
			addKV(i, num, group, l)
		}(i)
	}
	wg.Wait()
	interval := time.Since(t).Seconds()

	b.Log("Average write per second: ", float64(group*num)/interval)
}
