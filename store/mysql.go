package store

import (
	"context"

	"github.com/jinmukeji/go-pkg/mysqldb"
	tx "github.com/jinmukeji/plat-pkg/transaction"
	"github.com/jinzhu/gorm"
)

// complains compiling error if MySqlStore doesn't implement Tx interface.
var _ tx.Tx = (*MySqlStore)(nil)

type MySqlStore struct {
	client *mysqldb.DbClient
}

func NewMySqlStore(client *mysqldb.DbClient) *MySqlStore {
	return &MySqlStore{
		client: client,
	}
}

// 实现事务控制接口

type txDbKey string

const txDbCtxKey txDbKey = "txDb"

func (s *MySqlStore) DB(ctx context.Context) *gorm.DB {
	if v := ctx.Value(txDbCtxKey); v != nil {
		return v.(*gorm.DB)
	} else {
		return s.client.DB
	}
}

func (s *MySqlStore) BeginTx(ctx context.Context) context.Context {
	db := s.client.DB.Begin() // 从原始的 DB 开启事务
	return context.WithValue(ctx, txDbCtxKey, db)
}

func (s *MySqlStore) CommitTx(ctx context.Context) {
	db := s.DB(ctx)
	db.Commit()
}

func (s *MySqlStore) RollbackTx(ctx context.Context) {
	db := s.DB(ctx)
	db.Rollback()
}

func (s *MySqlStore) GetError(ctx context.Context) error {
	db := s.DB(ctx)
	return db.Error
}
