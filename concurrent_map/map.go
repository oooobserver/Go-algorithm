package concurrentmap

import (
	"sync"
	"sync/atomic"
)

type Map struct {
	mu sync.Mutex

	read atomic.Pointer[readOnly]

	dirty map[any]*entry

	misses int
}

type readOnly struct {
	m       map[any]*entry
	amended bool
}

func (m *Map) loadReadOnly() readOnly {
	if p := m.read.Load(); p != nil {
		return *p
	}
	return readOnly{}
}

//
// Read process
//

func (m *Map) Get(key any) (value any, ok bool) {
	read := m.loadReadOnly()
	e, ok := read.m[key]
	if !ok && read.amended {
		m.mu.Lock()
		// Double check
		read = m.loadReadOnly()
		e, ok = read.m[key]
		if !ok && read.amended {
			e, ok = m.dirty[key]
			m.missLocked()
		}
		m.mu.Unlock()
	}

	if !ok {
		return nil, false
	}

	return e.load()
}

// If misses hit the threshold, flip
func (m *Map) missLocked() {
	m.misses++
	if m.misses < len(m.dirty) {
		return
	}
	m.read.Store(&readOnly{m: m.dirty})
	m.dirty = nil
	m.misses = 0
}

//
// Write process include: put, swap
//

func (m *Map) Put(key, value any) {
	_, _ = m.Swap(key, value)
}

// The present indicate that wether the key is in delete state
//
// Delete state include nil and expunged
func (m *Map) Swap(key, value any) (previous any, present bool) {
	read := m.loadReadOnly()
	// Key present in the read map, can just swap
	if e, ok := read.m[key]; ok {
		if v, ok := e.trySwap(&value); ok {
			if v == nil {
				return nil, false
			}
			return *v, true
		}
	}

	// Dirty map
	m.mu.Lock()
	read = m.loadReadOnly()
	if e, ok := read.m[key]; ok {
		// The entry was previously expunged, now have to add to dirty map
		if e.unexpungeLocked() {
			m.dirty[key] = e
		}
		if v := e.swapLocked(&value); v != nil {
			present = true
			previous = *v
		}
	} else if e, ok := m.dirty[key]; ok {
		if v := e.swapLocked(&value); v != nil {
			present = true
			previous = *v
		}
	} else {
		// Adding the first new key to the dirty map.
		// Make sure it is allocated and mark the read-only map as incomplete.
		if !read.amended {
			m.dirtyLocked()
			m.read.Store(&readOnly{m: read.m, amended: true})
		}
		m.dirty[key] = newEntry(value)
	}

	m.mu.Unlock()
	return previous, present
}

// Make sure dirty map is not nil
//
// If nil, during iterating the read map, not load the expunge entry
func (m *Map) dirtyLocked() {
	if m.dirty != nil {
		return
	}

	read := m.loadReadOnly()
	m.dirty = make(map[any]*entry, len(read.m))
	for k, e := range read.m {
		if !e.tryExpungeLocked() {
			m.dirty[k] = e
		}
	}
}

func (m *Map) CompareAndSwap(key, old, new any) bool {
	read := m.loadReadOnly()
	if e, ok := read.m[key]; ok {
		return e.tryCompareAndSwap(old, new)
	} else if !read.amended {
		return false
	}

	// Go to dirty map
	m.mu.Lock()
	defer m.mu.Unlock()
	read = m.loadReadOnly()
	swapped := false
	if e, ok := read.m[key]; ok {
		swapped = e.tryCompareAndSwap(old, new)
	} else if e, ok := m.dirty[key]; ok {
		swapped = e.tryCompareAndSwap(old, new)
		m.missLocked()
	}
	return swapped
}

//
// Delete process
//

// Delete deletes the value for a key.
func (m *Map) Delete(key any) {
	m.LoadAndDelete(key)
}

func (m *Map) LoadAndDelete(key any) (value any, loaded bool) {
	read := m.loadReadOnly()
	e, ok := read.m[key]
	if !ok && read.amended {
		m.mu.Lock()
		read = m.loadReadOnly()
		e, ok = read.m[key]
		if !ok && read.amended {
			e, ok = m.dirty[key]
			delete(m.dirty, key)
			// Regardless of whether the entry was present, record a miss: this key
			// will take the slow path until the dirty map is promoted to the read
			// map.
			m.missLocked()
		}
		m.mu.Unlock()
	}
	if ok {
		return e.delete()
	}
	return nil, false
}

func (m *Map) Range(f func(key, value any) bool) {
	read := m.loadReadOnly()
	if read.amended {
		m.mu.Lock()
		read = m.loadReadOnly()
		if read.amended {
			read = readOnly{m: m.dirty}
			copyRead := read
			m.read.Store(&copyRead)
			m.dirty = nil
			m.misses = 0
		}
		m.mu.Unlock()
	}

	for k, e := range read.m {
		v, ok := e.load()
		if !ok {
			continue
		}
		if !f(k, v) {
			break
		}
	}
}
