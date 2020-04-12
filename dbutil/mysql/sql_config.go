package mysql

import (
	"time"

	"github.com/go-sql-driver/mysql"

	"github.com/micro/go-micro/v2/config"
	"github.com/micro/go-micro/v2/config/reader"
)

// StandardConfig 创建一个标准的 *mysql.Config
func NewConfig() *mysql.Config {
	cfg := mysql.NewConfig()

	// 显式设定以下关键参数

	// https://github.com/go-sql-driver/mysql#timetime-support
	cfg.ParseTime = true

	// https://github.com/go-sql-driver/mysql#unicode-support
	cfg.Collation = "utf8mb4_general_ci"

	// https://github.com/go-sql-driver/mysql#loc
	cfg.Loc = time.UTC

	// 默认使用 tcp 网络连接
	cfg.Net = "tcp"

	return cfg
}

type microDbOptions struct {
	// 网络方式, tcp 或者 udp
	Network string `json:"network" yaml:"network"`

	// 服务器地址 - host:port
	Address string `json:"address" yaml:"address"`

	// 用户名
	Username string `json:"username" yaml:"username"`

	// 密码
	Password string `json:"password" yaml:"password"`

	// 数据库名
	Database string `json:"database" yaml:"database"`

	// 连接超时
	// The value must be a decimal number with a unit suffix ("ms", "s", "m", "h"), such as "30s", "0.5m" or "1m30s".
	DialTimeout string `json:"dial_timeout" yaml:"dial_timeout"`

	// 是否转换时间
	ParseTime bool `json:"parse_time" yaml:"parse_time"`

	// 字符排序
	Collation string `json:"collation" yaml:"collation"`

	// 区域设置
	Locale string `json:"locale" yaml:"locale"`
}

// NewDbClientFromConfig 通过 Micro Config 的配置创建 DbClient
func NewConfigFromMicroConfigKey(cfgKey ...string) (*mysql.Config, error) {
	opts := microDbOptions{}
	if err := config.Get(cfgKey...).Scan(&opts); err != nil {
		return nil, err
	}

	return mapConfig(&opts)
}

// NewDbClientFromConfigValue 通过 Micro Config 的 reader.Value 创建 DbClient
func OpenDBFromMicroConfigValue(v reader.Value) (*mysql.Config, error) {
	opts := microDbOptions{}
	if err := v.Scan(&opts); err != nil {
		return nil, err
	}

	return mapConfig(&opts)
}

func mapConfig(opts *microDbOptions) (*mysql.Config, error) {
	cfg := NewConfig()

	if len(opts.Network) > 0 {
		cfg.Net = opts.Network
	}
	cfg.Addr = opts.Address
	cfg.User = opts.Username
	cfg.Passwd = opts.Password
	cfg.DBName = opts.Database

	if len(opts.DialTimeout) > 0 {
		timeout, err := time.ParseDuration(opts.DialTimeout)
		if err != nil {
			return nil, err
		}
		cfg.Timeout = timeout
	}

	cfg.ParseTime = opts.ParseTime

	if len(opts.Collation) > 0 {
		cfg.Collation = opts.Collation
	}

	if len(opts.Locale) > 0 {
		loc, err := time.LoadLocation(opts.Locale)
		if err != nil {
			return nil, err
		}
		cfg.Loc = loc
	}

	return cfg, nil
}
