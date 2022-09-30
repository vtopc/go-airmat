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
BenchmarkMakeSlices/8-4 	 1460142	       819.4 ns/op	    4080 B/op	       8 allocs/op
BenchmarkMakeSlices/16-4         	    8475	    126708 ns/op	 1048563 B/op	      16 allocs/op
BenchmarkMakeSlices/24-4         	      96	  20826771 ns/op	268435442 B/op	      24 allocs/op
BenchmarkPool/8-4                	 3468649	       306.7 ns/op	       0 B/op	       0 allocs/op
BenchmarkPool/16-4               	   47094	     24404 ns/op	       0 B/op	       0 allocs/op
BenchmarkPool/24-4               	      91	  11934050 ns/op	      19 B/op	       0 allocs/op
```
