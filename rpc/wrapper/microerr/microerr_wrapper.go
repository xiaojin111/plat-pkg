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

const defaultSvcName = "com.jinmuhealth.platform"

func MicroErrWrapper(fn server.HandlerFunc) server.HandlerFunc {

	return func(ctx context.Context, req server.Request, rsp interface{}) error {
		svcName := defaultSvcName
		if svc, ok := micro.FromContext(ctx); ok {
			svcName = svc.Server().Options().Name
		}

		err := fn(ctx, req, rsp)
		if err != nil {
			style := strings.ToLower(errorStyleFromContext(ctx))

			switch style {
			case "microsimple":
				// 使用简化版信息输出，不输出内部错误信息
				if e, ok := err.(*errors.RpcError); ok {
					return wrapError(svcName, e.Error())
				}

			case "microdetailed":
				// 使用详细版信息输出，输出内部错误信息
				if e, ok := err.(*errors.RpcError); ok {
					return wrapError(svcName, e.DetailedError())
				}
			case "raw":
				// 不做处理，直接输出原始的
				// return err
			default:
				// return err
			}

			return err
		}

		return nil
	}
}

func wrapError(id, detail string) error {
	err := &me.Error{
		Id:     id,
		Code:   520,
		Detail: detail,
		Status: "Application Error",
	}

	return err
}

func errorStyleFromContext(ctx context.Context) string {
	style := "Raw" // default

	if md, ok := metadata.FromContext(ctx); ok {
		// available style:
		//  - RpcError
		//  - MicroSimple
		//  - microdetailed
		style = md[http.CanonicalHeaderKey("x-err-style")]
	}

	return style
}
