package mysql

import (
	"time"

	"github.com/go-sql-driver/mysql"
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

	return cfg
}
