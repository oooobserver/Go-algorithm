package sort

import (
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
