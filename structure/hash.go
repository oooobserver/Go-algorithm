package structure

import (
	"math/rand"
	"time"
)

type hash_tuple struct {
	key   int
	value int
}

type hash struct {
	hash_function func(int) int
	store         [][]hash_tuple
}

func NewHash() hash {
	var h hash
	var rng = rand.New(rand.NewSource(time.Now().UnixNano()))
	a := rng.Intn(113)
	b := rng.Intn(113)
	h.hash_function = func(n int) int {
		return (((a*n + b) % 113) % 50)
	}
	h.store = make([][]hash_tuple, 50)

	return h
}

func (h *hash) Hash_set(key, value int) {
	index := h.hash_function(key)
	if h.store[index] == nil {
		h.store[index] = make([]hash_tuple, 5)
		h.store[index][0] = hash_tuple{key, value}
	} else {
		for i := 0; i < 5; i++ {
			if h.store[index][i].value == 0 {
				h.store[index][i] = hash_tuple{key, value}
				break
			}
		}
	}
}

func (h *hash) Hash_get(key int) int {
	index := h.hash_function(key)
	for _, t := range h.store[index] {
		if t.key == key {
			return t.value
		}
	}

	return -1
}
