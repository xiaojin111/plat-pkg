package mysqldb

import (
	"crypto/tls"
	"crypto/x509"
	"errors"
	"io/ioutil"
	"time"

	// import mysql driver fo gorm
	"github.com/go-sql-driver/mysql"
	"github.com/jinmukeji/plat-pkg/v2/mysqldb/gormlogger"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// DbClient 是数据访问管理器
// type DbClient struct {
// 	*gorm.DB
// 	opts Options
// }

// StandardConfig 创建一个标准的 *mysql.Config
func NewStandardConfig() *mysql.Config {
	cfg := mysql.NewConfig()

	// 显式设定以下关键参数

	// https://github.com/go-sql-driver/mysql#timetime-support
	cfg.ParseTime = true

	// https://github.com/go-sql-driver/mysql#unicode-support
	cfg.Collation = "utf8mb4_general_ci"

	// https://github.com/go-sql-driver/mysql#loc
	cfg.Loc = time.UTC

	return cfg
}

func NewTLSConfig(rootCert string) (*tls.Config, error) {
	rootCertPool := x509.NewCertPool()
	pem, err := ioutil.ReadFile(rootCert)
	if err != nil {
		return nil, err
	}
	if ok := rootCertPool.AppendCertsFromPEM(pem); !ok {
		return nil, errors.New("Failed to append PEM.")
	}

	return &tls.Config{
		RootCAs: rootCertPool,
	}, nil
}

type DbClient  = gorm.DB

func OpenDB(opt ...Option) (*gorm.DB, error) {
	options := NewOptions(opt...)

	mysqlCfg := options.mysqlCfg

	if options.tlsCfg != nil {
		err := mysql.RegisterTLSConfig(tlsKey, options.tlsCfg)
		if err != nil {
			return nil, err
		}

		mysqlCfg.TLSConfig = tlsKey
	}

	dsn := mysqlCfg.FormatDSN()

	db, err := gorm.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	// gorm setting
	db.SingularTable(true)
	// db.DB().SetMaxOpenConns(options.MaxConnections)
	db.SetLogger(gormlogger.NewWithLevel(mysqlCfg.Addr, mysqlCfg.DBName, options.logLevel))
	db = db.LogMode(true)
	// 禁止没有 WHERE 语句的 DELETE 或 UPDATE 操作执行，否则抛出 error
	db = db.BlockGlobalUpdate(true)
	// 重置 SetNow 的时间获取方式为总是获取UTC时区时间
	db = db.SetNowFuncOverride(func() time.Time {
		return time.Now().UTC()
	})

	return db, nil
}

// TODO: 删除无用代码

// // NewDbClient 根据传入的 options 返回一个新的 DbClient
// func NewDbClient(opts ...Option) (*DbClient, error) {
// 	options := NewOptions(opts...)
// 	return NewDbClientFromOptions(options)
// }

// func NewDbClientFromOptions(options Options) (*DbClient, error) {
// 	db, err := Open(options)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return &DbClient{db, options}, nil
// }

// // Options 返回 DbClient 的 Options.
// func (c *DbClient) Options() Options {
// 	return c.opts
// }
