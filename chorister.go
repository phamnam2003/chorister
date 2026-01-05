package chorister

import "github.com/phamnam2003/ants/v2"

const (
	ErrUnimplemented = "unimplemented method"
)

type Chorister struct {
	// cpool is the underlying goroutine pool used for managing concurrent tasks.
	cpool CPool

	// options holds the configuration options for the Chorister.
	options *Options
}

// NewChorister creates a new Chorister instance with the given pool size and options.
func NewChorister(poolSize int, opts ...Option) (*Chorister, error) {
	cOpts := loadOptions(opts...)

	pool, err := ants.NewPool(poolSize, ants.WithOptions(cOpts.Options))
	if err != nil {
		return nil, err
	}

	c := &Chorister{
		cpool:   pool,
		options: cOpts,
	}

	return c, nil
}
