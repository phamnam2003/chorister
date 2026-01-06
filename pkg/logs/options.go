package logs

import (
	"github.com/phamnam2003/chorister/pkg/generic"
)

// LOptions holds configuration options for the logger.
// It can be extended in the future to include various logging settings.
type LOptions struct {
	// rotate indicates whether log rotation is enabled.
	// that make logger can be write log file daily, monthly, ...
	EnableRotate bool

	// LogPath is the path log where logs will be stored.
	// if rotate is false, logs will be printed to standard output.
	LogPath string

	// Prefix is the prefix string added to each log entry.
	// It helps in identifying logs from different sources or modules.
	Prefix string
}

// WithEnableRotate sets the EnableRotate option for LOptions.
func WithEnableRotate(enable bool) generic.Option[LOptions] {
	return func(opts *LOptions) {
		opts.EnableRotate = enable
	}
}

// WithLogPath sets the LogDir option for LOptions.
// It joins the provided logPath with the system's file path separator.
// This func makes sure the log path avoid traversal attack.
func WithLogPath(logPath string) generic.Option[LOptions] {
	return func(opts *LOptions) {
		opts.LogPath = logPath
	}
}

// WithPrefix sets the Prefix option for LOptions.
// It adds a specified prefix to each log entry.
func WithPrefix(prefix string) generic.Option[LOptions] {
	return func(opts *LOptions) {
		opts.Prefix = prefix
	}
}
