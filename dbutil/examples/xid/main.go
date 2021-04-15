package main

import (
	"fmt"
	"time"

	"gitee.com/jt-heath/plat-pkg/v2/dbutil/mysql"
	"github.com/rs/xid"
)

type DBEvent struct {
	EventID   xid.ID     `gorm:"primary_key;column:event_id"`
	Type      string     `gorm:"column:type"`
	RefID     xid.ID     `gorm:"column:ref_id"`
	CreatedAt time.Time  // 创建时间
	UpdatedAt time.Time  // 更新时间
	DeletedAt *time.Time // 删除时间
}

func (u DBEvent) TableName() string {
	return "event"
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
	id := xid.New()
	e := DBEvent{
		EventID: id,
		Type:    "XXX",
		RefID:   xid.New(),
	}

	// insert
	if err := db.Create(&e).Error; err != nil {
		panic(err)
	}

	// query
	var re DBEvent
	if err := db.Where("event_id = ?", id).First(&re).Error; err != nil {
		panic(err)
	}

	fmt.Printf("Event: %v", re)
}
