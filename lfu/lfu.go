package lfu

type node[T comparable] struct {
	freq int
	key  T
	val  any
}

type LFU[T comparable] struct {
	capacity int
	size     int
	cache    map[T]node[T]
	heap     hp[T]
}

type hp[T comparable] []*node[T]

func (h hp[T]) Len() int           { return len(h) }
func (h hp[T]) Less(i, j int) bool { a, b := h[i], h[j]; return a.freq < b.freq || a.freq == b.freq }
func (h *hp[T]) Swap(i, j int)     { (*h)[i], (*h)[j] = (*h)[j], (*h)[i] }
func (h *hp[T]) Push(n *node[T])   { (*h) = append((*h), n) }
func (h *hp[T]) Pop() (_ any)      { (*h) = (*h)[:len((*h))-1]; return }

func NewLFU[T comparable](capacity int) LFU[T] {
	lfu := LFU[T]{
		capacity: capacity,
		cache:    make(map[T]node[T]),
		heap:     make([]*node[T], 0),
	}

	return lfu
}

func (l *LFU[T]) Get(key T) (val any, ok bool) {
	if node, ok := l.cache[key]; ok {
		node.freq++
		return node.val, true
	}

	return nil, false
}
