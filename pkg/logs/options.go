package logs

import (
	"github.com/phamnam2003/chorister/pkg/generic"
	"go.uber.org/zap"
)

type PrefixLogger struct {
	Key, Value string
}

// LOptions holds configuration options for the logger.
// It can be extended in the future to include various logging settings.
type LOptions struct {
	// Writers is a slice of LogWriter pointers.
	// Each LogWriter encapsulates an io.Writer, a zapcore.Encoder,
	// and a zapcore.LevelEnabler for structured logging.
	Writers []*LogWriter

	// ZapOpts holds additional options for the zap logger.
	// This allows further customization of the zap logger behavior.
	ZapOpts []zap.Option

	// Prefix is the prefix string added to each log entry.
	// It helps in identifying logs from different sources or modules.
	Prefix PrefixLogger
}

// WithPrefix sets the Prefix option for LOptions.
// It adds a specified prefix to each log entry.
func WithPrefix(prefix PrefixLogger) generic.Option[LOptions] {
	return func(opts *LOptions) {
		opts.Prefix = prefix
	}
}

func WithWriters(writers ...*LogWriter) generic.Option[LOptions] {
	return func(opts *LOptions) {
		opts.Writers = writers
	}
}

func WithZapOptions(zapOpts ...zap.Option) generic.Option[LOptions] {
	return func(opts *LOptions) {
		opts.ZapOpts = zapOpts
	}
}
