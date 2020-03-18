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
	defaultLogLevel = "INFO"
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
			Value:   defaultLogLevel,
		},
	}
}

func setupLogger(c *cli.Context, svc string) {
	std := mlog.StandardLogger()

	// Setup Log level
	lv := mlog.GetLevel()
	if logLevel := c.String("log_level"); logLevel != "" {
		if level, err := logrus.ParseLevel(logLevel); err != nil {
			log.Fatal(err.Error())
		} else {
			lv = level
			// setup standard logger
			std.SetLevel(lv)
		}
	}
	log.Infof("Log Level: %s", lv)

	// Setup formatter
	if logFormat := c.String("log_format"); strings.ToLower(logFormat) == "logstash" {
		// logstash 日式形式下注入 svc 字段，用来输出当前 service 的名称
		f := mlog.NewLogstashFormatter(logrus.Fields{
			"svc": svc,
		})

		std.SetFormatter(f)
	}

	// Hijack micro's logger
	ml.DefaultLogger = mll.NewLogger(
		mll.WithLogger(std),
	)
}
