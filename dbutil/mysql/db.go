package mysql

import (
	"time"

	// import mysql driver fo gorm
	"github.com/go-sql-driver/mysql"
	"github.com/jinmukeji/plat-pkg/v2/dbutil/gormlogger"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type DB = gorm.DB

const (
	tlsKey = "custom"
)

// OpenDB is alias of OpenGormDB.
func OpenDB(opt ...Option) (*DB, error) {
	return OpenGormDB(opt...)
}

// OpenGormDB 打开一个 *gorm.DB 的连接.
func OpenGormDB(opt ...Option) (*DB, error) {
	options := NewOptions(opt...)

	mysqlCfg := options.MySqlCfg

	if options.TLSCfg != nil {
		err := mysql.RegisterTLSConfig(tlsKey, options.TLSCfg)
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
	db.SetLogger(gormlogger.NewWithLevel(mysqlCfg.Addr, mysqlCfg.DBName, options.LogLevel))
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
