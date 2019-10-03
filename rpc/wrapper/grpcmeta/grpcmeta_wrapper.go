package cid

import (
	"context"
	"net/http"
	"strings"

	"github.com/micro/go-micro/server"

	"github.com/micro/go-micro/metadata"
	gmetadata "google.golang.org/grpc/metadata"
)

var excluded = map[string]bool{
	// "accept-encoding":true,
}

// GrpcMetadataWrapper is a handler wrapper that map gRPC metadata to Micro metadata
func GrpcMetadataWrapper(fn server.HandlerFunc) server.HandlerFunc {

	return func(ctx context.Context, req server.Request, rsp interface{}) error {

		gmd, ok := gmetadata.FromIncomingContext(ctx)
		if ok {
			md, _ := metadata.FromContext(ctx)
			for k, v := range gmd {
				if !isExcluded(k) {
					md[http.CanonicalHeaderKey(k)] = combineValues(v)
				}
			}
		}

		err := fn(ctx, req, rsp)
		return err
	}
}

func isExcluded(key string) bool {
	if v, ok := excluded[strings.ToLower(key)]; ok {
		return v
	}

	return false
}

func combineValues(v []string) string {
	return strings.Join(v, ",")
}
