# go-airmat

Float Air Mattress is a container for a slice that could be used with [sync.Pool](https://pkg.go.dev/sync#Pool).
Has a minimum allocations.

## Benchmarks
```shell
go test -bench=. -cpu=4 -benchmem
goos: darwin
goarch: amd64
pkg: github.com/vtopc/go-airmat
cpu: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz
BenchmarkMakeSlices/8-4 	 1677363	       731.3 ns/op	    4080 B/op	       8 allocs/op
BenchmarkMakeSlices/16-4         	   10000	    105324 ns/op	 1048563 B/op	      16 allocs/op
BenchmarkMakeSlices/24-4         	     100	  15483531 ns/op	268435443 B/op	      24 allocs/op
BenchmarkPool/8-4                	 4156912	       288.8 ns/op	       0 B/op	       0 allocs/op
BenchmarkPool/16-4               	   56299	     19531 ns/op	       9 B/op	       0 allocs/op
BenchmarkPool/24-4               	     133	   8669618 ns/op	      13 B/op	       0 allocs/op
```
