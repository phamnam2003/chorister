// Package logs provides logging functionalities using the Zap logging library.
package logs

import (
	"errors"
	"sync"

	"github.com/phamnam2003/ants/v2/pkg/syncx"
	"github.com/phamnam2003/chorister/pkg/generic"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	ErrInvalidIOWriters = errors.New("io writers is empty")
)

// CLogger is an alias for zap.Logger
// Because zap logger have bridge with OpenTelemetry log sdk
// this is helpful for centralized logging management
type CLogger struct {
	*zap.Logger

	options *LOptions
	lock    sync.Locker
}

func NewCLogger(opts ...generic.Option[LOptions]) (*CLogger, error) {
	lOpts := generic.LoadGenericOptions(opts...)

	if len(lOpts.RotateWriter) == 0 {
		return nil, ErrInvalidIOWriters
	}

	if lOpts.Encoder == nil {
		lOpts.Encoder = zapcore.NewJSONEncoder(zap.NewDevelopmentEncoderConfig())
	}

	cores := make([]zapcore.Core, 0, len(lOpts.RotateWriter))
	for _, w := range lOpts.RotateWriter {
		cores = append(cores, zapcore.NewCore(
			lOpts.Encoder,
			zapcore.AddSync(w),
			zapcore.InfoLevel,
		))
	}
	logger := zap.New(zapcore.NewTee(cores...), lOpts.ZapOpts...)

	if lOpts.Prefix != "" {
		logger = logger.With(zap.Field{Key: "prefix", Type: zapcore.StringType, String: lOpts.Prefix})
	}

	clog := &CLogger{
		Logger:  logger,
		options: lOpts,
		lock:    syncx.NewSpinLock(),
	}

	return clog, nil
}
