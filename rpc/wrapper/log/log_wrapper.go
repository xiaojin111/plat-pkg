package log

import (
	"context"
	"time"

	rc "github.com/jinmukeji/plat-pkg/v2/rpc/cid"
	"github.com/jinmukeji/plat-pkg/v2/rpc/errors"

	"github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-micro/v2/server"
)

const (
	logCidKey     = "cid"
	logLatencyKey = "latency"
	logRpcCallKey = "rpc.call"

	// rpcMetadata   = "[RPC METADATA]"
	rpcFailed = "[RPC ERR]"
	rpcOk     = "[RPC OK]"

	errorField     = "error"
	errorCodeField = "errcode"
)

func helperLogger() *logger.Helper {
	dl, ok := logger.DefaultLogger.(*logger.Helper)
	if !ok {
		return logger.NewHelper(logger.DefaultLogger)
	}

	return dl
}

// LogWrapper is a handler wrapper that logs server request.
func LogWrapper(fn server.HandlerFunc) server.HandlerFunc {

	return func(ctx context.Context, req server.Request, rsp interface{}) error {
		start := time.Now()
		cid := rc.CidFromContext(ctx)

		// 注入一个包含 cid Field 的 logger.Entry
		hl := helperLogger()
		cl := hl.WithFields(map[string]interface{}{logCidKey: cid})
		c := logger.NewContext(ctx, cl)

		err := fn(c, req, rsp)
		// RPC 计算经历的时间长度
		//no time.Since in order to format it well after
		end := time.Now()
		latency := end.Sub(start)

		// l.Infof("%s %s", rpcMetadata, flatMetadata(md))

		l := cl.WithFields(map[string]interface{}{
			logRpcCallKey: req.Method(),
			logLatencyKey: latency.String(),
		})

		// Log rpc call execution result
		switch v := err.(type) {
		case nil:
			l.Info(rpcOk)
		case *errors.RpcError:
			l.WithFields(map[string]interface{}{
				errorField:     v.DetailedError(),
				errorCodeField: v.Code,
			}).
				Warn(rpcFailed)
		case error:
			l.WithFields(map[string]interface{}{errorField: err.Error()}).
				Warn(rpcFailed)
		default:
			l.Errorf("unknown error type: %v", v)
		}

		return err
	}
}

// flatMetadata 将 Metadata 打平为 "k=v" 形式的字符串序列
// func flatMetadata(md metadata.Metadata) string {
// 	var buffer bytes.Buffer
// 	for k, v := range md {
// 		buffer.WriteString(strconv.Quote(k))
// 		buffer.WriteString("=")
// 		buffer.WriteString(strconv.Quote(v))
// 		buffer.WriteString(" ")
// 	}

// 	return buffer.String()
// }
