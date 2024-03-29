package hashtable

import "shenye.com/util"

// For now, only support these types of the key: int, string, rune
type tuple struct {
	key   interface{}
	value interface{}
}

type hash struct {
	hash_function func(interface{}) int
	store         [][]tuple
}

const hash_length = 100

func New() hash {
	var h hash

	a := int(util.RandomInt(0, 114))
	b := int(util.RandomInt(0, 114))

	h.hash_function = func(n interface{}) int {
		value := 0
		switch v := n.(type) {
		case int:
			value = v
		case rune:
			value = int(v)
		case string:
			tmp := 0
			for _, s := range v {
				tmp = tmp*10 + int(s)
			}
		}
		return (((a*value + b) % 113) % hash_length)
	}
	h.store = make([][]tuple, hash_length)

	return h
}

// func (h *hash) Set(key, value int) {
// 	index := h.hash_function(key)
// 	if h.store[index] == nil {
// 		h.store[index] = make([]Tuple, 5)
// 		h.store[index][0] = Tuple{key, value}
// 	} else {
// 		for i := 0; i < 5; i++ {
// 			if h.store[index][i].Value == 0 {
// 				h.store[index][i] = Tuple{key, value}
// 				break
// 			}
// 		}
// 	}
// }

// func (h *hash) Get(key int) int {
// 	index := h.hash_function(key)
// 	for _, t := range h.store[index] {
// 		if t.Key == key {
// 			return t.Value
// 		}
// 	}

// 	return -1
// }
