package minheap

import "sort"

type Interface interface {
	sort.Interface // Len, Less, Swap
	Push(x any)    // add x as element Len()
	Pop() any      // remove and return the last element
}

func Init(h Interface) {
	// heapify
	n := h.Len()
	for i := n/2 - 1; i >= 0; i-- {
		down(h, i, n)
	}
}

// O(log n)
func Push(h Interface, x any) {
	h.Push(x)
	up(h, h.Len()-1)
}

// O(log n)
func Pop(h Interface) any {
	n := h.Len() - 1
	h.Swap(0, n)
	down(h, 0, n)
	return h.Pop()
}

// O(log n), removes and returns the element at index i
func Remove(h Interface, i int) any {
	n := h.Len() - 1
	if n != i { // If not the last element, should swap and re-establishes
		h.Swap(i, n)
		if !down(h, i, n) {
			up(h, i)
		}
	}
	return h.Pop()
}

// O(log n), re-establishes the heap ordering after the element at index i has changed its value.
func Fix(h Interface, i int) {
	if !down(h, i, h.Len()) {
		up(h, i)
	}
}

// O(log n)
func up(h Interface, j int) {
	for {
		parent := get_parent(j)
		if parent == j || !h.Less(j, parent) {
			break
		}
		h.Swap(parent, j)
		j = parent
	}
}

func get_parent(i int) int {
	return (i - 1) / 2
}

func get_left_child(i int) int {
	return 2*i + 1
}

func down(h Interface, i, length int) bool {
	cur := i
	for {
		lc := get_left_child(cur)
		if lc >= length || lc < 0 { // over length or int overflow
			break
		}

		// Get the min value of two childs
		min := lc
		if rc := lc + 1; rc < length && h.Less(rc, lc) {
			min = rc
		}

		// If the min value of child is greater than parent, stop
		if !h.Less(min, cur) {
			break
		}

		h.Swap(cur, min)
		cur = min
	}
	return cur > i
}

// Below is the example of how you use std heap
// type pair struct {
// 	f int
// 	c byte
// }

// type hp []pair

// func (h hp) Len() int           { return len(h) }
// func (h hp) Less(i, j int) bool { a, b := h[i], h[j]; return a.f < b.f || a.f == b.f && a.c < b.c }
// func (h hp) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
// func (hp) Push(any)             {}
// func (hp) Pop() (_ any)         { return }

// type hp[T comparable] []*node[T]

// func (h hp[T]) Len() int           { return len(h) }
// func (h hp[T]) Less(i, j int) bool { a, b := h[i], h[j]; return a.freq < b.freq || a.freq == b.freq }
// func (h *hp[T]) Swap(i, j int)     { (*h)[i], (*h)[j] = (*h)[j], (*h)[i] }
// func (h *hp[T]) Push(n *node[T])   { (*h) = append((*h), n) }
// func (h *hp[T]) Pop() (_ any)      { (*h) = (*h)[:len((*h))-1]; return }
