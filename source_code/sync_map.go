package sourcecode

// import (
// 	"sync/atomic"
// )

// /*
// Loads, stores, and deletes run in amortized constant time.

// The Map type is optimized for two common use cases:
// 1. when the entry for a given key is only ever written once but read many times
// 2. when multiple goroutines read, write, and overwrite entries for disjoint sets of keys
// */

// type Map struct {
// 	mu Mutex

// 	/*
// 	   read contains the portion of the map's contents that are safe for concurrent access (with or without mu held)

// 	   The read field itself is always safe to load, but must only be stored with mu held

// 	   Entries stored in read may be updated concurrently without mu, but updating a previously-expunged entry requires that the entry be copied to the dirty map and unexpunged with mu held

// 	*/

// 	read atomic.Pointer[readOnly]

// 	// Expunged entries are not stored in the dirty map.
// 	//
// 	// If the dirty map is nil, the next write to the map will initialize it by making a shallow copy of the clean map
// 	dirty map[any]*entry

// 	misses int
// }

// type readOnly struct {
// 	m map[any]*entry
// 	// true if the dirty map contains some key not in m.
// 	amended bool
// }

// // expunged is an arbitrary pointer that marks entries which have been deleted
// var expunged = new(any)

// type entry struct {
// 	p atomic.Pointer[any]
// }

// func newEntry(i any) *entry {
// 	e := &entry{}
// 	e.p.Store(&i)
// 	return e
// }

// func (m *Map) loadReadOnly() readOnly {
// 	if p := m.read.Load(); p != nil {
// 		return *p
// 	}
// 	return readOnly{}
// }

// func (m *Map) Load(key any) (value any, ok bool) {
// 	read := m.loadReadOnly()
// 	e, ok := read.m[key]
// 	if !ok && read.amended {
// 		m.mu.Lock()
// 		//Double-check, avoid reporting a spurious miss while we were blocked on m.mu.
// 		read = m.loadReadOnly()
// 		e, ok = read.m[key]
// 		if !ok && read.amended {
// 			e, ok = m.dirty[key]
// 			// Regardless of whether the entry was present, record a miss
// 			m.missLocked()
// 		}
// 		m.mu.Unlock()
// 	}
// 	if !ok {
// 		return nil, false
// 	}
// 	return e.load()
// }

// func (e *entry) load() (value any, ok bool) {
// 	p := e.p.Load()
// 	if p == nil || p == expunged {
// 		return nil, false
// 	}
// 	return *p, true
// }

// // Store sets the value for a key.
// func (m *Map) Store(key, value any) {
// 	_, _ = m.Swap(key, value)
// }

// // tryCompareAndSwap compare the entry with the given old value and swaps
// // it with a new value if the entry is equal to the old value, and the entry
// // has not been expunged.
// //
// // If the entry is expunged, tryCompareAndSwap returns false and leaves
// // the entry unchanged.
// func (e *entry) tryCompareAndSwap(old, new any) bool {
// 	p := e.p.Load()
// 	if p == nil || p == expunged || *p != old {
// 		return false
// 	}

// 	// Copy the interface after the first load to make this method more amenable
// 	// to escape analysis: if the comparison fails from the start, we shouldn't
// 	// bother heap-allocating an interface value to store.
// 	nc := new
// 	for {
// 		if e.p.CompareAndSwap(p, &nc) {
// 			return true
// 		}
// 		p = e.p.Load()
// 		if p == nil || p == expunged || *p != old {
// 			return false
// 		}
// 	}
// }

// // unexpungeLocked ensures that the entry is not marked as expunged.
// //
// // If the entry was previously expunged, it must be added to the dirty map
// // before m.mu is unlocked.
// func (e *entry) unexpungeLocked() (wasExpunged bool) {
// 	return e.p.CompareAndSwap(expunged, nil)
// }

// // swapLocked unconditionally swaps a value into the entry.
// //
// // The entry must be known not to be expunged.
// func (e *entry) swapLocked(i *any) *any {
// 	return e.p.Swap(i)
// }

// // LoadOrStore returns the existing value for the key if present.
// // Otherwise, it stores and returns the given value.
// // The loaded result is true if the value was loaded, false if stored.
// func (m *Map) LoadOrStore(key, value any) (actual any, loaded bool) {
// 	// Avoid locking if it's a clean hit.
// 	read := m.loadReadOnly()
// 	if e, ok := read.m[key]; ok {
// 		actual, loaded, ok := e.tryLoadOrStore(value)
// 		if ok {
// 			return actual, loaded
// 		}
// 	}

// 	m.mu.Lock()
// 	read = m.loadReadOnly()
// 	if e, ok := read.m[key]; ok {
// 		if e.unexpungeLocked() {
// 			m.dirty[key] = e
// 		}
// 		actual, loaded, _ = e.tryLoadOrStore(value)
// 	} else if e, ok := m.dirty[key]; ok {
// 		actual, loaded, _ = e.tryLoadOrStore(value)
// 		m.missLocked()
// 	} else {
// 		if !read.amended {
// 			// We're adding the first new key to the dirty map.
// 			// Make sure it is allocated and mark the read-only map as incomplete.
// 			m.dirtyLocked()
// 			m.read.Store(&readOnly{m: read.m, amended: true})
// 		}
// 		m.dirty[key] = newEntry(value)
// 		actual, loaded = value, false
// 	}
// 	m.mu.Unlock()

// 	return actual, loaded
// }

// // tryLoadOrStore atomically loads or stores a value if the entry is not
// // expunged.
// //
// // If the entry is expunged, tryLoadOrStore leaves the entry unchanged and
// // returns with ok==false.
// func (e *entry) tryLoadOrStore(i any) (actual any, loaded, ok bool) {
// 	p := e.p.Load()
// 	if p == expunged {
// 		return nil, false, false
// 	}
// 	if p != nil {
// 		return *p, true, true
// 	}

// 	// Copy the interface after the first load to make this method more amenable
// 	// to escape analysis: if we hit the "load" path or the entry is expunged, we
// 	// shouldn't bother heap-allocating.
// 	ic := i
// 	for {
// 		if e.p.CompareAndSwap(nil, &ic) {
// 			return i, false, true
// 		}
// 		p = e.p.Load()
// 		if p == expunged {
// 			return nil, false, false
// 		}
// 		if p != nil {
// 			return *p, true, true
// 		}
// 	}
// }

// // LoadAndDelete deletes the value for a key, returning the previous value if any.
// // The loaded result reports whether the key was present.
// func (m *Map) LoadAndDelete(key any) (value any, loaded bool) {
// 	read := m.loadReadOnly()
// 	e, ok := read.m[key]
// 	if !ok && read.amended {
// 		m.mu.Lock()
// 		read = m.loadReadOnly()
// 		e, ok = read.m[key]
// 		if !ok && read.amended {
// 			e, ok = m.dirty[key]
// 			delete(m.dirty, key)
// 			// Regardless of whether the entry was present, record a miss: this key
// 			// will take the slow path until the dirty map is promoted to the read
// 			// map.
// 			m.missLocked()
// 		}
// 		m.mu.Unlock()
// 	}
// 	if ok {
// 		return e.delete()
// 	}
// 	return nil, false
// }

// // Delete deletes the value for a key.
// func (m *Map) Delete(key any) {
// 	m.LoadAndDelete(key)
// }

// func (e *entry) delete() (value any, ok bool) {
// 	for {
// 		p := e.p.Load()
// 		if p == nil || p == expunged {
// 			return nil, false
// 		}
// 		if e.p.CompareAndSwap(p, nil) {
// 			return *p, true
// 		}
// 	}
// }

// // trySwap swaps a value if the entry has not been expunged.
// //
// // If the entry is expunged, trySwap returns false and leaves the entry
// // unchanged.
// func (e *entry) trySwap(i *any) (*any, bool) {
// 	for {
// 		p := e.p.Load()
// 		if p == expunged {
// 			return nil, false
// 		}
// 		if e.p.CompareAndSwap(p, i) {
// 			return p, true
// 		}
// 	}
// }

// // Swap swaps the value for a key and returns the previous value if any.
// // The loaded result reports whether the key was present.
// func (m *Map) Swap(key, value any) (previous any, loaded bool) {
// 	read := m.loadReadOnly()
// 	if e, ok := read.m[key]; ok {
// 		if v, ok := e.trySwap(&value); ok {
// 			if v == nil {
// 				return nil, false
// 			}
// 			return *v, true
// 		}
// 	}

// 	m.mu.Lock()
// 	read = m.loadReadOnly()
// 	if e, ok := read.m[key]; ok {
// 		if e.unexpungeLocked() {
// 			// The entry was previously expunged, which implies that there is a
// 			// non-nil dirty map and this entry is not in it.
// 			m.dirty[key] = e
// 		}
// 		if v := e.swapLocked(&value); v != nil {
// 			loaded = true
// 			previous = *v
// 		}
// 	} else if e, ok := m.dirty[key]; ok {
// 		if v := e.swapLocked(&value); v != nil {
// 			loaded = true
// 			previous = *v
// 		}
// 	} else {
// 		if !read.amended {
// 			// We're adding the first new key to the dirty map.
// 			// Make sure it is allocated and mark the read-only map as incomplete.
// 			m.dirtyLocked()
// 			m.read.Store(&readOnly{m: read.m, amended: true})
// 		}
// 		m.dirty[key] = newEntry(value)
// 	}
// 	m.mu.Unlock()
// 	return previous, loaded
// }

// // CompareAndSwap swaps the old and new values for key
// // if the value stored in the map is equal to old.
// // The old value must be of a comparable type.
// func (m *Map) CompareAndSwap(key, old, new any) bool {
// 	read := m.loadReadOnly()
// 	if e, ok := read.m[key]; ok {
// 		return e.tryCompareAndSwap(old, new)
// 	} else if !read.amended {
// 		return false // No existing value for key.
// 	}

// 	m.mu.Lock()
// 	defer m.mu.Unlock()
// 	read = m.loadReadOnly()
// 	swapped := false
// 	if e, ok := read.m[key]; ok {
// 		swapped = e.tryCompareAndSwap(old, new)
// 	} else if e, ok := m.dirty[key]; ok {
// 		swapped = e.tryCompareAndSwap(old, new)
// 		// We needed to lock mu in order to load the entry for key,
// 		// and the operation didn't change the set of keys in the map
// 		// (so it would be made more efficient by promoting the dirty
// 		// map to read-only).
// 		// Count it as a miss so that we will eventually switch to the
// 		// more efficient steady state.
// 		m.missLocked()
// 	}
// 	return swapped
// }

// // CompareAndDelete deletes the entry for key if its value is equal to old.
// // The old value must be of a comparable type.
// //
// // If there is no current value for key in the map, CompareAndDelete
// // returns false (even if the old value is the nil interface value).
// func (m *Map) CompareAndDelete(key, old any) (deleted bool) {
// 	read := m.loadReadOnly()
// 	e, ok := read.m[key]
// 	if !ok && read.amended {
// 		m.mu.Lock()
// 		read = m.loadReadOnly()
// 		e, ok = read.m[key]
// 		if !ok && read.amended {
// 			e, ok = m.dirty[key]
// 			// Don't delete key from m.dirty: we still need to do the “compare” part
// 			// of the operation. The entry will eventually be expunged when the
// 			// dirty map is promoted to the read map.
// 			//
// 			// Regardless of whether the entry was present, record a miss: this key
// 			// will take the slow path until the dirty map is promoted to the read
// 			// map.
// 			m.missLocked()
// 		}
// 		m.mu.Unlock()
// 	}
// 	for ok {
// 		p := e.p.Load()
// 		if p == nil || p == expunged || *p != old {
// 			return false
// 		}
// 		if e.p.CompareAndSwap(p, nil) {
// 			return true
// 		}
// 	}
// 	return false
// }

// // Range calls f sequentially for each key and value present in the map.
// // If f returns false, range stops the iteration.
// //
// // Range does not necessarily correspond to any consistent snapshot of the Map's
// // contents: no key will be visited more than once, but if the value for any key
// // is stored or deleted concurrently (including by f), Range may reflect any
// // mapping for that key from any point during the Range call. Range does not
// // block other methods on the receiver; even f itself may call any method on m.
// //
// // Range may be O(N) with the number of elements in the map even if f returns
// // false after a constant number of calls.
// func (m *Map) Range(f func(key, value any) bool) {
// 	// We need to be able to iterate over all of the keys that were already
// 	// present at the start of the call to Range.
// 	// If read.amended is false, then read.m satisfies that property without
// 	// requiring us to hold m.mu for a long time.
// 	read := m.loadReadOnly()
// 	if read.amended {
// 		// m.dirty contains keys not in read.m. Fortunately, Range is already O(N)
// 		// (assuming the caller does not break out early), so a call to Range
// 		// amortizes an entire copy of the map: we can promote the dirty copy
// 		// immediately!
// 		m.mu.Lock()
// 		read = m.loadReadOnly()
// 		if read.amended {
// 			read = readOnly{m: m.dirty}
// 			copyRead := read
// 			m.read.Store(&copyRead)
// 			m.dirty = nil
// 			m.misses = 0
// 		}
// 		m.mu.Unlock()
// 	}

// 	for k, e := range read.m {
// 		v, ok := e.load()
// 		if !ok {
// 			continue
// 		}
// 		if !f(k, v) {
// 			break
// 		}
// 	}
// }

// func (m *Map) missLocked() {
// 	m.misses++
// 	if m.misses < len(m.dirty) {
// 		return
// 	}
// 	m.read.Store(&readOnly{m: m.dirty})
// 	m.dirty = nil
// 	m.misses = 0
// }

// func (m *Map) dirtyLocked() {
// 	if m.dirty != nil {
// 		return
// 	}

// 	read := m.loadReadOnly()
// 	m.dirty = make(map[any]*entry, len(read.m))
// 	for k, e := range read.m {
// 		if !e.tryExpungeLocked() {
// 			m.dirty[k] = e
// 		}
// 	}
// }

// func (e *entry) tryExpungeLocked() (isExpunged bool) {
// 	p := e.p.Load()
// 	for p == nil {
// 		if e.p.CompareAndSwap(nil, expunged) {
// 			return true
// 		}
// 		p = e.p.Load()
// 	}
// 	return p == expunged
// }
