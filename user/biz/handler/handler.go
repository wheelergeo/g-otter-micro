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

// RpcUserCommonDeptCreate implements the UserServiceImpl interface.
func (s *UserServiceImpl) RpcUserCommonDeptCreate(ctx context.Context, req *user.RpcUserCommonDeptCreateReq) (resp *user.RpcUserCommonDeptCreateResp, err error) {
	resp, err = service.NewRpcUserCommonDeptCreateService(ctx).Run(req)

	return resp, err
}

// RpcUserCommonDeptRetrieve implements the UserServiceImpl interface.
func (s *UserServiceImpl) RpcUserCommonDeptRetrieve(ctx context.Context, req *user.RpcUserCommonDeptRetrieveReq) (resp *user.RpcUserCommonDeptRetrieveResp, err error) {
	resp, err = service.NewRpcUserCommonDeptRetrieveService(ctx).Run(req)

	return resp, err
}

// RpcUserCommonDeptRetrieveTree implements the UserServiceImpl interface.
func (s *UserServiceImpl) RpcUserCommonDeptRetrieveTree(ctx context.Context, req *user.RpcUserCommonDeptRetrieveTreeReq) (resp *user.RpcUserCommonDeptRetrieveTreeResp, err error) {
	resp, err = service.NewRpcUserCommonDeptRetrieveTreeService(ctx).Run(req)

	return resp, err
}

// RpcUserCommonDeptUpdate implements the UserServiceImpl interface.
func (s *UserServiceImpl) RpcUserCommonDeptUpdate(ctx context.Context, req *user.RpcUserCommonDeptUpdateReq) (resp *user.RpcUserCommonDeptUpdateResp, err error) {
	resp, err = service.NewRpcUserCommonDeptUpdateService(ctx).Run(req)

	return resp, err
}

// RpcUserCommonDeptDelete implements the UserServiceImpl interface.
func (s *UserServiceImpl) RpcUserCommonDeptDelete(ctx context.Context, req *user.RpcUserCommonDeptDeleteReq) (resp *user.RpcUserCommonDeptDeleteResp, err error) {
	resp, err = service.NewRpcUserCommonDeptDeleteService(ctx).Run(req)

	return resp, err
}
