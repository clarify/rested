package resource

import (
	"context"
	"fmt"
	"log"
)

// LogLevel defines log levels
type LogLevel int

// Log levels
const (
	LogLevelDebug LogLevel = iota
	LogLevelInfo
	LogLevelWarn
	LogLevelError
	LogLevelFatal
)

// LoggerLevel sets the logging level of the framework.
var LoggerLevel = LogLevelInfo

// Logger is the function used by rest-layer to log messages. By default
// it does nothing but you can customize it to plug any logger.
var Logger = func(ctx context.Context, level LogLevel, msg string, fields map[string]interface{}) {
	_ = log.Output(2, msg)
}

func logPanicf(ctx context.Context, format string, a ...interface{}) {
	msg := fmt.Sprintf(format, a...)
	if LoggerLevel <= LogLevelFatal && Logger != nil {
		Logger(ctx, LogLevelFatal, msg, nil)
	}
	panic(msg)
}
