package log

import (
	"context"
	"time"

	rc "github.com/jinmukeji/plat-pkg/rpc/cid"

	mlog "github.com/jinmukeji/go-pkg/log"
	"github.com/micro/go-micro/server"
)

var (
	// log is the package global logger
	logger = mlog.StandardLogger()
)

const (
	logCidKey     = "cid"
	logLatencyKey = "latency"
	logRpcCallKey = "rpc.call"

	// rpcMetadata   = "[RPC METADATA]"
	rpcFailed = "[RPC ERR]"
	rpcOk     = "[RPC OK]"
)

// LogWrapper is a handler wrapper that logs server request.
func LogWrapper(fn server.HandlerFunc) server.HandlerFunc {

	return func(ctx context.Context, req server.Request, rsp interface{}) error {
		start := time.Now()
		cid := rc.CidFromContext(ctx)

		// 注入一个包含 cid Field 的 logger.Entry
		c := contextWithLogger(ctx, cid)

		err := fn(c, req, rsp)
		// RPC 计算经历的时间长度
		//no time.Since in order to format it well after
		end := time.Now()
		latency := end.Sub(start)

		// l.Infof("%s %s", rpcMetadata, flatMetadata(md))

		l := logger.
			WithField(logCidKey, cid).
			WithField(logRpcCallKey, req.Method()).
			WithField(logLatencyKey, latency.String())

		// Log rpc call execution result
		if err != nil {
			l.WithError(err).Warn(rpcFailed)
		} else {
			l.Info(rpcOk)
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
