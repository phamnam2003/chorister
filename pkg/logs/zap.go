// Package logs provides logging functionalities using the Zap logging library.
package logs

import "go.uber.org/zap"

// CLogger is an alias for zap.Logger
// Because zap logger have bridge with OpenTelemetry log sdk
// this is helpful for centralized logging management
type CLogger *zap.Logger
