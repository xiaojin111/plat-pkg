package formatmeta

import (
	"context"
	"net/http"

	"github.com/micro/go-micro/v2/server"

	"github.com/micro/go-micro/v2/metadata"
)

// FormatMetadataWrapper is a handler wrapper that format all metadata keys as http.CanonicalHeaderKey.
func FormatMetadataWrapper(fn server.HandlerFunc) server.HandlerFunc {

	return func(ctx context.Context, req server.Request, rsp interface{}) error {

		md, ok := metadata.FromContext(ctx)
		if ok && len(md) > 0 {
			nmd := metadata.Metadata{}
			for k, v := range md {
				nmd[http.CanonicalHeaderKey(k)] = v
			}

			ctx = metadata.NewContext(ctx, nmd)
		}

		err := fn(ctx, req, rsp)
		return err
	}
}
