package sort

import (
	"fmt"
	"math/rand"
	"slices"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

var rng = rand.New(rand.NewSource(time.Now().UnixNano()))

func TestInsertionSort(t *testing.T) {
	testArray := make([][]int, 10)
	for i := range testArray {
		size := rng.Intn(100)
		testArray[i] = make([]int, size)

		for j := 0; j < size; j++ {
			testArray[i][j] = rng.Int()
		}
	}

	for _, arr := range testArray {
		tmp := make([]int, len(arr))
		copy(tmp, arr)
		InsertionSort(tmp)
		slices.Sort(arr)
		require.Equal(t, tmp, arr)
	}
}

// func MergeSort1(nums []int, p ...int) {
// 	// Start, End, Middle
// 	var a, b, c int

// 	// Default parameter
// 	if len(p) == 0 {
// 		a = 0
// 		b = len(nums)
// 	} else {
// 		a = p[0]
// 		b = p[1]
// 	}
// 	if b-a > 1 {
// 		c = (a + b + 1) / 2
// 		MergeSort1(nums, a, c)
// 		MergeSort1(nums, c, b)

// 		merge1(nums, nums[a:c], nums[c:b], a, b)
// 	}
// }

// func merge1(nums, left, right []int, start, end int) {
// 	i, j := len(left)-1, len(right)-1
// 	for ; end > start; end-- {
// 		if (j == -1) || (i >= 0 && (left[i] > right[j])) {
// 			nums[end-1] = left[i]
// 			i--
// 		} else {
// 			nums[end-1] = right[j]
// 			j--
// 		}
// 	}
// }

func TestMergeSort(t *testing.T) {
	testArray := make([][]int, 10)
	for i := range testArray {
		size := rng.Intn(100)
		testArray[i] = make([]int, size)

		for j := 0; j < size; j++ {
			testArray[i][j] = rng.Int()
		}
	}

	for _, arr := range testArray {
		tmp := make([]int, len(arr))
		copy(tmp, arr)
		MergeSort(tmp)
		slices.Sort(arr)
		require.Equal(t, tmp, arr)
	}
}

func TestCountingSort(t *testing.T) {
	testArray := make([][]int, 10)
	for i := range testArray {
		size := rng.Intn(100)
		testArray[i] = make([]int, size)

		for j := 0; j < size; j++ {
			testArray[i][j] = rng.Intn(100)
		}
	}

	for _, arr := range testArray {
		tmp := make([]int, len(arr))
		copy(tmp, arr)
		CountingSort(tmp)
		slices.Sort(arr)
		require.Equal(t, tmp, arr)
	}
}

func TestRadixSort(t *testing.T) {
	testArray := make([][]int, 10)
	for i := range testArray {
		// Avoid empty list
		size := rng.Intn(100) + 1
		testArray[i] = make([]int, size)

		for j := 0; j < size; j++ {
			testArray[i][j] = rng.Int()
		}
	}

	for _, arr := range testArray {
		tmp := make([]int, len(arr))
		copy(tmp, arr)
		RadixSort(tmp)
		slices.Sort(arr)
		require.Equal(t, tmp, arr)
	}
}

func TestBinarySort(t *testing.T) {
	testArray := make([][]int, 10)
	for i := range testArray {
		// Avoid empty list
		size := rng.Intn(100) + 1
		testArray[i] = make([]int, size)

		for j := 0; j < size; j++ {
			testArray[i][j] = rng.Int()
		}
	}

	for _, arr := range testArray {
		tmp := make([]int, len(arr))
		copy(tmp, arr)
		BinarySort(tmp)
		slices.Sort(arr)
		require.Equal(t, tmp, arr)
	}
}

func BenchmarkSort(t *testing.B) {
	nums := []int{1, 14, -5, 7, 98, 116}
	for i := 0; i < t.N; i++ {
		InsertionSort(nums)
	}
}

func ExampleInsertionSort() {
	nums := []int{1, 14, -5, 7, 98, 116}
	InsertionSort(nums)
	fmt.Println(nums)
	// Output:
	// [-5 1 7 14 98 116]
}
