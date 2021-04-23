package visited

// List is a reusable list with very efficient resets. Inspired by the C++
// implementation in hnswlib it can be reset with zero memrory writes in the
// array by moving the match target instead of altering the list. Only after a
// version overflow do we need to actually reset
type List struct {
	store   []uint8
	version uint8
}

func NewList(size int) *List {
	return &List{
		// start at 1 since the initial value of the list is already 0, so we need
		// something to differentiate from that
		version: 1,
		store:   make([]uint8, size),
	}
}

func (l *List) Visit(node uint64) {
	l.store[node] = l.version
}

func (l *List) Visited(node uint64) bool {
	return l.store[node] == l.version
}

func (l *List) Reset() {
	l.version++

	if l.version == 0 {
		// 0 is not a valid version because it conflicts with the initial value of
		// the array
		l.version = 1
	}

	// we have overflowed and need an actual reset
	if l.version == 1 {
		for i := range l.store {
			l.store[i] = 0
		}
	}
}
