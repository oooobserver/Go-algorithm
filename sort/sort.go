package sort

import (
	"slices"

	binaryheap "shenye.com/binary_heap"
)

// Insertion sort, O(n^2)
func InsertionSort(nums []int) {
	// Base case
	if len(nums) < 2 {
		return
	}

	for i := 1; i < len(nums); i++ {
		temp := nums[i]
		j := i - 1
		for j >= 0 && nums[j] > temp {
			nums[j+1] = nums[j]
			j -= 1
		}
		nums[j+1] = temp
	}
}

// Merge sort, O(nlogn)
// T(n) = 2T(n/2) + f(n)
func MergeSort(nums []int, p ...int) {
	var a, b, c int

	// Default parameter
	if len(p) == 0 {
		a = 0
		b = len(nums)
	} else {
		a = p[0]
		b = p[1]
	}
	if b-a > 1 {
		c = (a + b + 1) / 2
		MergeSort(nums, a, c)
		MergeSort(nums, c, b)

		l, r := make([]int, c-a), make([]int, b-c)
		copy(l, nums[a:c])
		copy(r, nums[c:b])
		merge(nums, l, r, a, b)
	}
}

func merge(nums, left, right []int, start, end int) {
	i, j := 0, 0
	for ; start < end; start++ {
		if (j == len(right)) || (i < len(left) && (left[i] < right[j])) {
			nums[start] = left[i]
			i++
		} else {
			nums[start] = right[j]
			j++
		}
	}
}

// CountingSort suit for non-negative number, O(n)
func CountingSort(nums []int) {
	if len(nums) == 0 {
		return
	}

	// O(n) + O(n), find the min and max value
	minVal, maxVal := slices.Min(nums), slices.Max(nums)
	r := maxVal - minVal + 1

	// O(u)
	count := make([]int, r)

	// O(n)
	for _, num := range nums {
		count[num-minVal]++
	}

	i := 0
	for key, freq := range count {
		for j := 0; j < freq; j++ {
			nums[i] = key + minVal
			i++
		}
	}
}

// RadixSort, use base 10, O(n+logu*n)
func RadixSort(nums []int) {
	// O(n)
	maxNum := slices.Max(nums)

	exp := 1
	// O(logu)
	for maxNum/exp > 0 {
		radix_countingSort(nums, exp) // O(n)
		exp *= 10
	}
}

func radix_countingSort(arr []int, exp int) {
	n := len(arr)
	output := make([]int, n)
	count := make([]int, 10)

	// O(n)
	for i := 0; i < n; i++ {
		index := arr[i] / exp
		count[index%10]++
	}

	for i := 1; i < 10; i++ {
		count[i] += count[i-1]
	}

	// O(n)
	// The reason why reverse is to not break related order
	// Because in previous, put order is from last to least
	for i := n - 1; i >= 0; i-- {
		index := arr[i] / exp
		output[count[index%10]-1] = arr[i]
		count[index%10]--
	}

	// O(n)
	for i := 0; i < n; i++ {
		arr[i] = output[i]
	}
}

// O(nlogn)
func BinarySort(nums []int) {
	bh := binaryheap.New()
	bh.Build(nums)
	for i := len(nums) - 1; i >= 0; i-- {
		nums[i] = bh.DeleteMax()
	}
}
