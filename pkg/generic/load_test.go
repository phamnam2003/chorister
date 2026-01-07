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
	require.Nil(t, opts.Logger)
	require.Nil(t, opts.PanicHandler)
	require.Zero(t, opts.ExpiryDuration)
	require.Zero(t, opts.MaxBlockingTasks)
	require.False(t, opts.PreAlloc)
	require.False(t, opts.Nonblocking)
	require.False(t, opts.DisablePurge)
	require.False(t, opts.EnableMetrics)
	require.Nil(t, opts.CLogger.Logger)

	opts = generic.LoadGenericOptions(chorister.WithPreAlloc(true), chorister.WithDisablePurge(true))
	require.True(t, opts.PreAlloc)
	require.True(t, opts.DisablePurge)
	require.False(t, opts.EnableMetrics)

	lOpts := generic.LoadGenericOptions(logs.WithEnableRotate(true), logs.WithEnableSampling(true))
	require.True(t, lOpts.EnableRotate)
	require.Empty(t, lOpts.Prefix)
	require.True(t, lOpts.EnableSampling)
}
