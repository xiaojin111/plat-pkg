package log

import (
	"context"

	"github.com/sirupsen/logrus"
)

type loggerContextKey string

var defaultLoggerContextKey = loggerContextKey("ctx-logger")

func contextWithLogger(ctx context.Context, cid string) context.Context {
	c := context.WithValue(ctx, defaultLoggerContextKey, logger.WithField(logCidKey, cid))
	return c
}

func LoggerFromContext(ctx context.Context) *logrus.Entry {
	if logger, ok := ctx.Value(defaultLoggerContextKey).(*logrus.Entry); ok {
		return logger
	}
	return nil
}
