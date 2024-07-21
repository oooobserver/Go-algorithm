package skiplist

import (
	"fmt"
	"math/rand"
)

const (
	maxLevel = 32
)

type SkipList struct {
	head   *node
	level  int
	length int
}

type node struct {
	key   int
	value int
	nexts []*node
}

func NewSkipList() *SkipList {
	head := newNode(0, 0, maxLevel)
	return &SkipList{
		head:   head,
		level:  1,
		length: 0,
	}
}

func newNode(key int, value int, level int) *node {
	return &node{
		key:   key,
		value: value,
		nexts: make([]*node, level),
	}
}

func (sl *SkipList) Length() int {
	return sl.length
}

// O(logn)
func (sl *SkipList) search(key int) *node {
	cur := sl.head
	/* Four situations:
	- the next node in this level is bigger: go to next node
	- the next node in this level is smaller: go to next level
	- the next node in this level is nil: go to next level
	- the next node in this level is same: found

	After one loop, the next node of current is smaller/same/nil
	*/

	for i := sl.level - 1; i >= 0; i-- {
		// Keep going until next element is smaller or equal
		for cur.nexts[i] != nil && cur.nexts[i].key < key {
			cur = cur.nexts[i]
		}

		// If not same, go to next level
		if cur.nexts[i] != nil && cur.nexts[i].key == key {
			return cur.nexts[i]
		}
	}
	return nil
}

func (sl *SkipList) Get(key int) (int, bool) {
	node := sl.search(key)
	if node == nil {
		return 0, false
	}

	return node.value, true
}

// If key is present then update otherwise insert a new KV
func (sl *SkipList) Put(key, value int) {
	node := sl.search(key)
	if node != nil {
		node.value = value
		return
	}

	level := sl.randomLevel()
	if level > sl.level {
		sl.level = level
	}

	node = newNode(key, value, level)

	// Reconfig
	cur := sl.head
	for i := level - 1; i >= 0; i-- {
		for cur.nexts[i] != nil && cur.nexts[i].key < key {
			cur = cur.nexts[i]
		}
		// Add node behind nil or the first node larger
		node.nexts[i] = cur.nexts[i]
		cur.nexts[i] = node
	}
	sl.length++
}

func (sl *SkipList) randomLevel() int {
	level := 1
	for rand.Float64() < 0.5 && level < maxLevel {
		level++
	}
	return level
}

func (sl *SkipList) Delete(key int) {
	n := sl.search(key)
	if n == nil {
		return
	}

	cur := sl.head
	for i := sl.level - 1; i >= 0; i-- {
		for cur.nexts[i] != nil && cur.nexts[i].key < key {
			cur = cur.nexts[i]
		}

		// TODO: improve-from bottom to up
		if cur.nexts[i] == nil || cur.nexts[i].key > key {
			continue
		}

		cur.nexts[i] = cur.nexts[i].nexts[i]
	}

	// Remove useless lvel
	for sl.level > 1 && sl.head.nexts[sl.level-1] == nil {
		sl.level--
	}

	sl.length--
}

// Implement the fmt.print feature
func (sl *SkipList) String() string {
	cur := sl.head.nexts[0]
	res := ""
	for cur != nil {
		res += fmt.Sprintf("[%d/%d]   ", cur.key, cur.value)
		cur = cur.nexts[0]
	}

	res += "\n"
	return res
}

type SkipListIterator struct {
	node  *node
	key   int
	value int
}

func (l *SkipList) NewSkipListIterator() *SkipListIterator {
	iter := SkipListIterator{
		node: l.head,
	}
	iter.Next()
	return &iter
}

func (iter *SkipListIterator) Key() int {
	return iter.key
}

func (iter *SkipListIterator) Value() int {
	return iter.value
}

func (iter *SkipListIterator) Next() bool {
	iter.node = iter.node.nexts[0]
	if iter.node == nil {
		return false
	}

	iter.key = iter.node.key
	iter.value = iter.node.value
	return true
}
