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

func search(nums []int, target int) int {
	l, r := 0, len(nums)-1

	for l <= r {
		mid := (l + r + 1) / 2
		if nums[mid] == target {
			return mid
		}

		if nums[l] < nums[mid] {
			if target > nums[mid] || target < nums[l] {
				l = mid + 1
			} else {
				r = mid - 1
			}
		}

		if nums[mid] < nums[r] {
			if target > nums[r] || target < nums[mid] {
				r = mid - 1
			} else {
				l = mid + 1
			}
		}
	}
	return -1
}

func main() {
	nums := []int{4, 5, 6, 7, 0, 1, 2}
	r := search(nums, 0)
	fmt.Println(r)
}
