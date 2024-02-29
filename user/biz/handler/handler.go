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

// RpcUserCommonRoleCreate implements the UserServiceImpl interface.
func (s *UserServiceImpl) RpcUserCommonRoleCreate(ctx context.Context, req *user.RpcUserCommonRoleCreateReq) (resp *user.RpcUserCommonRoleCreateResp, err error) {
	resp, err = service.NewRpcUserCommonRoleCreateService(ctx).Run(req)

	return resp, err
}

// RpcUserCommonRoleRetrieve implements the UserServiceImpl interface.
func (s *UserServiceImpl) RpcUserCommonRoleRetrieve(ctx context.Context, req *user.RpcUserCommonRoleRetrieveReq) (resp *user.RpcUserCommonRoleRetrieveResp, err error) {
	resp, err = service.NewRpcUserCommonRoleRetrieveService(ctx).Run(req)

	return resp, err
}

// RpcUserCommonRoleUpdate implements the UserServiceImpl interface.
func (s *UserServiceImpl) RpcUserCommonRoleUpdate(ctx context.Context, req *user.RpcUserCommonRoleUpdateReq) (resp *user.RpcUserCommonRoleUpdateResp, err error) {
	resp, err = service.NewRpcUserCommonRoleUpdateService(ctx).Run(req)

	return resp, err
}

// RpcUserCommonRoleDelete implements the UserServiceImpl interface.
func (s *UserServiceImpl) RpcUserCommonRoleDelete(ctx context.Context, req *user.RpcUserCommonRoleDeleteReq) (resp *user.RpcUserCommonRoleDeleteResp, err error) {
	resp, err = service.NewRpcUserCommonRoleDeleteService(ctx).Run(req)

	return resp, err
}

// RpcUserCommonPostCreate implements the UserServiceImpl interface.
func (s *UserServiceImpl) RpcUserCommonPostCreate(ctx context.Context, req *user.RpcUserCommonPostCreateReq) (resp *user.RpcUserCommonPostCreateResp, err error) {
	resp, err = service.NewRpcUserCommonPostCreateService(ctx).Run(req)

	return resp, err
}

// RpcUserCommonPostRetrieve implements the UserServiceImpl interface.
func (s *UserServiceImpl) RpcUserCommonPostRetrieve(ctx context.Context, req *user.RpcUserCommonPostRetrieveReq) (resp *user.RpcUserCommonPostRetrieveResp, err error) {
	resp, err = service.NewRpcUserCommonPostRetrieveService(ctx).Run(req)

	return resp, err
}

// RpcUserCommonPostUpdate implements the UserServiceImpl interface.
func (s *UserServiceImpl) RpcUserCommonPostUpdate(ctx context.Context, req *user.RpcUserCommonPostUpdateReq) (resp *user.RpcUserCommonPostUpdateResp, err error) {
	resp, err = service.NewRpcUserCommonPostUpdateService(ctx).Run(req)

	return resp, err
}

// RpcUserCommonPostDelete implements the UserServiceImpl interface.
func (s *UserServiceImpl) RpcUserCommonPostDelete(ctx context.Context, req *user.RpcUserCommonPostDeleteReq) (resp *user.RpcUserCommonPostDeleteResp, err error) {
	resp, err = service.NewRpcUserCommonPostDeleteService(ctx).Run(req)

	return resp, err
}

// RpcUserCommonKvCreate implements the UserServiceImpl interface.
func (s *UserServiceImpl) RpcUserCommonKvCreate(ctx context.Context, req *user.RpcUserCommonKvCreateReq) (resp *user.RpcUserCommonKvCreateResp, err error) {
	resp, err = service.NewRpcUserCommonKvCreateService(ctx).Run(req)

	return resp, err
}

// RpcUserCommonKvRetrieve implements the UserServiceImpl interface.
func (s *UserServiceImpl) RpcUserCommonKvRetrieve(ctx context.Context, req *user.RpcUserCommonKvRetrieveReq) (resp *user.RpcUserCommonKvRetrieveResp, err error) {
	resp, err = service.NewRpcUserCommonKvRetrieveService(ctx).Run(req)

	return resp, err
}

// RpcUserCommonKvUpdate implements the UserServiceImpl interface.
func (s *UserServiceImpl) RpcUserCommonKvUpdate(ctx context.Context, req *user.RpcUserCommonKvUpdateReq) (resp *user.RpcUserCommonKvUpdateResp, err error) {
	resp, err = service.NewRpcUserCommonKvUpdateService(ctx).Run(req)

	return resp, err
}

// RpcUserCommonKvDelete implements the UserServiceImpl interface.
func (s *UserServiceImpl) RpcUserCommonKvDelete(ctx context.Context, req *user.RpcUserCommonKvDeleteReq) (resp *user.RpcUserCommonKvDeleteResp, err error) {
	resp, err = service.NewRpcUserCommonKvDeleteService(ctx).Run(req)

	return resp, err
}
