package leetlist

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func Print(n *ListNode) {
	for n != nil {
		fmt.Printf("%d   ", n.Val)
		n = n.Next
	}
	fmt.Println()
}

func AddNode(n *ListNode, val int) {
	next := n.Next
	n.Next = &ListNode{Val: val, Next: next}
}

func GenTestList() *ListNode {
	h := &ListNode{0, nil}
	for i := range 5 {
		AddNode(h, i)
	}
	return h
}
