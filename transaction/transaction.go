package transaction

import "context"

// Tx 事务处理方法接口
type Tx interface {
	// 开启事务，并返回新的 Context
	BeginTx(ctx context.Context) context.Context

	// 提交事务
	CommitTx(ctx context.Context)

	// 回滚事务
	RollbackTx(ctx context.Context)

	// 获取 error
	GetError(ctx context.Context) error
}
