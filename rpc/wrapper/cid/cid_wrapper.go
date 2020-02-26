package cid

import (
	"context"

	rc "github.com/jinmukeji/plat-pkg/v2/rpc/cid"
	"github.com/micro/go-micro/v2/server"
)

// CidWrapper is a handler wrapper that generate a new cid if cid is not found from request
func CidWrapper(fn server.HandlerFunc) server.HandlerFunc {

	return func(ctx context.Context, req server.Request, rsp interface{}) error {
		cid := rc.CidFromContext(ctx)
		// 如果没有找到 cid，则生成一个新的
		if cid == "" {
			cid = rc.NewCid()
			ctx = rc.ContextWithCid(ctx, cid)
		}
		err := fn(ctx, req, rsp)
		return err
	}
}
