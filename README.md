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
BenchmarkMakeSlices-4   	  419877	      2636 ns/op	   16368 B/op	      10 allocs/op
BenchmarkMattress-4     	 9954985	       118.1 ns/op	       0 B/op	       0 allocs/op
BenchmarkPool-4         	 2420287	       485.0 ns/op	       0 B/op	       0 allocs/op
```
