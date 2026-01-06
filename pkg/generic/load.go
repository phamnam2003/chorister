package generic

// Option represents the optional function.
type Option[Os any] func(opts *Os)

// LoadGenericOptions loads options from the given Option functions.
func LoadGenericOptions[Os any](options ...Option[Os]) *Os {
	opts := new(Os)
	for _, option := range options {
		option(opts)
	}
	return opts
}
