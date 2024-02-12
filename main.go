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

func reverseKGroup(head *ListNode, k int) *ListNode {
	tmp := &ListNode{}
	tmp.Next = head
	cur := tmp

	for cur != nil {
		test := cur
		for i := 0; i < k; i++ {
			test = test.Next
			if test == nil {
				return tmp.Next
			}
		}

		cur.Next = reverse(cur.Next, k)

		for i := 0; i < k; i++ {
			cur = cur.Next
		}
	}

	return tmp.Next
}

func reverse(node *ListNode, target int) *ListNode {
	var prev, next, start *ListNode
	start = node
	for i := 0; i < target; i++ {
		next = node.Next
		node.Next = prev
		prev = node
		node = next
	}

	start.Next = next

	return prev
}

func main() {
}
