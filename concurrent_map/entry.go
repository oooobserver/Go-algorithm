package concurrentmap

import "sync/atomic"

// expunged is an arbitrary pointer that marks entries which have been deleted from the dirty map.
var expunged = new(any)

type entry struct {
	p atomic.Pointer[any]
}

func newEntry(i any) *entry {
	e := &entry{}
	e.p.Store(&i)
	return e
}

func (e *entry) load() (value any, ok bool) {
	p := e.p.Load()
	if p == nil || p == expunged {
		return nil, false
	}
	return *p, true
}

// Swaps a value if the entry has not been expunged
//
// If the entry is expunged, trySwap returns false and leaves the entry unchanged
func (e *entry) trySwap(i *any) (*any, bool) {
	for {
		p := e.p.Load()
		if p == expunged {
			return nil, false
		}

		// Because this method does not hold lock
		// the entry may change after checking
		// so use compare-and-swap and for loop
		if e.p.CompareAndSwap(p, i) {
			return p, true
		}
	}
}

// unexpungeLocked ensures that the entry is not marked as expunged.
//
// If the entry was previously expunged, it must be added to the dirty map before m.mu is unlocked.
func (e *entry) unexpungeLocked() (wasExpunged bool) {
	return e.p.CompareAndSwap(expunged, nil)
}

// Unconditionally swaps a value into the entry.
//
// The entry must be known not to be expunged.
func (e *entry) swapLocked(i *any) *any {
	return e.p.Swap(i)
}

// If entry is nil, change to expunge
func (e *entry) tryExpungeLocked() (isExpunged bool) {
	p := e.p.Load()
	// Because read map can change without lock
	// So still need for loop here
	for p == nil {
		if e.p.CompareAndSwap(nil, expunged) {
			return true
		}
		p = e.p.Load()
	}
	return p == expunged
}

func (e *entry) tryCompareAndSwap(old, new any) bool {
	p := e.p.Load()
	if p == nil || p == expunged || *p != old {
		return false
	}

	nc := new
	for {
		if e.p.CompareAndSwap(p, &nc) {
			return true
		}
		p = e.p.Load()
		if p == nil || p == expunged || *p != old {
			return false
		}
	}
}

func (e *entry) delete() (value any, ok bool) {
	for {
		p := e.p.Load()
		if p == nil || p == expunged {
			return nil, false
		}
		if e.p.CompareAndSwap(p, nil) {
			return *p, true
		}
	}
}
