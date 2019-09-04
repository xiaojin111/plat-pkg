package jwt

import (
	"context"
	"net/http"

	"github.com/micro/go-micro/metadata"
)

const (
	// MetaJwtKey Metadata 中 jwt 的 key.
	MetaJwtKey = "x-jwt"
)

func ContextWithJwt(ctx context.Context, jwt string) context.Context {
	// go 底层源码里面对 Key 传递的时候做了 CanonicalMIMEHeaderKey 处理
	return metadata.NewContext(ctx, map[string]string{
		http.CanonicalHeaderKey(MetaJwtKey): jwt,
	})
}

// JwtFromContext 从 Context 中获取 jwt 的值
func JwtFromContext(ctx context.Context) string {
	jwt := ""
	if md, ok := metadata.FromContext(ctx); ok {
		jwt = md[http.CanonicalHeaderKey(MetaJwtKey)]
	}
	return jwt
}
