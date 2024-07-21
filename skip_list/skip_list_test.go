package skiplist

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestBasic(t *testing.T) {
	sl := NewSkipList()
	sl.Put(3, 3)
	sl.Put(6, 6)
	sl.Put(7, 7)
	sl.Put(1, 1)
	sl.Put(8, 8)
	sl.Put(9, 9)
	sl.Put(4, 4)

	// fmt.Println(sl)
	res, _ := sl.Get(6)
	require.Equal(t, res, 6)
	res, _ = sl.Get(8)
	require.Equal(t, res, 8)

	sl.Delete(6)
	_, ok := sl.Get(6)
	require.False(t, ok)
}

func TestLeetcode(t *testing.T) {
	sl := NewSkipList()
	sl.Put(3, 3)
	sl.Put(1, 1)
	sl.Put(2, 2)

	require.Equal(t, sl.String(), "[1/1]   [2/2]   [3/3]   \n")
	_, ok := sl.Get(0)
	require.False(t, ok)

	sl.Put(4, 4)
	res, ok := sl.Get(1)
	require.Equal(t, res, 1)
	require.True(t, ok)

	sl.Delete(0)
	sl.Delete(1)
	require.Equal(t, sl.String(), "[2/2]   [3/3]   [4/4]   \n")
	_, ok = sl.Get(1)
	require.False(t, ok)
}

func TestSkipListAll(t *testing.T) {
	l := NewSkipList()
	num := 100_000

	for i := 0; i < num; i++ {
		l.Put(i, i)
	}

	for i := 0; i < num; i++ {
		if i%2 == 0 {
			l.Delete(i)
		}
	}

	for i := 0; i < num; i++ {
		if i%2 != 0 {
			v, has := l.Get(i)
			assert.True(t, has)
			assert.Equal(t, i, v)
		}
	}
}

func TestSkipListIterator(t *testing.T) {
	l := NewSkipList()
	num := 100_000

	for i := 0; i < num; i++ {
		l.Put(i, i+1)
	}

	iter := l.NewSkipListIterator()
	for i := 0; i < num; i++ {
		assert.Equal(t, i, iter.key)
		assert.Equal(t, i+1, iter.value)
		iter.Next()
	}
}
