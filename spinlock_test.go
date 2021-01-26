package spinlock

import (
	"sync"
	"testing"
	"time"
)

// ensure it implenents the Locker interface.
var _ sync.Locker = (*Spinlock)(nil)

func TestSpinLockWorks(t *testing.T) {
	var lock Spinlock
	val := 0
	locked := make(chan int)

	go func() {
		lock.Lock()
		locked <- 1
		time.Sleep(3 * time.Second)
		val++
		lock.Unlock()
	}()
	<-locked
	if val != 0 {
		t.Errorf("Didn't wait for the lock correctly")
	}
	lock.Lock()
	if val != 1 {
		t.Errorf("Didn't get the lock")
	}
	lock.Unlock()
}

func BenchmarkSpinLockParallel(b *testing.B) {
	var lock Spinlock
	b.SetParallelism(3)
	b.RunParallel(func(pb *testing.PB) {
		foo := 1
		for pb.Next() {
			lock.Lock()
			foo = doWork(foo+1, 2)
			lock.Unlock()
		}
		_ = foo
	})
}

func BenchmarkMutexParallel(b *testing.B) {
	var lock sync.Mutex
	b.SetParallelism(3)
	b.RunParallel(func(pb *testing.PB) {
		foo := 1
		for pb.Next() {
			lock.Lock()
			foo = doWork(foo+1, 2)
			lock.Unlock()
		}
		_ = foo
	})
}

func BenchmarkSpinlock(b *testing.B) {
	var lock Spinlock
	foo := 0
	for x := 0; x < b.N; x++ {
		lock.Lock()
		foo = doWork(foo+1, 2)
		lock.Unlock()
	}
	_ = foo
}

func BenchmarkMutex(b *testing.B) {
	var lock sync.Mutex
	foo := 0
	for x := 0; x < b.N; x++ {
		lock.Lock()
		foo = doWork(foo+1, 2)
		lock.Unlock()
	}
	_ = foo
}

func BenchmarkDoWork(b *testing.B) {
	foo := 0
	for x := 0; x < b.N; x++ {
		foo = doWork(foo+1, 2)
	}
	_ = foo
}

func doWork(x, y int) int {
	return x % y
}
