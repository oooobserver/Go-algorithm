package linkedlist

import (
	"fmt"
	"testing"
)

func TestT(t *testing.T) {
	sl := newSkipList()
	sl.insert(3, "three")
	sl.insert(6, "six")
	sl.insert(7, "seven")
	sl.insert(1, "one")
	sl.insert(8, "eight")
	sl.insert(9, "nine")
	sl.insert(4, "four")

	sl.print()

	fmt.Println("Searching for key 6:", sl.search(6))
	fmt.Println("Searching for key 8:", sl.search(8))

	sl.delete(6)
	fmt.Println("After deleting key 6:")
	sl.print()
}
