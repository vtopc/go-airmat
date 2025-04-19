# go-airmat

![Logo](https://github.com/vtopc/go-airmat/blob/master/images/logo.png?raw=true)

Float Air Mattress is a container for a slice that could be used with [sync.Pool](https://pkg.go.dev/sync#Pool).
Has a minimum allocations.

## Benchmarks
```shell
go test -bench=. -cpu=4 -benchmem
goos: darwin
goarch: amd64
pkg: github.com/vtopc/go-airmat
cpu: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz
BenchmarkMakeSlices/8-4 	 1573507	       719.0 ns/op	    4080 B/op	       8 allocs/op
BenchmarkMakeSlices/16-4         	   10000	    125151 ns/op	 1048562 B/op	      16 allocs/op
BenchmarkMakeSlices/24-4         	     100	  18851436 ns/op	268435446 B/op	      24 allocs/op
BenchmarkPool/8-4                	 3459918	       289.3 ns/op	       0 B/op	       0 allocs/op
BenchmarkPool/16-4               	   54115	     20723 ns/op	       9 B/op	       0 allocs/op
BenchmarkPool/24-4               	     110	  11347946 ns/op	      15 B/op	       0 allocs/op
```

## TODO
1. Shrink huge Mattresses(https://wundergraph.com/blog/golang-sync-pool), e.g.:
  - https://github.com/golang/go/issues/27735#issuecomment-739169121
  - https://github.com/valyala/bytebufferpool
