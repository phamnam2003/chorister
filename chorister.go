package chorister

import (
	"errors"

	"github.com/phamnam2003/ants/v2"
	"github.com/phamnam2003/chorister/pkg/generic"
)

var (
	ErrUnimplemented   = errors.New("unimplemented method")
	ErrNoConfigurePool = errors.New("no configure")
)

type Chorister struct {
	// cpool is the underlying goroutine pool used for managing concurrent tasks.
	CPool

	// options holds the configuration options for the Chorister.
	options *Options
}

// NewChorister creates a new Chorister instance with the given pool size and options.
func NewChorister(poolSize int, opts ...generic.Option[Options]) (*Chorister, error) {
	options := generic.LoadGenericOptions(opts...)

	pool, err := ants.NewPool(poolSize, ants.WithOptions(options.Options))
	if err != nil {
		return nil, err
	}

	c := &Chorister{
		CPool:   pool,
		options: options,
	}

	return c, nil
}
