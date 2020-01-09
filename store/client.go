package store

import (
	"github.com/jinmukeji/go-pkg/mysqldb"
	"github.com/micro/go-micro/config"
	"github.com/micro/go-micro/config/reader"
)

// NewDbClientFromConfig 通过 Micro Config 的配置创建 DbClient
func NewDbClientFromConfig(cfgKey ...string) (*mysqldb.DbClient, error) {
	opts := mysqldb.NewOptions()
	if err := config.Get(cfgKey...).Scan(&opts); err != nil {
		return nil, err
	}

	return mysqldb.NewDbClientFromOptions(opts)
}

// NewDbClientFromConfigValue 通过 Micro Config 的 reader.Value 创建 DbClient
func NewDbClientFromConfigValue(v reader.Value) (*mysqldb.DbClient, error) {
	opts := mysqldb.NewOptions()
	if err := v.Scan(&opts); err != nil {
		return nil, err
	}

	return mysqldb.NewDbClientFromOptions(opts)
}
