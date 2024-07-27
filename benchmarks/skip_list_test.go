package benchmarks

import (
	"fmt"
	"testing"
	"time"

	skiplist "shenye.com/skip_list"
)

func BenchmarkSkipListWrite(b *testing.B) {
	l := skiplist.NewSkipList()
	num := 100_000

	t := time.Now()
	for i := 0; i < num; i++ {
		l.Put(i, i)
	}
	interval := time.Since(t).Seconds()
	b.Log("Time consume: ", interval, "Average write per second: ", float64(num)/interval)
}

func BenchmarkSkipListRead(b *testing.B) {
	l := skiplist.NewSkipList()

	num := 100_000

	for i := 0; i < num; i++ {
		l.Put(i, i)
	}

	t := time.Now()
	// Test the correctness
	for i := 0; i < num; i++ {
		l.Get(i)
	}
	interval := time.Since(t).Seconds()
	fmt.Println("Average read per second: ", float64(num)/interval)
}

func BenchmarkSkipListDelete(b *testing.B) {
	l := skiplist.NewSkipList()

	num := 100_000
	for i := 0; i < num; i++ {
		l.Put(i, i)
	}

	t := time.Now()
	// Test the correctness
	for i := 0; i < num; i++ {
		l.Delete(i)
	}
	interval := time.Since(t).Seconds()
	fmt.Println("Average delete per second: ", float64(num)/interval)
}
