package logs

import (
	"io"

	"go.uber.org/zap/zapcore"
)

// LogWriter is a structure that encapsulates an io.Writer along with
// a zapcore.Encoder and a zapcore.LevelEnabler to facilitate structured
// logging with level-based filtering.
type LogWriter struct {
	zapcore.Core

	// Writer is the underlying io.Writer where log entries will be written.
	// This could be a file, JSON, console, network connection, etc.
	Writer io.Writer

	// Encoder defines how log entries are formatted.
	// This value allows customization of log output format.
	Encoder zapcore.Encoder

	// LevelEnabler determines which log levels are enabled for this writer.
	// This allows for level-based filtering of log entries.
	LevelEnabler zapcore.LevelEnabler
}

// Valid checks if the LogWriter has all necessary components properly set.
// It returns true if Writer, Encoder, and LevelEnabler are all non-nil.
func (lw *LogWriter) Valid() bool {
	return lw != nil && lw.Writer != nil && lw.Encoder != nil && lw.LevelEnabler != nil
}
