package handler

import (
	"context"

	user "github.com/wheelergeo/g-otter-gen/user"
	"github.com/wheelergeo/g-otter-micro-user/biz/service"
)

// UserServiceImpl implements the last service interface defined in the IDL.
type UserServiceImpl struct{}

// RpcWebLogin implements the UserServiceImpl interface.
func (s *UserServiceImpl) RpcWebLogin(ctx context.Context, req *user.RpcWebLoginReq) (resp *user.RpcWebLoginResp, err error) {
	resp, err = service.NewRpcWebLoginService(ctx).Run(req)

	return resp, err
}

// RpcUpdateUserOnline implements the UserServiceImpl interface.
func (s *UserServiceImpl) RpcUpdateUserOnline(ctx context.Context, req *user.RpcUpdateUserOnlineReq) (resp *user.RpcUpdateUserOnlineResp, err error) {
	resp, err = service.NewRpcUpdateUserOnlineService(ctx).Run(req)

	return resp, err
}
