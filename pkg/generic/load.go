package generic

// Option represents the optional function.
type Option[O any] func(opts *O)

// LoadGenericOptions loads options from the given Option functions.
func LoadGenericOptions[O any](options ...Option[O]) *O {
	opts := new(O)
	for _, option := range options {
		option(opts)
	}
	return opts
}
