package airmat_test

import (
	"strconv"
	"testing"

	"github.com/vtopc/go-airmat"
)

var parallelRequests = []int{8, 16, 24}

// Global variable for Benchmarking - https://itnext.io/the-top-10-most-common-mistakes-ive-seen-in-go-projects-4b79d4f6cd65
var (
	s    []int
	mm   []*airmat.Mattress[int]
	pool = airmat.NewPool[int]()
)

func BenchmarkMakeSlices(b *testing.B) {
	for _, requests := range parallelRequests {
		b.Run(strconv.Itoa(requests), func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				for i := 0; i < requests; i++ {
					v := make([]int, 2<<i)
					s = v
				}
			}
		})
	}
}

func BenchmarkPool(b *testing.B) {
	for _, requests := range parallelRequests {
		b.Run(strconv.Itoa(requests), func(b *testing.B) {
			mm = make([]*airmat.Mattress[int], requests)

			for n := 0; n < b.N; n++ {
				// emulates parallel requests
				for i := 0; i < requests; i++ {
					m := pool.Get()
					m.SetSize(2 << i)
					mm[i] = m
				}

				for i := 0; i < requests; i++ {
					pool.Put(mm[i])
				}
			}
		})
	}
}
