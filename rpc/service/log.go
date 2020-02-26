package service

import (
	"strings"

	mlog "github.com/jinmukeji/go-pkg/v2/log"
	"github.com/micro/cli/v2"
	"github.com/sirupsen/logrus"
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

func setupLogger(c *cli.Context, logger *mlog.Logger, svc string) {
	// Setup formatter
	logFormat := c.String("log_format")
	if strings.ToLower(logFormat) == "logstash" {
		formatter := mlog.NewLogstashFormatter(logrus.Fields{
			"svc": svc,
		})
		log.SetFormatter(formatter)
	}

	// Setup Log level
	logLevel := c.String("log_level")
	if logLevel != "" {
		if level, err := logrus.ParseLevel(logLevel); err != nil {
			log.Fatal(err.Error())
		} else {
			log.SetLevel(level)
			log.Debugf("Log Level: %s", level)
		}
	}
}
