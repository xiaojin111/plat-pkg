package rpc

import (
	"context"
	"net/http"

	"github.com/micro/go-micro/metadata"
)

const (
	// MetaCidKey Metadata 中 cid 的 key.
	MetaCidKey = "x-cid"
)

func ContextWithCid(ctx context.Context, cid string) context.Context {
	// go 底层源码里面对 Key 传递的时候做了 CanonicalMIMEHeaderKey 处理
	return metadata.NewContext(ctx, map[string]string{
		http.CanonicalHeaderKey(MetaCidKey): cid,
	})
}

// CidFromContext 从 Context 中获取 cid 的值
func CidFromContext(ctx context.Context) string {
	cid := ""
	if md, ok := metadata.FromContext(ctx); ok {
		cid = md[http.CanonicalHeaderKey(MetaCidKey)]
	}
	return cid
}
