package airmat

import (
	"sync"
)

// Pool of Mattress-es.
// Just a wrapper for https://pkg.go.dev/sync#Pool.
type Pool[T any] struct {
	sp *sync.Pool
}

func NewPool[T any]() *Pool[T] {
	return &Pool[T]{
		sp: &sync.Pool{
			New: func() any {
				// The Pool's New function should generally only return pointer
				// types, since a pointer can be put into the return interface
				// value without an allocation:
				return &Mattress[T]{}
			},
		},
	}
}

// Get selects an arbitrary Mattress from the Pool, removes it from the Pool, and returns it to the caller.
// It might return a Slice with some data, but since it should be written by index it's not an issue.
func (p *Pool[T]) Get() *Mattress[T] {
	return p.sp.Get().(*Mattress[T])
}

// Put adds x to the pool.
// Should be called when Mattress is not need anymore.
func (p *Pool[T]) Put(x *Mattress[T]) {
	p.sp.Put(x)
}
