package linkedlist

import (
	"fmt"
	"math/rand"
)

const (
	maxLevel = 16
)

type node struct {
	key     int
	value   interface{}
	forward []*node
}

type skipList struct {
	head   *node
	level  int
	length int
}

func newNode(key int, value interface{}, level int) *node {
	return &node{
		key:     key,
		value:   value,
		forward: make([]*node, level),
	}
}

func newSkipList() *skipList {
	head := newNode(0, nil, maxLevel)
	return &skipList{
		head:   head,
		level:  1,
		length: 0,
	}
}

func (sl *skipList) randomLevel() int {
	level := 1
	for rand.Float64() < 0.5 && level < maxLevel {
		level++
	}
	return level
}

func (sl *skipList) insert(key int, value interface{}) {
	update := make([]*node, maxLevel)
	current := sl.head

	for i := sl.level - 1; i >= 0; i-- {
		for current.forward[i] != nil && current.forward[i].key < key {
			current = current.forward[i]
		}
		update[i] = current
	}

	level := sl.randomLevel()
	if level > sl.level {
		for i := sl.level; i < level; i++ {
			update[i] = sl.head
		}
		sl.level = level
	}

	node := newNode(key, value, level)
	for i := 0; i < level; i++ {
		node.forward[i] = update[i].forward[i]
		update[i].forward[i] = node
	}

	sl.length++
}

func (sl *skipList) search(key int) interface{} {
	current := sl.head
	for i := sl.level - 1; i >= 0; i-- {
		for current.forward[i] != nil && current.forward[i].key < key {
			current = current.forward[i]
		}
	}
	current = current.forward[0]
	if current != nil && current.key == key {
		return current.value
	}
	return nil
}

func (sl *skipList) delete(key int) {
	update := make([]*node, maxLevel)
	current := sl.head

	for i := sl.level - 1; i >= 0; i-- {
		for current.forward[i] != nil && current.forward[i].key < key {
			current = current.forward[i]
		}
		update[i] = current
	}

	current = current.forward[0]
	if current != nil && current.key == key {
		for i := 0; i < sl.level; i++ {
			if update[i].forward[i] != current {
				break
			}
			update[i].forward[i] = current.forward[i]
		}

		for sl.level > 1 && sl.head.forward[sl.level-1] == nil {
			sl.level--
		}

		sl.length--
	}
}

func (sl *skipList) print() {
	for i := sl.level - 1; i >= 0; i-- {
		current := sl.head.forward[i]
		fmt.Printf("Level %d: ", i)
		for current != nil {
			fmt.Printf("(%d, %v) ", current.key, current.value)
			current = current.forward[i]
		}
		fmt.Println()
	}
}

// func main() {
	
// }
