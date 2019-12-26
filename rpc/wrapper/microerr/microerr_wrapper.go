package microerr

import (
	"context"
	"net/http"
	"strings"

	"github.com/micro/go-micro"
	me "github.com/micro/go-micro/errors"
	"github.com/micro/go-micro/metadata"
	"github.com/micro/go-micro/server"

	"github.com/jinmukeji/plat-pkg/rpc/errors"
)

// 错误样式
const (
	// ErrStyleSimple 简易模式错误输出样式。通常用于面向最终用户输出，本样式隐藏内部错误信息。
	ErrStyleSimple = "microsimple"
	// ErrStyleDetailed 详细模式错误输出样式。通常用于面向开发者或系统运维人员输出，本样式输出内部错误信息。
	ErrStyleDetailed = "microdetailed"
	// ErrStyleDetailed Micro原始错误输出样式。
	ErrStyleRaw = "raw"
)

const (
	ErrorStyleMetaKey      = "x-err-style"       // Metadata key / HTTP Header
	ApplicationErrorStatus = "Application Error" // 错误状态
	ApplicationErrorCode   = 520                 // 错误状态码
)

func MicroErrWrapper(fn server.HandlerFunc) server.HandlerFunc {

	return func(ctx context.Context, req server.Request, rsp interface{}) error {
		var svcName string
		if svc, ok := micro.FromContext(ctx); ok {
			svcName = svc.Server().Options().Name
		}

		err := fn(ctx, req, rsp)
		if err != nil {
			var rErr error
			style := errorStyleFromContext(ctx)

			switch style {
			case ErrStyleSimple:
				// 使用简化版信息输出，不输出内部错误信息
				if e, ok := err.(*errors.RpcError); ok {
					rErr = wrapError(svcName, e.Error())
				} else {
					// fallback，转为原始的
					rErr = err
				}

			case ErrStyleDetailed:
				// 使用详细版信息输出，输出内部错误信息
				if e, ok := err.(*errors.RpcError); ok {
					rErr = wrapError(svcName, e.DetailedError())
				} else {
					// fallback，转为原始的
					rErr = err
				}
			case ErrStyleRaw:
				// 不做处理，直接输出原始的
				rErr = err
			default:
				// 其它未知的，直接输出原始的
				rErr = err
			}

			return rErr
		}

		return nil
	}
}

func wrapError(id, detail string) error {
	err := &me.Error{
		Id:     id,
		Code:   ApplicationErrorCode,
		Detail: detail,
		Status: ApplicationErrorStatus,
	}

	return err
}

func errorStyleFromContext(ctx context.Context) string {
	style := ErrStyleRaw // default

	if md, ok := metadata.FromContext(ctx); ok {
		// available style:
		//  - RpcError
		//  - MicroSimple
		//  - MicroDetailed
		style = md[http.CanonicalHeaderKey(ErrorStyleMetaKey)]
		style = strings.ToLower(style)
	}

	return style
}
