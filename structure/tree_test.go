package structure

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestBSTDelete(t *testing.T) {
	bt := newTestBST()
	bt.BinaryDelete(4)
	bt.BinaryDelete(2)
	bt.BinaryDelete(1)
	res := []int{3, 5, 6, 7, 8}
	test_res := bt.InorderTrav()
	require.Equal(t, res, test_res)
}

func TestBSTInsert(t *testing.T) {
	bt := newTestBST()
	biggest := BinarySubTreeLast(bt.root)
	res := append(bt.InorderTrav(), biggest.item+1)
	bt.InsertAfter(biggest, biggest.item+1)
	test_res := bt.InorderTrav()
	require.Equal(t, res, test_res)
}

func TestBinarySearch(t *testing.T) {
	bt := newTestBST()
	require.True(t, bt.BinarySearch(1))
	require.True(t, bt.BinarySearch(4))
	require.True(t, bt.BinarySearch(7))
	require.False(t, bt.BinarySearch(9))
}
