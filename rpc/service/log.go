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
			Name:    "log_level",
			Usage:   "Log level. [TRACE, DEBUG, INFO, WARN, ERROR, FATAL, PANIC]",
			EnvVars: []string{"LOG_LEVEL"},
			Value:   DefaultLogLevel,
		},
	}
}

func setupLogger(c *cli.Context, svc string) {

	// Setup Log level
	lv := ml.InfoLevel
	logLevel := c.String("log_level")
	if logLevel != "" {
		if level, err := ml.GetLevel(logLevel); err != nil {
			log.Fatal(err.Error())
		} else {
			lv = level
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
	}

	ml.DefaultLogger = mll.NewLogger(
		ml.WithLevel(lv),
		fmtOpt,
	)
}
