package chorister

import (
	"github.com/phamnam2003/chorister/pkg/generic"
)

const (
	ErrUnimplemented = "unimplemented method"
)

type Chorister struct {
	// cpool is the underlying goroutine pool used for managing concurrent tasks.
	CPool

	// options holds the configuration options for the Chorister.
	options *COptions
}

// NewChorister creates a new Chorister instance with the given pool size and options.
func NewChorister(pool CPool, opts ...generic.Option[COptions]) (*Chorister, error) {
	options := generic.LoadGenericOptions(opts...)

	c := &Chorister{
		CPool:   pool,
		options: options,
	}

	return c, nil
}
