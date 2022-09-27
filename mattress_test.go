package airmat

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const size = 10

func TestNewMattress(t *testing.T) {
	type S struct {
		F int
	}

	t.Run("simple", func(t *testing.T) {
		m := NewMattress[int](size)
		assert.Len(t, m.Slice, size)
	})

	t.Run("struct", func(t *testing.T) {
		m := NewMattress[S](size)
		assert.Len(t, m.Slice, size)

		m.Slice[1].F = 42
		assert.Equal(t, m.Slice[1].F, 42)
	})

	t.Run("struct_ptr", func(t *testing.T) {
		m := NewMattress[*S](size)
		assert.Len(t, m.Slice, size)

		// TODO: automate?
		if m.Slice[1] == nil {
			m.Slice[1] = &S{}
		}

		m.Slice[1].F = 42
		assert.Equal(t, m.Slice[1].F, 42)
	})
}

func TestMattress_Grow(t *testing.T) {
	tests := []struct {
		name string
		m    *Mattress[string]
		len  int
		cap  int
	}{
		{
			name: "nil",
			m:    &Mattress[string]{},
			len:  size,
			cap:  size,
		},
		{
			name: "equal",
			m:    NewMattress[string](size),
			len:  size,
			cap:  size,
		},
		{
			name: "grow_one",
			m:    NewMattress[string](size),
			len:  size + 1,
			cap:  size * 2, // might be flaky assert
		},
		{
			name: "grow_twice",
			m:    NewMattress[string](size),
			len:  size * 2,
			cap:  size * 2, // might be flaky assert
		},
		{
			name: "shrink_one",
			m:    NewMattress[string](size),
			len:  size - 1,
			cap:  size,
		},
		{
			name: "shrink_twice",
			m:    NewMattress[string](size),
			len:  size / 2,
			cap:  size,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.m.SetSize(tt.len)
			assert.Equal(t, tt.len, len(tt.m.Slice))
			assert.Equal(t, tt.cap, cap(tt.m.Slice))
		})
	}
}
