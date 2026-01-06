// Package logs provides logging functionalities using the Zap logging library.
package logs

import (
	"sync"

	"github.com/phamnam2003/ants/v2/pkg/syncx"
	"github.com/phamnam2003/chorister/pkg/generic"
	"go.uber.org/zap"
)

// CLogger is an alias for zap.Logger
// Because zap logger have bridge with OpenTelemetry log sdk
// this is helpful for centralized logging management
type CLogger struct {
	*zap.Logger

	options *LOptions
	lock    sync.Locker
}

func newCLogger(logger *zap.Logger, opts ...generic.Option[LOptions]) *CLogger {
	lOpts := generic.LoadGenericOptions(opts...)

	clog := &CLogger{
		Logger:  logger,
		options: lOpts,
		lock:    syncx.NewSpinLock(),
	}

	return clog
}
