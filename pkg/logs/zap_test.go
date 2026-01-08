package logs_test

import (
	"os"
	"testing"

	"github.com/phamnam2003/chorister/pkg/logs"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func TestLogger(t *testing.T) {
	var (
		logger *logs.CLogger
		err    error
	)

	logger, err = logs.NewCLogger(
		logs.WithPrefix(logs.PrefixLogger{Key: "env", Value: "test"}),
		logs.WithZapOptions(zap.AddCaller(), zap.AddStacktrace(zap.ErrorLevel)),
	)
	require.Error(t, err)
	require.ErrorIs(t, err, logs.ErrInvalidIOWriters)
	require.Nil(t, logger)

	logger, err = logs.NewCLogger(
		logs.WithPrefix(logs.PrefixLogger{Key: "env", Value: "test"}),
		logs.WithZapOptions(zap.AddCaller(), zap.AddStacktrace(zap.ErrorLevel)),
		logs.WithWriters(nil),
	)
	require.Error(t, err)
	require.ErrorIs(t, err, logs.ErrInvalidIOWriters)
	require.Nil(t, logger)

	logger, err = logs.NewCLogger(
		logs.WithPrefix(logs.PrefixLogger{Key: "env", Value: "test"}),
		logs.WithZapOptions(zap.AddCaller(), zap.AddStacktrace(zap.ErrorLevel)),
		logs.WithWriters(&logs.LogWriter{
			Writer:       os.Stdout,
			Encoder:      zapcore.NewJSONEncoder(zap.NewProductionConfig().EncoderConfig),
			LevelEnabler: zap.InfoLevel,
		}),
	)
	require.NoError(t, err)
	require.NotNil(t, logger)

	logger, err = logs.NewCLogger(
		logs.WithPrefix(logs.PrefixLogger{Key: "env", Value: "test"}),
		logs.WithZapOptions(zap.AddCaller(), zap.AddStacktrace(zap.ErrorLevel)),
		logs.WithWriters(&logs.LogWriter{
			Writer:  os.Stdout,
			Encoder: zapcore.NewJSONEncoder(zap.NewProductionConfig().EncoderConfig),
			// LevelEnabler: zap.InfoLevel,
		}),
	)
	require.Error(t, err)
	require.ErrorIs(t, err, logs.ErrInvalidIOWriters)
	require.Nil(t, logger)
}
