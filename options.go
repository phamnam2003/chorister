package chorister

import (
	"time"

	"github.com/phamnam2003/ants/v2"
	"github.com/phamnam2003/chorister/pkg/logs"
)

// Option represents the optional function.
type Option func(opts *Options)

// loadOptions loads options from the given Option functions.
func loadOptions(options ...Option) *Options {
	opts := new(Options)
	for _, option := range options {
		option(opts)
	}
	return opts
}

// COptions represents chorister-specific configuration options.
type COptions struct {
	// EnableMetrics indicates whether to enable metrics collection.
	EnableMetrics bool

	// CLogger to create logs and zap logger used for bridge with log sdk OpenTelemetry.
	CLogger logs.CLogger
}

// Options represents configuration options for the chorister package.
type Options struct {
	// Options from the ants package.
	// It serves as the base configuration for goroutine pool management.
	// Includes settings like logger, expiry duration, panic handler, non-blocking, etc.
	ants.Options

	// COptions represents additional configuration options specific to the chorister package.
	COptions
}

// WithOptions accepts the whole Options config.
func WithOptions(options Options) Option {
	return func(opts *Options) {
		*opts = options
	}
}

// WithExpiryDuration sets up the interval time of cleaning up goroutines.
func WithExpiryDuration(expiryDuration time.Duration) Option {
	return func(opts *Options) {
		opts.ExpiryDuration = expiryDuration
	}
}

// WithPreAlloc indicates whether it should malloc for workers.
func WithPreAlloc(preAlloc bool) Option {
	return func(opts *Options) {
		opts.PreAlloc = preAlloc
	}
}

// WithMaxBlockingTasks sets up the maximum number of goroutines that are blocked when it reaches the capacity of pool.
func WithMaxBlockingTasks(maxBlockingTasks int) Option {
	return func(opts *Options) {
		opts.MaxBlockingTasks = maxBlockingTasks
	}
}

// WithNonblocking indicates that pool will return nil when there is no available workers.
func WithNonblocking(nonblocking bool) Option {
	return func(opts *Options) {
		opts.Nonblocking = nonblocking
	}
}

// WithPanicHandler sets up panic handler.
func WithPanicHandler(panicHandler func(any)) Option {
	return func(opts *Options) {
		opts.PanicHandler = panicHandler
	}
}

// WithAntsLogger sets up a customized logger.
func WithAntsLogger(logger ants.Logger) Option {
	return func(opts *Options) {
		opts.Logger = logger
	}
}

// WithDisablePurge indicates whether we turn off automatically purge.
func WithDisablePurge(disable bool) Option {
	return func(opts *Options) {
		opts.DisablePurge = disable
	}
}

// WithEnableMetrics indicates whether we turn on metrics.
func WithEnableMetrics(enable bool) Option {
	return func(opts *Options) {
		opts.EnableMetrics = enable
	}
}

// WithCLogger sets up a customized zap logger.
func WithCLogger(cLogger logs.CLogger) Option {
	return func(opts *Options) {
		opts.CLogger = cLogger
	}
}
