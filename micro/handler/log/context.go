package log

import (
	"context"

	"github.com/micro/go-micro/v2/logger"
)

type loggerContextKey string

var defaultLoggerContextKey = loggerContextKey("ctx-logger")

func contextWithLogger(ctx context.Context, hl *logger.Helper) context.Context {
	c := context.WithValue(ctx, defaultLoggerContextKey, hl)
	return c
}

func LoggerFromContext(ctx context.Context) *logger.Helper {
	if hl, ok := ctx.Value(defaultLoggerContextKey).(*logger.Helper); ok {
		return hl
	}
	return nil
}
