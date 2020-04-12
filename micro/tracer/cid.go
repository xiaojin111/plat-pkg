package tracer

import (
	id "github.com/jinmukeji/go-pkg/v2/id-gen"
)

// NewCid 生成一个新的 cid. 使用 xid 形式。
func NewCid() string {
	cid := id.NewXid()
	return cid
}
