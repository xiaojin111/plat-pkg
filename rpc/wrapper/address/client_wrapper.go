package address

import (
	"context"

	"github.com/micro/go-micro/client"
)

type AddressWrapper struct {
	client.Client
	address string
}

func (w *AddressWrapper) Call(ctx context.Context, req client.Request, rsp interface{}, opts ...client.CallOption) error {
	if len(w.address) > 0 {
		o := []client.CallOption{}
		if len(opts) > 0 {
			o = append(o, opts...)
		}
		o = append(o, client.WithAddress(w.address))

		return w.Client.Call(ctx, req, rsp, o...)
	}

	return w.Client.Call(ctx, req, rsp, opts...)
}
func NewAddressWrapper(address string) client.Wrapper {
	return func(c client.Client) client.Client {
		return &AddressWrapper{c, address}
	}
}
