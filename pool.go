// Package chorister implements a simple worker pool for concurrent task execution.
package chorister

import (
	"time"
)

// CPool defines the interface for a worker pool.
// Worker pool based from github.com/phamnam2003/ants
type CPool interface {
	// Free returns the number of available workers, -1 indicates this pool is unlimited.
	Free() int

	// Cap returns the capacity of this pool.
	Cap() int

	// Running returns the number of workers currently running.
	Running() int

	// Release closes this pool and releases the worker queue.
	Release()

	// ReleaseTimeout is like Release but with a timeout, it waits all workers to exit before timing out.
	ReleaseTimeout(timeout time.Duration) error

	// Submit submits a task to the pool.
	//
	// Note that you are allowed to call Pool.Submit() from the current Pool.Submit(),
	// but what calls for special attention is that you will get blocked with the last
	// Pool.Submit() call once the current Pool runs out of its capacity, and to avoid this,
	// you should instantiate a Pool with ants.WithNonblocking(true).
	Submit(task func()) error

	Reboot()

	// Waiting returns the number of tasks waiting to be executed.
	Waiting() int

	// Tune changes the capacity of this pool, note that it is noneffective to the infinite or pre-allocation pool.
	Tune(size int)

	// IsClosed indicates whether the pool is closed.
	IsClosed() bool
}

// UnimplementCPool is a placeholder struct for unimplemented CPool methods.
type UnimplementCPool struct{}

// Free is a placeholder method for unimplemented CPool.Free.
func (u *UnimplementCPool) Free() int {
	return 0
}

// Cap is a placeholder method for unimplemented CPool.Cap.
func (u *UnimplementCPool) Cap() int {
	return 0
}

// Running is a placeholder method for unimplemented CPool.Running.
func (u *UnimplementCPool) Running() int {
	return 0
}

// Release is a placeholder method for unimplemented CPool.Release.
func (u *UnimplementCPool) Release() {}

// ReleaseTimeout is a placeholder method for unimplemented CPool.ReleaseTimeout.
func (u *UnimplementCPool) ReleaseTimeout(timeout time.Duration) error {
	return ErrUnimplemented
}

// Submit is a placeholder method for unimplemented CPool.Submit.
func (u *UnimplementCPool) Submit(task func()) error {
	return ErrUnimplemented
}

// Reboot is a placeholder method for unimplemented CPool.Reboot.
func (u *UnimplementCPool) Reboot() {}

// Waiting is a placeholder method for unimplemented CPool.Waiting.
func (u *UnimplementCPool) Waiting() int {
	return 0
}

// Tune is a placeholder method for unimplemented CPool.Tune.
func (u *UnimplementCPool) Tune(size int) {}

// IsClosed is a placeholder method for unimplemented CPool.IsClosed.
func (u *UnimplementCPool) IsClosed() bool {
	return false
}

// UnimplementCPoolWithFunc is a placeholder struct for unimplemented CPoolWithFunc methods.
type UnimplementCPoolWithFunc[T any] struct {
	*UnimplementCPool
}

// CPoolWithFunc defines the interface for a generic worker pool that accepts tasks with arguments.
type CPoolWithFunc[T any] interface {
	CPool
	Invoke(arg T) error
}

// Invoke is a placeholder method for unimplemented CPoolWithFunc.Invoke.
func (u *UnimplementCPoolWithFunc[T]) Invoke(arg T) error {
	return ErrUnimplemented
}
