// Package logs provides logging functionalities using the Zap logging library.
package logs

import (
	"errors"

	"github.com/phamnam2003/chorister/pkg/generic"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	ErrInvalidIOWriters = errors.New("io writers is invalid")
)

// CLogger is an alias for zap.Logger
// Because zap logger have bridge with OpenTelemetry log sdk
// this is helpful for centralized logging management
type CLogger struct {
	*zap.Logger

	options *LOptions
}

// NewCLogger creates a new CLogger instance based on the provided LOptions.
// It configures the logger with specified writers, encoders, and log levels.
// It use zap logger under the hood.
func NewCLogger(opts ...generic.Option[LOptions]) (*CLogger, error) {
	lOpts := generic.LoadGenericOptions(opts...)

	if len(lOpts.Writers) == 0 {
		return nil, ErrInvalidIOWriters
	}

	cores := make([]zapcore.Core, 0, len(lOpts.Writers))
	for _, w := range lOpts.Writers {
		if !w.Valid() {
			return nil, ErrInvalidIOWriters
		}
		cores = append(cores, zapcore.NewCore(
			w.Encoder,
			zapcore.AddSync(w.Writer),
			w.LevelEnabler,
		))
	}
	logger := zap.New(zapcore.NewTee(cores...), lOpts.ZapOpts...)

	if lOpts.Prefix.Key != "" && lOpts.Prefix.Value != "" {
		logger = logger.With(zap.String(lOpts.Prefix.Key, lOpts.Prefix.Value))
	}

	clog := &CLogger{
		Logger:  logger,
		options: lOpts,
	}

	return clog, nil
}

// Sync flushes any buffered log entries.
// It is wrapper around zap.Sync method.
func (cl *CLogger) Sync() error {
	if cl.Logger != nil {
		return cl.Logger.Sync()
	}

	return nil
}
