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

		min := lc                                        // left child
		if rc := lc + 1; rc < length && h.Less(rc, lc) { // right child exit and right <= left
			min = rc // right child
		}

		// The min value of child is greater than parent, stop
		if !h.Less(min, cur) {
			break
		}

		h.Swap(cur, min)
		cur = min
	}
	return cur > i
}
