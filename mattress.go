package airmat

// Mattress is a sync.Pool compatible slice,
// should be updated by the index, not with append(...) func.
type Mattress[T any] []T

func NewMattress[T any](size int) Mattress[T] {
	return make([]T, size)
}

// SetSize extends or shrinks length of underlying Slice.
func (m *Mattress[T]) SetSize(size int) {
	if m == nil || *m == nil {
		*m = make([]T, size)

		return
	}

	l := len(*m)

	switch {
	case l == size:
		break
	case l < size:
		m.extend(size)
	case l > size:
		m.shrink(size)
	}
}

func (m *Mattress[T]) extend(l int) {
	diff := l - len(*m)

	*m = append(*m, make([]T, diff)...)
}

func (m *Mattress[T]) shrink(l int) {
	*m = (*m)[:l]
}
