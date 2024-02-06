package structure

import (
	"math/rand"
	"time"
)

type Tuple struct {
	Key   int
	Value int
}

type hash struct {
	hash_function func(int) int
	store         [][]Tuple
}

func NewHash() hash {
	var h hash
	var rng = rand.New(rand.NewSource(time.Now().UnixNano()))
	a := rng.Intn(113)
	b := rng.Intn(113)
	h.hash_function = func(n int) int {
		return (((a*n + b) % 113) % 50)
	}
	h.store = make([][]Tuple, 50)

	return h
}

func (h *hash) Hash_set(key, value int) {
	index := h.hash_function(key)
	if h.store[index] == nil {
		h.store[index] = make([]Tuple, 5)
		h.store[index][0] = Tuple{key, value}
	} else {
		for i := 0; i < 5; i++ {
			if h.store[index][i].Value == 0 {
				h.store[index][i] = Tuple{key, value}
				break
			}
		}
	}
}

func (h *hash) Hash_get(key int) int {
	index := h.hash_function(key)
	for _, t := range h.store[index] {
		if t.Key == key {
			return t.Value
		}
	}

	return -1
}
