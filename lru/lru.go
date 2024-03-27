package lru

import "fmt"

type LRUCache struct {
	capacity int
	size     int
	cache    map[int]*node
	head     *node
}

type node struct {
	key  int
	val  int
	next *node
	prev *node
}

func Constructor(capacity int) LRUCache {
	lru := LRUCache{
		capacity: capacity,
		cache:    make(map[int]*node),
		head:     &node{},
	}
	lru.head.next = lru.head
	return lru
}

func (this *LRUCache) Get(key int) int {
	if node, ok := this.cache[key]; ok {

		// Push front
		this.push_front(node)
		return node.val
	}

	fmt.Printf("get: ")
	this.Print()
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	// In the cache
	if this.Get(key) != -1 {
		this.cache[key].val = value
		this.push_front(this.cache[key])
	} else {
		// Can directly put
		if this.size < this.capacity {
			this.size++

			// Insert
			next := this.head.next
			tmp := &node{
				key:  key,
				val:  value,
				next: next,
				prev: this.head,
			}
			this.head.next = tmp
			next.prev = tmp
			this.cache[key] = tmp
		} else {
			candidate := this.head.prev
			this.push_front(candidate)
			delete(this.cache, candidate.key)

			candidate.key = key
			candidate.val = value
			this.cache[key] = candidate

		}

	}

	fmt.Printf("put: ")
	this.Print()
}

func (this *LRUCache) push_front(node *node) {
	// fmt.Printf("Push front\n")
	node.prev.next = node.next
	node.next.prev = node.prev

	node.next = this.head.next
	node.prev = this.head

	this.head.next.prev = node
	this.head.next = node
}

func (this *LRUCache) Print() {
	node := this.head.next
	for node != nil && node != this.head {
		fmt.Printf("%d\t", node.val)
		node = node.next
	}
	fmt.Printf("\n")
}
