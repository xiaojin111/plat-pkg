package main

import (
	"fmt"
	"time"

	"github.com/jinmukeji/go-pkg/v2/log"
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
	log.SetLevel(log.DebugLevel)

	cfg := mysql.NewStandardConfig()
	cfg.User = "root"
	cfg.Passwd = "p@ssw0rd"
	cfg.Net = "tcp"
	cfg.Addr = "127.0.0.1:3306"
	cfg.DBName = "platform"
	cfg.Timeout = 10 * time.Second // Dial timeout

	db, err := mysql.OpenDB(
		mysql.WithMySQLConfig(cfg),
		mysql.WithLogLevel(log.DebugLevel),
	)
	if err != nil {
		log.Fatalln(err)
	}

	defer db.Close()

	// Read
	var u DBUser
	// find user with id 1
	log.Infoln("reading data")
	if err := db.First(&u, 1).Error; err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("User: %v", u)
}
