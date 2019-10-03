package cid

import (
	"context"
	"net/http"

	"github.com/micro/go-micro/metadata"
)

const (
	// MetaCidKey Metadata 中 cid 的 key.
	MetaCidKey = "X-Cid"
)

func ContextWithCid(ctx context.Context, cid string) context.Context {

	md, ok := metadata.FromContext(ctx)
	if !ok {
		md = make(metadata.Metadata)
	} else {
		md = metadata.Copy(md)
	}

	md[http.CanonicalHeaderKey(MetaCidKey)] = cid
	return metadata.NewContext(ctx, md)

	// go 底层源码里面对 Key 传递的时候做了 CanonicalMIMEHeaderKey 处理
	// return metadata.NewContext(ctx, map[string]string{
	// 	http.CanonicalHeaderKey(MetaCidKey): cid,
	// })
}

// CidFromContext 从 Context 中获取 cid 的值
func CidFromContext(ctx context.Context) string {
	cid := ""
	if md, ok := metadata.FromContext(ctx); ok {
		cid = md[http.CanonicalHeaderKey(MetaCidKey)]
	}
	return cid
}
