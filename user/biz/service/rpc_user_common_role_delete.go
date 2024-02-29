package service

import (
	"context"

	"github.com/cloudwego/kitex/pkg/klog"
	user "github.com/wheelergeo/g-otter-gen/user"
	"github.com/wheelergeo/g-otter-micro-user/biz/dal/model"
	"github.com/wheelergeo/g-otter-micro-user/biz/dal/mysql"
	"github.com/wheelergeo/g-otter-pkg/constants"
)

type RpcUserCommonRoleDeleteService struct {
	ctx context.Context
} // NewRpcUserCommonRoleDeleteService new RpcUserCommonRoleDeleteService
func NewRpcUserCommonRoleDeleteService(ctx context.Context) *RpcUserCommonRoleDeleteService {
	return &RpcUserCommonRoleDeleteService{ctx: ctx}
}

// Run create note info
func (s *RpcUserCommonRoleDeleteService) Run(req *user.RpcUserCommonRoleDeleteReq) (resp *user.RpcUserCommonRoleDeleteResp, err error) {
	err = req.IsValid()
	if err != nil {
		klog.CtxErrorf(s.ctx, err.Error())
		resp.Status.Code = constants.ParamInvalid
		resp.Status.Message = constants.ParamInvalidMsg
		return
	}
	resp = new(user.RpcUserCommonRoleDeleteResp)
	resp.Status = new(user.BaseResp)

	err = mysql.DB.Delete(&model.UserCommonRole{}, req.Ids).Error
	if err != nil {
		klog.CtxErrorf(s.ctx, err.Error())
		resp.Status.Code = constants.UserCommonRoleDeleteFailed
		resp.Status.Message = constants.UserCommonRoleDeleteFailedMsg
		return
	}

	resp.Status.Code = constants.UserCommonRoleDeleteSuccess
	resp.Status.Message = constants.UserCommonRoleDeleteSuccessMsg
	return
}
