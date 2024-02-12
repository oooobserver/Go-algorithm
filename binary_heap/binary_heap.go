package binaryheap

import "fmt"

type BinaryHeap struct {
	items []int
}

func New() BinaryHeap {
	return BinaryHeap{}
}

func (bh BinaryHeap) String() string {
	var res string
	for _, item := range bh.items {
		res += fmt.Sprintf("%d  ", item)
	}
	return res
}

func (bh *BinaryHeap) Insert(val int) {
	bh.items = append(bh.items, val)
	bh.heapifyUp(len(bh.items) - 1)
}

func (bh *BinaryHeap) DeleteMax() int {
	l := len(bh.items)
	if l == 0 {
		fmt.Println("The heap is empty")
		return 0
	}
	res := bh.items[0]
	bh.items[0], bh.items[l-1] = bh.items[l-1], bh.items[0]
	bh.items = bh.items[:l-1]
	bh.heapifyDown(0)
	return res
}

// Build a binary heap from the given array, note that the BH is implicit datastructure
// Note this only take O(n)
func (bh *BinaryHeap) Build(nums []int) {
	l := len(nums)

	// O(n)
	for i := 0; i < l; i++ {
		bh.items = append(bh.items, nums[i])
	}

	for i := l / 2; i >= 0; i-- { // O(n/2)
		bh.heapifyDown(i) // O(logn-logi)
	}
}

// O(logn), put the big value from bottem up
func (bh *BinaryHeap) heapifyUp(index int) {
	// Base case
	if index == 0 {
		return
	}

	parent := getParent(index)

	if bh.items[parent] < bh.items[index] {
		bh.items[parent], bh.items[index] = bh.items[index], bh.items[parent]
		bh.heapifyUp(parent)
	}
}

// O(logn), put the small value from top to bottom
func (bh *BinaryHeap) heapifyDown(index int) {
	maxIndex, maxVal := bh.maxChild(index)

	// Base case
	if maxIndex == -1 {
		return
	}
	if maxVal > bh.items[index] {
		bh.items[maxIndex], bh.items[index] = bh.items[index], bh.items[maxIndex]
		bh.heapifyDown(maxIndex)
	}
}

func getLeft(index int) int {
	return 2*index + 1
}

func getRight(index int) int {
	return 2*index + 2
}

func getParent(index int) int {
	return (index - 1) / 2
}

func (bh *BinaryHeap) maxChild(index int) (maxIndex int, maxVal int) {
	l, r := getLeft(index), getRight(index)
	// Base case
	if l >= len(bh.items) && r >= len(bh.items) {
		return -1, 0
	} else if r >= len(bh.items) {
		return l, bh.items[l]
	}

	if bh.items[l] < bh.items[r] {
		return r, bh.items[r]
	} else {
		return l, bh.items[l]
	}
}
