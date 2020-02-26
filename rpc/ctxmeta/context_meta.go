package ctxmeta

import (
	"context"
	"net/http"

	"github.com/micro/go-micro/v2/metadata"
)

const (
	// MetaJwtKey Metadata 中 jwt 的 key.
	MetaJwtKey = "x-jwt"

	// MetaAppIdKey Metadata 中 APP ID 的 key.
	MetaAppIdKey = "x-app-id"
)

// ContextWithJwt 向 Context (Metadata) 中存入 jwt 的值
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

// ContextWithAppId 向 Context (Metadata) 中存入 APP ID 的值
func ContextWithAppId(ctx context.Context, appId string) context.Context {
	// go 底层源码里面对 Key 传递的时候做了 CanonicalMIMEHeaderKey 处理
	return metadata.NewContext(ctx, map[string]string{
		http.CanonicalHeaderKey(MetaAppIdKey): appId,
	})
}

// AppIdFromContext 从 Context 中获取 APP ID 的值
func AppIdFromContext(ctx context.Context) string {
	appId := ""
	if md, ok := metadata.FromContext(ctx); ok {
		appId = md[http.CanonicalHeaderKey(MetaAppIdKey)]
	}
	return appId
}
