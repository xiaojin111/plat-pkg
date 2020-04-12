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

// TODO: 删除无用代码

// // Options 是 DbClient 的配置参数
// type Options struct {
// 	// 是否启用日志
// 	EnableLog bool `json:"enable_log" yaml:"enable_log"`

// 	// 连接超时
// 	// The value must be a decimal number with a unit suffix ("ms", "s", "m", "h"), such as "30s", "0.5m" or "1m30s".
// 	DialTimeout string `json:"dial_timeout" yaml:"dial_timeout"`

// 	// 是否转换时间
// 	ParseTime bool `json:"parse_time" yaml:"parse_time"`

// 	// 最大连接数
// 	MaxConnections int `json:"max_connections" yaml:"max_connections"`

// 	// 服务器地址 - host:port
// 	Address string `json:"address" yaml:"address"`

// 	// 用户名
// 	Username string `json:"username" yaml:"username"`

// 	// 密码
// 	Password string `json:"password" yaml:"password"`

// 	// 数据库名
// 	Database string `json:"database" yaml:"database"`

// 	// 字符集
// 	Charset string `json:"charset" yaml:"charset"`

// 	// 字符排序
// 	Collation string `json:"collation" yaml:"collation"`

// 	// 区域设置
// 	Locale string `json:"locale" yaml:"locale"`

// 	// 阻止全局更新
// 	// 开启选项时，将禁止没有 WHERE 语句的 DELETE 或 UPDATE 操作执行，否则抛出 error
// 	BlockGlobalUpdate bool `json:"block_global_update" yaml:"block_global_update"`
// }

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
