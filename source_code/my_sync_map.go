package sourcecode

import (
	"sync"
	"sync/atomic"
)

/*
Loads, stores, and deletes run in amortized constant time.

The Map type is optimized for two common use cases:
1. when the entry for a given key is only ever written once but read many times
2. when multiple goroutines read, write, and overwrite entries for disjoint sets of keys
*/

type Map struct {
	mu     sync.Mutex
	read   atomic.Pointer[readOnly]
	dirty  map[any]*entry
	misses int
}

type readOnly struct {
	m    map[any]*entry
	miss bool
}

type entry struct {
	p atomic.Pointer[any]
}

var deleted = new(any)

func (e *entry) load() (value any, ok bool) {
	p := e.p.Load()
	if p == nil || p == deleted {
		return nil, false
	}

	return *p, true
}

func newEntry(i any) *entry {
	e := &entry{}
	e.p.Store(&i)
	return e
}

// Read process
func (m *Map) loadReadOnly() readOnly {
	if p := m.read.Load(); p != nil {
		return *p
	}

	return readOnly{}
}

func (m *Map) Load(key any) (value any, ok bool) {
	read := m.loadReadOnly()
	v, ok := read.m[key]
	if !ok && read.miss {
		m.mu.Lock()
		read := m.loadReadOnly()
		v, ok = read.m[key]
		if !ok && read.miss {
			v, ok = m.dirty[key]
			m.missLocked()
		}
		m.mu.Unlock()
	}

	if !ok {
		return nil, false
	}

	return v.load()
}

func (m *Map) missLocked() {
	m.misses++
	if m.misses < len(m.dirty) {
		return
	}

	m.read.Store(&readOnly{m: m.dirty})
	m.dirty = nil
	m.misses = 0
}

// Write process
func (m *Map) Insert(key, value any) {
	_, _ = m.Swap(key, value)
}

// Swaps the value for a key and returns the previous value if any.
func (m *Map) Swap(key, value any) (previous any, loaded bool) {
	read := m.loadReadOnly()
	if e, ok := read.m[key]; ok {
		if v, ok := e.trySwap(&value); ok {
			if v == nil {
				return nil, false
			}
			return *v, true
		}
	}

	// The key is in the dirty map, need to hold the lock
	m.mu.Lock()
	// Always double check
	read = m.loadReadOnly()
	if e, ok := read.m[key]; ok {
		if e.undeleteLocked() {
			// If the entry is deleted in the read map
			// This must add in the dirty map
			m.dirty[key] = e
		}
		// The entry is not deleted in the read map
		// But could be nil
		if v := e.swapLocked(&value); v != nil {
			loaded = true
			previous = *v
		}

	} else if e, ok := m.dirty[key]; ok {
		if v := e.swapLocked(&value); v != nil {
			loaded = true
			previous = *v
		}
	} else {
		// This is a new key
		// Not in read or dirty
		if !read.miss {
			m.dirtyLocked()
			m.read.Store(&readOnly{m: read.m, miss: true})
		}
		m.dirty[key] = newEntry(value)
	}
	m.mu.Unlock()
	return previous, loaded
}

// Swaps a value if the entry has not been deleted.
// If the entry is deleted, returns false and leaves the entry unchanged
func (e *entry) trySwap(i *any) (*any, bool) {
	// Use `for` because corruent load without lock
	// while at CAP, the actual value point by pointer may already changed
	for {
		p := e.p.Load()
		if p == deleted {
			return nil, false
		}
		if e.p.CompareAndSwap(p, i) {
			return p, true
		}
	}
}

// Ensures that the entry is not marked as deleted. If the entry was previously deleted,
// it must be added to the dirty map before m.mu is unlocked.
func (e *entry) undeleteLocked() (wasDeleted bool) {
	return e.p.CompareAndSwap(deleted, nil)
}

// Swap a value into the entry, the entry must be known not to be deleted
func (e *entry) swapLocked(i *any) *any {
	return e.p.Swap(i)
}

func (m *Map) dirtyLocked() {
	if m.dirty != nil {
		return
	}

	read := m.loadReadOnly()
	m.dirty = make(map[any]*entry, len(read.m))
	// for k, e := range read.m {
	// 	// if !e.tryExpungeLocked() {
	// 	// 	m.dirty[k] = e
	// 	// }
	// }
}

// Atomically loads or stores a value if the entry is not deleted
// If the entry is deleted, leaves the entry unchanged
// Only load when the entry is nil
func (e *entry) tryLoadOrStore(i any) (actual any, loaded, ok bool) {
	p := e.p.Load()
	if p == deleted {
		return nil, false, false
	}
	if p != nil {
		return *p, true, true
	}

	ic := i
	for {
		if e.p.CompareAndSwap(nil, &ic) {
			return i, false, true
		}
		p = e.p.Load()
		if p == deleted {
			return nil, false, false
		}
		if p != nil {
			return *p, true, true
		}
	}
}
