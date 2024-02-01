package algorithm

// Insertion sort
func InsertionSort(nums []int) []int {
	// Base case
	if len(nums) < 2 {
		return nums
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
	return nums
}

// Merge sort
func MergeSort(nums []int, p ...int) {
	var a, b, c int

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
