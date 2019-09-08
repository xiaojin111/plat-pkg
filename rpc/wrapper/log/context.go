package log

import (
	"context"

	"github.com/sirupsen/logrus"
)

type loggerContextKey string

var defaultLoggerContextKey = loggerContextKey("ctx-logger")

func contextWithLogger(ctx context.Context, l *logrus.Entry) context.Context {
	c := context.WithValue(ctx, defaultLoggerContextKey, l)
	return c
}

func LoggerFromContext(ctx context.Context) *logrus.Entry {
	if l, ok := ctx.Value(defaultLoggerContextKey).(*logrus.Entry); ok {
		return l
	}
	return nil
}
