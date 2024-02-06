package structure

import (
	"fmt"
	"testing"
)

func TestTree(t *testing.T) {
	bt := newTestBinaryTree()

	a := bt.InorderTrav()
	fmt.Println(a)
	fmt.Println(bt.root.item)
}
