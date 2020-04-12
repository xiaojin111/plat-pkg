package main

import (
	"fmt"
	"time"

	"github.com/jinmukeji/plat-pkg/v2/dbutil/mysql"
)

type DBUser struct {
	UserID    int64 `gorm:"primary_key;column:user_id"`
	Username  string
	FirstName string
	LastName  string
	CreatedAt time.Time  // 创建时间
	UpdatedAt time.Time  // 更新时间
	DeletedAt *time.Time // 删除时间
}

func (u DBUser) TableName() string {
	return "user"
}

func main() {
	cfg := mysql.NewConfig()
	cfg.User = "root"
	cfg.Passwd = "p@ssw0rd"
	cfg.Net = "tcp"
	cfg.Addr = "127.0.0.1:3306"
	cfg.DBName = "platform"
	cfg.Timeout = 10 * time.Second // Dial timeout

	db, err := mysql.OpenGormDB(
		mysql.WithMySQLConfig(cfg),
	)
	if err != nil {
		panic(err)
	}

	defer db.Close()

	// Read
	var u DBUser
	// find user with id 1
	if err := db.First(&u, 1).Error; err != nil {
		panic(err)
	}

	fmt.Printf("User: %v", u)
}
