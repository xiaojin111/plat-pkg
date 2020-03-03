package service

import (
	"strings"

	mlog "github.com/jinmukeji/go-pkg/v2/log"
	"github.com/micro/cli/v2"
	"github.com/sirupsen/logrus"

	ml "github.com/micro/go-micro/v2/logger"
	mll "github.com/micro/go-plugins/logger/logrus/v2"
)

const (
	DefaultLogLevel = "INFO"
)

var (
	// log is the package global logger
	log = mlog.StandardLogger()
)

func logCliFlags() []cli.Flag {
	return []cli.Flag{
		// 日志相关
		&cli.StringFlag{
			Name:    "log_format",
			Usage:   "Log format. Empty string or LOGSTASH.",
			EnvVars: []string{"LOG_FORMAT"},
			Value:   "",
		},

		&cli.StringFlag{
			Name:  "log_level",
			Usage: "Log level. [TRACE, DEBUG, INFO, WARN, ERROR, PANIC, FATAL]",
			// the first environment variable that resolves is used as the default
			EnvVars: []string{"LOG_LEVEL", "MICRO_LOG_LEVEL"},
			Value:   DefaultLogLevel,
		},
	}
}

func setupLogger(c *cli.Context, svc string) {

	// Setup Log level
	lv := ml.InfoLevel
	logLevel := c.String("log_level")
	if logLevel != "" {
		if level, err := ml.GetLevel(strings.ToLower(logLevel)); err != nil {
			log.Fatal(err.Error())
		} else {
			lv = level
			// setup standard logger
			mlog.StandardLogger().SetLevel(loggerToLogrusLevel(lv))
			log.Debugf("Log Level: %s", level)
		}
	}

	// Setup formatter
	logFormat := c.String("log_format")
	fmtOpt := mll.WithTextTextFormatter(mlog.DefaultTextFormatter())
	if strings.ToLower(logFormat) == "logstash" {
		f := mlog.NewLogstashFormatter(logrus.Fields{
			"svc": svc,
		})
		fmtOpt = mll.WithJSONFormatter(f)

		// setup standard logger
		mlog.StandardLogger().SetFormatter(f)
	}

	// Hijack micro's logger
	ml.DefaultLogger = mll.NewLogger(
		ml.WithLevel(lv),
		fmtOpt,
	)
}

func loggerToLogrusLevel(level ml.Level) logrus.Level {
	switch level {
	case ml.TraceLevel:
		return logrus.TraceLevel
	case ml.DebugLevel:
		return logrus.DebugLevel
	case ml.InfoLevel:
		return logrus.InfoLevel
	case ml.WarnLevel:
		return logrus.WarnLevel
	case ml.ErrorLevel:
		return logrus.ErrorLevel
	case ml.FatalLevel:
		return logrus.FatalLevel
	default:
		return logrus.InfoLevel
	}
}
