package mysqldb

import (
	"fmt"
	"time"

	// import mysql driver fo gorm
	"github.com/jinmukeji/plat-pkg/v2/mysqldb/gormlogger"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// DbClient 是数据访问管理器
type DbClient struct {
	*gorm.DB
	opts Options
}

func Open(options Options) (*gorm.DB, error) {
	// mysql 连接字符串格式:
	// 	`username:password@tcp(localhost:3306)/db_name?charset=utf8mb4&collation=utf8mb4_general_ci&parseTime=True&loc=utc&timeout=10s`
	connStr := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s&collation=%s&parseTime=%t&loc=%s&timeout=%s",
		options.Username,
		options.Password,
		options.Address,
		options.Database,
		options.Charset,
		options.Collation,
		options.ParseTime,
		options.Locale,
		options.DialTimeout)

	db, err := gorm.Open("mysql", connStr)
	if err != nil {
		return nil, err
	}

	// gorm setting
	db.SingularTable(true)
	db.DB().SetMaxOpenConns(options.MaxConnections)
	db.SetLogger(gormlogger.New(options.Address, options.Database))
	db = db.LogMode(options.EnableLog)
	// 禁止没有 WHERE 语句的 DELETE 或 UPDATE 操作执行，否则抛出 error
	db = db.BlockGlobalUpdate(options.BlockGlobalUpdate)
	// 重置 SetNow 的时间获取方式为总是获取UTC时区时间
	db = db.SetNowFuncOverride(func() time.Time {
		return time.Now().UTC()
	})

	return db, nil
}

// NewDbClient 根据传入的 options 返回一个新的 DbClient
func NewDbClient(opts ...Option) (*DbClient, error) {
	options := NewOptions(opts...)
	return NewDbClientFromOptions(options)
}

func NewDbClientFromOptions(options Options) (*DbClient, error) {
	db, err := Open(options)
	if err != nil {
		return nil, err
	}

	return &DbClient{db, options}, nil
}

// Options 返回 DbClient 的 Options.
func (c *DbClient) Options() Options {
	return c.opts
}
