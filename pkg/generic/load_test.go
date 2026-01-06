package generic_test

import (
	"testing"

	"github.com/phamnam2003/chorister"
	"github.com/phamnam2003/chorister/pkg/generic"
	"github.com/phamnam2003/chorister/pkg/logs"
	"github.com/stretchr/testify/require"
)

func TestLoadGenericOptions(t *testing.T) {
	opts := generic.LoadGenericOptions[chorister.Options]()
	require.Nil(t, opts.Options.Logger)
	require.Nil(t, opts.Options.PanicHandler)
	require.Zero(t, opts.Options.ExpiryDuration)
	require.Zero(t, opts.Options.MaxBlockingTasks)
	require.False(t, opts.Options.PreAlloc)
	require.False(t, opts.Options.Nonblocking)
	require.False(t, opts.Options.DisablePurge)
	require.False(t, opts.EnableMetrics)
	require.Nil(t, opts.CLogger.Logger)

	opts = generic.LoadGenericOptions(chorister.WithPreAlloc(true), chorister.WithDisablePurge(true))
	require.True(t, opts.Options.PreAlloc)
	require.True(t, opts.Options.DisablePurge)
	require.False(t, opts.EnableMetrics)

	LOG_DIR := "/var/log/chorister"
	lOpts := generic.LoadGenericOptions(logs.WithEnableRotate(true), logs.WithLogDir(LOG_DIR))
	require.True(t, lOpts.EnableRotate)
	require.Equal(t, LOG_DIR, lOpts.LogDir)
	require.Empty(t, lOpts.Prefix)
}
