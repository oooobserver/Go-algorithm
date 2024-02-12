package binaryheap

import (
	"fmt"
	"testing"
)

func TestInsert(t *testing.T) {
	bh := New()
	bh.Insert(114)
	bh.Insert(514)
	bh.Insert(12)
	bh.Insert(311)
	fmt.Println(bh)
}
