package linkedlist

import (
	"fmt"
)

type Node struct {
	Value int
	Next  *Node
}

type LinkedList struct {
	Head *Node
}

func New() LinkedList {
	return LinkedList{}
}

// Keep the order while appending, O(n)
func (l *LinkedList) OrderAppend(v int) {
	newNode := &Node{v, nil}

	if l.Head == nil {
		l.Head = newNode
		return
	}

	current := l.Head
	for current.Next != nil {
		current = current.Next
	}

	current.Next = newNode
}

// More efficient, but break the order, O(1)
func (l *LinkedList) Append(v int) {
	newNode := &Node{v, nil}
	if l.Head == nil {
		l.Head = newNode
		return
	}

	newNode.Next = l.Head
	l.Head = newNode
}

// Delete by value, O(n)
func (l *LinkedList) Delete(v int) {
	if l.Head == nil {
		fmt.Println("LinkedList delete error: empty list")
		return
	}

	cur := l.Head
	if cur.Value == v {
		l.Head = l.Head.Next
		return
	}

	for cur.Next != nil {
		if cur.Next.Value == v {
			cur.Next = cur.Next.Next
			return
		}
		cur = cur.Next
	}

	fmt.Println("LinkedList delete error: element not exit")
}

// Implement the fmt.print feature
func (l LinkedList) String() string {
	var res string
	cur := l.Head
	for cur != nil {
		res += fmt.Sprintf("%d \t", cur.Value)
		cur = cur.Next
	}
	return res
}

func (l *LinkedList) LL2Array() []int {
	res := make([]int, 0)
	helper := l.Head
	for helper != nil {
		res = append(res, helper.Value)
		helper = helper.Next
	}

	return res
}

func Array2LL(nums []int) LinkedList {
	head := &Node{nums[0], nil}
	cur := head
	for i := 1; i < len(nums); i++ {
		tmp := &Node{nums[i], nil}
		cur.Next = tmp
		cur = tmp
	}

	return LinkedList{Head: head}
}
