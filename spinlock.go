package spinlock

import (
	"runtime"
	"sync/atomic"
)

const (
	mutexUnlocked = iota
	mutexLocked
)

// Spinlock implements the Locker interface.
//
// It must not be copied.
type Spinlock struct {
	l int32
}

// Lock grabs and exlusive lock, blocking all other users.
// If the lock fails twice, runtime.Gosched() is called.
func (s *Spinlock) Lock() {
	if atomic.CompareAndSwapInt32(&s.l, mutexUnlocked, mutexLocked) {
		return
	}
	// Enable Lock to be inlined by removing the for loop.
	s.slowLock()
}

func (s *Spinlock) slowLock() {
	for {
		if atomic.CompareAndSwapInt32(&s.l, mutexUnlocked, mutexLocked) {
			return
		}
		runtime.Gosched()
	}
}

// Unlock will allow all other users of s to attempt to grab the lock.
// Grabing the lock is not in a deterministic order.
//
// Function will panic if a call to Unlock is made before Lock
func (s *Spinlock) Unlock() {
	if !atomic.CompareAndSwapInt32(&s.l, mutexLocked, mutexUnlocked) {
		panic("Tried to unlock when not locked")
	}
}
