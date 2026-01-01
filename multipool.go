package chorister

import "errors"

// CMultiPool defines the interface for a multi-worker pool.
// Multi-worker pool based from github.com/phamnam2003/ants
type CMultiPool interface {
	CPool

	// RunningByIndex returns the number of the currently running workers in the specific pool.
	RunningByIndex(idx int) (int, error)

	// FreeByIndex returns the number of available workers in the specific pool.
	FreeByIndex(idx int) (int, error)

	// WaitingByIndex returns the number of the currently waiting tasks in the specific pool.
	WaitingByIndex(idx int) (int, error)
}

// UnimplementCMultiPool is a placeholder struct for unimplemented CMultiPool methods.
type UnimplementCMultiPool struct {
	*UnimplementCPool
}

// FreeByIndex is a placeholder method for unimplemented CMultiPool.FreeByIndex.
func (u *UnimplementCMultiPool) FreeByIndex(idx int) (int, error) {
	return 0, errors.New(ErrUnimplemented)
}

// RunningByIndex is a placeholder method for unimplemented CMultiPool.RunningByIndex.
func (u *UnimplementCMultiPool) RunningByIndex(idx int) (int, error) {
	return 0, errors.New(ErrUnimplemented)
}

// WaitingByIndex is a placeholder method for unimplemented CMultiPool.WaitingByIndex.
func (u *UnimplementCMultiPool) WaitingByIndex(idx int) (int, error) {
	return 0, errors.New(ErrUnimplemented)
}

// CMultiPoolWithFunc defines the interface for a generic multi-worker pool that accepts tasks with arguments.
type CMultiPoolWithFunc[T any] interface {
	CMultiPool
	Invoke(arg T) error
}

// UnimplementCMultiPoolWithFunc is a placeholder struct for unimplemented CMultiPoolWithFunc methods.
type UnimplementCMultiPoolWithFunc[T any] struct {
	*UnimplementCMultiPool
}

// Invoke is a placeholder method for unimplemented CMultiPoolWithFunc.Invoke.
func (u *UnimplementCMultiPoolWithFunc[T]) Invoke(arg T) error {
	return errors.New(ErrUnimplemented)
}
