package airmat

type Mattress[T any] struct {
	Slice []T
}

func NewMattress[T any](size int) *Mattress[T] {
	return &Mattress[T]{
		Slice: make([]T, size),
	}
}

// Grow extends or shrinks length of underlying Slice.
func (m *Mattress[T]) Grow(size int) {
	if m.Slice == nil {
		m.Slice = make([]T, size)

		return
	}

	l := len(m.Slice)

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
	diff := l - len(m.Slice)

	// if diff < 0 {
	// 	return
	// }

	m.Slice = append(m.Slice, make([]T, diff)...)
}

func (m *Mattress[T]) shrink(l int) {
	// if len(m.Slice) < len {
	// 	return
	// }

	m.Slice = m.Slice[:l]
}
