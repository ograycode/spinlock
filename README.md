# Spinlock

[![Go Reference](https://pkg.go.dev/badge/github.com/ograycode/spinlock.svg)](https://pkg.go.dev/github.com/ograycode/spinlock)

An implementation of a [Spinlock](https://en.wikipedia.org/wiki/Spinlock) for Go.

## Performance

Early indications are that it is good for certain types of work.

Tested with go1.15.2 windows/amd64 on a i7-8750H CPU @ 2.20GHz, 2208 Mhz, 6 Cores, 12 Logical Processors
```
go test -benchmem -bench .                                                    
goos: windows
goarch: amd64
BenchmarkSpinLockParallel-12            42856836                25.3 ns/op             0 B/op          0 allocs/op
BenchmarkMutexParallel-12               19999933                62.8 ns/op             0 B/op          0 allocs/op
BenchmarkSpinlock-12                    100000000               10.9 ns/op             0 B/op          0 allocs/op
BenchmarkMutex-12                       100000000               10.9 ns/op             0 B/op          0 allocs/op
BenchmarkDoWork-12                      1000000000               0.254 ns/op           0 B/op          0 allocs/op
PASS
```

## TODO

- More tests.
- More benchmarks.