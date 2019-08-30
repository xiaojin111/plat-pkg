package log

import (
	"net/http"

	"github.com/micro/cli"
	"github.com/micro/micro/plugin"

	"strings"

	mlog "github.com/jinmukeji/go-pkg/log"
	"github.com/sirupsen/logrus"
)

type logPlugin struct {
	svcName             string
	logFormat, logLevel string
}

var (
	log *mlog.Logger
)

func init() {
	log = mlog.StandardLogger()
}

func (p *logPlugin) Flags() []cli.Flag {
	return []cli.Flag{
		// 日志相关
		cli.StringFlag{
			Name:        "log_format",
			Usage:       "Log format. Empty string or LOGSTASH.",
			EnvVar:      "LOG_FORMAT",
			Value:       "",
			Destination: &(p.logFormat),
		},

		cli.StringFlag{
			Name:        "log_level",
			Usage:       "Log level. [TRACE, DEBUG, INFO, WARN, ERROR, FATAL, PANIC]",
			EnvVar:      "LOG_LEVEL",
			Value:       "INFO",
			Destination: &(p.logLevel),
		},
	}
}

func (p *logPlugin) Commands() []cli.Command {
	return nil
}

func (p *logPlugin) Handler() plugin.Handler {
	return func(h http.Handler) http.Handler {
		// 什么都不包装，透传
		return h
	}
}

func (p *logPlugin) Init(ctx *cli.Context) error {
	// Setup formatter
	if strings.ToLower(p.logFormat) == "logstash" {
		formatter := mlog.NewLogstashFormatter(logrus.Fields{
			"svc": p.svcName,
		})
		log.SetFormatter(formatter)
	}

	// Setup Log level
	if p.logLevel != "" {
		if level, err := logrus.ParseLevel(p.logLevel); err != nil {
			log.Fatal(err.Error())
		} else {
			log.SetLevel(level)
		}
	}
	return nil
}

func (p *logPlugin) String() string {
	return "log"
}

func NewPlugin(svcName string) plugin.Plugin {
	return NewLog(svcName)
}

func NewLog(svcName string) plugin.Plugin {
	return &logPlugin{
		svcName: svcName,
	}
}
