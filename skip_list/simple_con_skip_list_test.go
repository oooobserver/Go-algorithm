package skiplist

import (
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

func addKV(base int, num int, group int, l *SCSkipList) {
	for i := 0; i < num; i++ {
		val := base + group*i
		l.Put(val, val)
	}
}

func TestPutCorrectness(t *testing.T) {
	l := NewSCSkipList()
	group := 10
	num := 1000

	var wg sync.WaitGroup
	wg.Add(group)
	for i := 1; i <= group; i++ {
		go func(i int) {
			defer wg.Done()
			addKV(i, num, group, l)
		}(i)
	}
	wg.Wait()

	for i := 1; i <= group*num; i++ {
		v, has := l.Get(i)
		assert.True(t, has)
		assert.Equal(t, i, v)
	}

	_, has := l.Get(0)
	assert.False(t, has)
}
