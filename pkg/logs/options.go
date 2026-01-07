package logs

import (
	"io"

	"github.com/phamnam2003/chorister/pkg/generic"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// LOptions holds configuration options for the logger.
// It can be extended in the future to include various logging settings.
type LOptions struct {
	// RotateWriter is a list of io.Writer where rotated logs will be written.
	// This allows flexibility in directing rotated logs to different destinations.
	RotateWriter []io.Writer

	// Encoder defines the format of the log entries.
	// It can be customized to use different encoders like JSON, console, etc.
	Encoder zapcore.Encoder

	// ZapOpts holds additional options for the zap logger.
	// This allows further customization of the zap logger behavior.
	ZapOpts []zap.Option

	// Prefix is the prefix string added to each log entry.
	// It helps in identifying logs from different sources or modules.
	Prefix string

	// EnableSampling indicates whether log sampling is enabled.
	// When enabled, it reduces the volume of logs by sampling similar log entries when your application is under high load.
	EnableSampling bool
}

// WithPrefix sets the Prefix option for LOptions.
// It adds a specified prefix to each log entry.
func WithPrefix(prefix string) generic.Option[LOptions] {
	return func(opts *LOptions) {
		opts.Prefix = prefix
	}
}

// WithEnableSampling sets the EnableSampling option for LOptions.
func WithEnableSampling(enable bool) generic.Option[LOptions] {
	return func(opts *LOptions) {
		opts.EnableSampling = enable
	}
}

func WithRotateWriters(rotateWriter ...io.Writer) generic.Option[LOptions] {
	return func(opts *LOptions) {
		opts.RotateWriter = rotateWriter
	}
}

func WithZapOptions(zapOpts ...zap.Option) generic.Option[LOptions] {
	return func(opts *LOptions) {
		opts.ZapOpts = zapOpts
	}
}
func WithEncoder(encoder zapcore.Encoder) generic.Option[LOptions] {
	return func(opts *LOptions) {
		opts.Encoder = encoder
	}
}
