package tree

import (
	"fmt"
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

func TestTreeTrav(t *testing.T) {
	bt := newTestBST()
	r := bt.inorderTravFor()
	fmt.Println(r)
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

func TestBinaryTravelPro(t *testing.T) {
	bt := newTestBST()
	bt.PreOrderDisplay()
	fmt.Println(bt.preorderPro())

	require.Equal(t, bt.inorderTravFor(), bt.inorderPro())

	bt.PostOrderDisplay()
	fmt.Println(bt.postorderPro())
}
