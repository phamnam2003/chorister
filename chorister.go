package chorister

import (
	"github.com/phamnam2003/ants/v2"
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
func NewChorister(poolSize int, opts ...generic.Option[Options]) (*Chorister, error) {
	cOpts := generic.LoadGenericOptions(opts...)

	pool, err := ants.NewPool(poolSize, ants.WithOptions(cOpts.Options))
	if err != nil {
		return nil, err
	}

	c := &Chorister{
		CPool:   pool,
		options: &cOpts.COptions,
	}

	return c, nil
}
