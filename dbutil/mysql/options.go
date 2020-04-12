package mysql

import (
	"crypto/tls"

	"github.com/go-sql-driver/mysql"
	"github.com/jinmukeji/go-pkg/v2/log"
)

type Options struct {
	MySqlCfg *mysql.Config
	TLSCfg   *tls.Config
	LogLevel log.Level
}

// Option 是设置 Options 的函数
type Option func(*Options)

func WithMySQLConfig(cfg *mysql.Config) Option {
	return func(o *Options) {
		o.MySqlCfg = cfg
	}
}

func WithTLS(cfg *tls.Config) Option {
	return func(o *Options) {
		o.TLSCfg = cfg
	}
}

func WithLogLevel(lv log.Level) Option {
	return func(o *Options) {
		o.LogLevel = lv
	}
}

// defaultOptions 返回默认配置的 Options
func defaultOptions() *Options {
	return &Options{
		LogLevel: log.GetLevel(),
	}
}

// NewOptions 设置 Options
func NewOptions(opt ...Option) *Options {
	opts := defaultOptions()

	for _, o := range opt {
		o(opts)
	}

	return opts
}
