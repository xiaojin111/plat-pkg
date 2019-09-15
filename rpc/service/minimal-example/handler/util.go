package handler

import (
	"context"

	cm "github.com/jinmukeji/plat-pkg/rpc/ctxmeta"
)

func appIdFromContext(ctx context.Context) string {
	return cm.AppIdFromContext(ctx)
}
