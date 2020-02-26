package store

import (
	"context"
	"fmt"

	tx "github.com/jinmukeji/plat-pkg/v2/transaction"
	"github.com/jinmukeji/plat-pkg/v2/mysqldb"
	"github.com/jinzhu/gorm"
)

// complains compiling error if MySqlStore doesn't implement interfaces.
var _ tx.Tx = (*MySqlStore)(nil)
var _ Closer = (*MySqlStore)(nil)

type MySqlStore struct {
	client *mysqldb.DbClient
}

func NewMySqlStore(client *mysqldb.DbClient) *MySqlStore {
	return &MySqlStore{
		client: client,
	}
}

// 实现 Closer 接口

func (s *MySqlStore) Close() error {
	if s.client != nil {
		return s.client.Close()
	}
	return nil
}

// 实现事务控制接口 tx.Tx

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

// Transaction start a transaction as a block,
// return error will rollback, otherwise to commit.
func (s *MySqlStore) Transaction(ctx context.Context, fc func(txs *MySqlStore) error) (err error) {
	tx := s.BeginTx(ctx)
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("%s", r)
			s.RollbackTx(tx)
			return
		}
	}()

	err = fc(s)
	if err == nil {
		s.CommitTx(tx)
		err = s.GetError(tx)
	}

	// Makesure rollback when Block error or Commit error
	if err != nil {
		s.RollbackTx(tx)
	}

	return
}
