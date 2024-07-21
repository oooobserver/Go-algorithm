package skiplist

import "sync"

//
// This is th simple concurrent skip list implmentation
// It just add RWLock to the normal list
//

type SCSkipList struct {
	mu    sync.RWMutex
	inner *SkipList
}

func NewSCSkipList() *SCSkipList {
	return &SCSkipList{
		inner: NewSkipList(),
	}
}

func (l *SCSkipList) Get(key int) (int, bool) {
	l.mu.RLock()
	defer l.mu.RUnlock()

	res, has := l.inner.Get(key)
	return res, has
}

func (l *SCSkipList) Put(key, value int) {
	l.mu.Lock()
	defer l.mu.Unlock()

	l.inner.Put(key, value)
}
