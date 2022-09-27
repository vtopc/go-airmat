package airmat_test

import (
	"strconv"
	"testing"

	"github.com/vtopc/go-airmat"
)

// Global variable for Benchmarking - https://itnext.io/the-top-10-most-common-mistakes-ive-seen-in-go-projects-4b79d4f6cd65
var (
	m    = &airmat.Mattress[int]{}
	s    []int
	pool = airmat.NewPool[int]()
)

func BenchmarkMakeSlices(b *testing.B) {
	for n := 0; n < b.N; n++ {
		for i := 0; i < 10; i++ {
			v := make([]int, 2<<i)
			s = v
		}
	}
}

func BenchmarkMattress(b *testing.B) {
	for n := 0; n < b.N; n++ {
		for i := 0; i < 10; i++ {
			m.SetSize(2 << i)
		}
	}
}

func BenchmarkPool(b *testing.B) {
	poolSizes := []int{8, 16, 24}

	for _, poolSize := range poolSizes {
		b.Run(strconv.Itoa(poolSize), func(b *testing.B) {
			mm := make([]*airmat.Mattress[int], poolSize)

			for n := 0; n < b.N; n++ {
				for i := 0; i < poolSize; i++ {
					m := pool.Get()
					m.SetSize(2 << i)
					mm[i] = m
				}

				for i := 0; i < poolSize; i++ {
					pool.Put(mm[i])
				}
			}
		})
	}
}
