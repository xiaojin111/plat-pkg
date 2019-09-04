package service

import (
	"strings"

	mlog "github.com/jinmukeji/go-pkg/log"
	"github.com/micro/cli"
	"github.com/sirupsen/logrus"
)

var (
	// log is the package global logger
	log = mlog.StandardLogger()

	logFormat, logLevel string
)

func logCliFlags() []cli.Flag {
	return []cli.Flag{
		// 日志相关
		cli.StringFlag{
			Name:        "log_format",
			Usage:       "Log format. Empty string or LOGSTASH.",
			EnvVar:      "LOG_FORMAT",
			Value:       "",
			Destination: &logFormat,
		},

		cli.StringFlag{
			Name:        "log_level",
			Usage:       "Log level. [TRACE, DEBUG, INFO, WARN, ERROR, FATAL, PANIC]",
			EnvVar:      "LOG_LEVEL",
			Value:       "INFO",
			Destination: &logLevel,
		},
	}
}

func setupLogger(logger *mlog.Logger, svc string) {
	// Setup formatter
	if strings.ToLower(logFormat) == "logstash" {
		formatter := mlog.NewLogstashFormatter(logrus.Fields{
			"svc": svc,
		})
		log.SetFormatter(formatter)
	}

	// Setup Log level
	if logLevel != "" {
		if level, err := logrus.ParseLevel(logLevel); err != nil {
			log.Fatal(err.Error())
		} else {
			log.SetLevel(level)
		}
	}
}
