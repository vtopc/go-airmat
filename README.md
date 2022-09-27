# go-airmat

Float Air Mattress is a container for a slice that could be used with [sync.Pool](https://pkg.go.dev/sync#Pool).
Has a minimum allocations.

## Benchmarks
BenchmarkMattress is kinda unreal situation, but BenchmarkPool is.
```shell
go test -bench=. -cpu=4 -benchmem
goos: darwin
goarch: amd64
pkg: github.com/vtopc/go-airmat
cpu: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz
BenchmarkMakeSlices-4   	  400750	      2688 ns/op	   16368 B/op	      10 allocs/op
BenchmarkMattress-4     	10026760	       120.1 ns/op	       0 B/op	       0 allocs/op
BenchmarkPool/8-4       	 4274638	       302.6 ns/op	       0 B/op	       0 allocs/op
BenchmarkPool/16-4      	   59446	     20567 ns/op	      30 B/op	       0 allocs/op
BenchmarkPool/24-4      	     100	  11741976 ns/op	 4023936 B/op	       0 allocs/op
```
