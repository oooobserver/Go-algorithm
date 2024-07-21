package skiplist

import (
	"sync"
	"sync/atomic"
)

// TODO: finish this

/* CSL
The most straightforward way to implement this is by using a read-write lock.
However, the performance of this method is not satisfactory due to the significant contention.

We notice that whether `put` or only affected the left side node
- put: modify the left side node's next point to the new add node
- get: keep the left side node, check if the next node is satisfy
- update: first get, then update

ConcurrentSkipList 中涉及到三类锁：

deleteMutex：一把全局的读写锁；get、put 操作取读锁，实现共享；delete 操作取写锁，实现全局互斥
keyMutex：每个 key 对应的一把互斥锁. 针对同一个 key 的 put 操作需要取 key 锁实现互斥
nodeMutex：每个 node 对应的一把读写锁. 在 get 检索过程中，会逐层对左边界节点加读锁；put 在插入新节点过程中，会逐层对左边界节点加写锁.
*/

type ConcurrentSkipList struct {
	head  *Node
	level atomic.Int32

	length atomic.Int32

	// Put, get, update get the read lock
	// Delete acquire the write lock
	DeleteMutex sync.RWMutex

	// The mutex of each key
	keyMutex sync.Map

	// Object pool, reuse the node object, reduce gc
	nodesCache sync.Pool

	compare func(key1, key2 any) bool
}

// Node represents a node in the skip list.
type Node struct {
	key   any
	value any
	next  []*Node
	lock  sync.RWMutex
}

func NewConcurrentSkipList(compareFunc func(key1, key2 any) bool) *ConcurrentSkipList {
	return &ConcurrentSkipList{
		head: &Node{
			next: make([]*Node, maxLevel),
		},

		nodesCache: sync.Pool{
			New: func() any {
				return &node{}
			},
		},

		compare: compareFunc,
	}
}

func (c *ConcurrentSkipList) Get(key any) (any, bool) {
	c.DeleteMutex.RLock()
	defer c.DeleteMutex.RUnlock()

	cur := c.head
	// 通过 last 记录上一层所在的节点位置，避免在下降的过程中对于同一个节点反复加多次左边界节点锁
	var last *Node

	for i := c.level.Load() - 1; i >= 0; i-- {
		// 在同一层中一路无锁穿越，直到来到左边界
		for cur.next[i] != nil && c.compare(cur.next[i].key, key) {
			cur = cur.next[i]
		}

		// 走到左边界
		// 通过 last 指针保证对同一个节点只会加一次左边界节点锁
		// Add lock at left side node
		// Use `last` to record last node add
		// QUESTION: only unlock at return, is there a issue
		if cur != last {
			cur.lock.RLock()
			defer cur.lock.RUnlock()
			last = cur
		}

		if cur.next[i] != nil && cur.next[i].key == key {
			return cur.next[i].value, true
		}
	}

	return 0, false
}
