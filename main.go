package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func convert(nums []int) *ListNode {
	head := &ListNode{nums[0], nil}
	cur := head
	for i := 1; i < len(nums); i++ {
		tmp := &ListNode{nums[i], nil}
		cur.Next = tmp
		cur = tmp
	}

	return head
}

func print(l *ListNode) {
	cur := l
	for cur != nil {
		fmt.Printf("%d   ", cur.Val)
		cur = cur.Next
	}
}

func main() {
}
