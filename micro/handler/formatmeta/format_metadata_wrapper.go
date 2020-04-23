package formatmeta

import (
	"context"

	"github.com/micro/go-micro/v2/server"

	pm "github.com/jinmukeji/plat-pkg/v2/micro/meta"
	"github.com/micro/go-micro/v2/metadata"
)

// FormatMetadataWrapper is a handler wrapper that format all metadata keys. [Deprecated]
func FormatMetadataWrapper(fn server.HandlerFunc) server.HandlerFunc {

	return func(ctx context.Context, req server.Request, rsp interface{}) error {

		md, ok := metadata.FromContext(ctx)
		if ok && len(md) > 0 {
			nmd := metadata.Metadata{}
			for k, v := range md {
				nmd[pm.StandardizeKey(k)] = v
			}

			ctx = metadata.NewContext(ctx, nmd)
		}

		err := fn(ctx, req, rsp)
		return err
	}
}
