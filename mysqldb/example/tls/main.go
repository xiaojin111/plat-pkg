package main

import (
	"fmt"
	"log"
	"time"

	"github.com/jinmukeji/plat-pkg/v2/mysqldb"
)

type DBDeveloper struct {
	DeveloperID int64 `gorm:"primary_key;column:developer_id"`
	CompanyName string
	Address     string
	Phone       string
	CreatedAt   time.Time  // 创建时间
	UpdatedAt   time.Time  // 更新时间
	DeletedAt   *time.Time // 删除时间
}

func (d DBDeveloper) TableName() string {
	return "developer"
}

func main() {
	cfg := mysqldb.NewStandardConfig()
	cfg.User = "jmtest"
	cfg.Passwd = "Qg34xCl9vc1F"
	cfg.Net = "tcp"
	cfg.Addr = "jinmu-test.cjzrjn31gtsw.rds.cn-north-1.amazonaws.com.cn:63306"
	cfg.DBName = "platform"
	cfg.Timeout = 10 * time.Second // Dial timeout

	tlsCfg, err := mysqldb.NewTLSConfig("rds-cn-ca-2019-root.pem")
	if err != nil {
		log.Fatal(err)
	}

	db, err := mysqldb.OpenDB(
		mysqldb.WithMySQLConfig(cfg),
		mysqldb.WithTLS(tlsCfg),
	)
	if err != nil {
		panic(err)
	}

	defer db.Close()

	// Read
	var d DBDeveloper
	// find developer with id 1
	log.Println("reading data")
	if err := db.First(&d, 1).Error; err != nil {
		panic(err)
	}

	fmt.Printf("User: %v", d)
}
