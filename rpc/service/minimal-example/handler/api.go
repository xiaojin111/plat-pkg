package handler

import (
	"context"
	"fmt"

	"github.com/jinmukeji/plat-pkg/rpc/errors"
	"github.com/jinmukeji/plat-pkg/rpc/errors/codes"

	echopb "github.com/jinmukeji/proto/gen/micro/idl/examples/echo/v1"
)

type EchoAPIService struct{}

func (svc *EchoAPIService) Hello(ctx context.Context, req *echopb.HelloRequest, rsp *echopb.HelloResponse) error {
	rsp.Greeting = "Hello " + req.Name
	return errors.ErrorWithCause(codes.Internal, fmt.Errorf("eeeeee"), "ppppp")
}

func (svc *EchoAPIService) GetUser(ctx context.Context, req *echopb.GetUserRequest, rsp *echopb.GetUserResponse) error {
	return nil
}

func (svc *EchoAPIService) ModifyUserProfile(ctx context.Context, req *echopb.ModifyUserProfileRequest, rsp *echopb.ModifyUserProfileResponse) error {
	return nil
}
