package jwt

import (
	"time"

	"github.com/jinmukeji/plat-pkg/rpc/jwt"
)

type Options struct {
	Enabled         bool          // 是否启用
	MaxExpInterval  time.Duration // 最大过期时间间隔
	HeaderKey       string        // HTTP Request Header 中的 jwt 使用的 key
	MicroConfigPath string        // Micro Config 中的 key
}

const (
	defaultMaxExpInterval  = 10 * time.Minute // 10分钟
	defaultMicroConfigPath = "platform/app-key"
)

func DefaultOptions() Options {
	return Options{
		Enabled:         false,
		MaxExpInterval:  defaultMaxExpInterval,
		HeaderKey:       jwt.MetaJwtKey,
		MicroConfigPath: defaultMicroConfigPath,
	}
}
