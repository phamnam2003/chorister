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

	// LogDir is the directory where logs will be stored.
	// if rotate is false, logs will be printed to standard output.
	LogDir string

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

// WithLogPath sets the LogPath option for LOptions.
// It joins the provided logPath with the system's file path separator.
// This func makes sure the log path avoid traversal attack.
func WithLogDir(logDir string) generic.Option[LOptions] {
	return func(opts *LOptions) {
		opts.LogDir = logDir
	}
}

// WithPrefix sets the Prefix option for LOptions.
// It adds a specified prefix to each log entry.
func WithPrefix(prefix string) generic.Option[LOptions] {
	return func(opts *LOptions) {
		opts.Prefix = prefix
	}
}
