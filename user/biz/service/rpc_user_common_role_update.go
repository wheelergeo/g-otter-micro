package service

import (
	"context"

	"github.com/cloudwego/kitex/pkg/klog"
	user "github.com/wheelergeo/g-otter-gen/user"
	"github.com/wheelergeo/g-otter-micro-user/biz/dal/model"
	"github.com/wheelergeo/g-otter-micro-user/biz/dal/mysql"
	"github.com/wheelergeo/g-otter-pkg/constants"
)

type RpcUserCommonRoleUpdateService struct {
	ctx context.Context
} // NewRpcUserCommonRoleUpdateService new RpcUserCommonRoleUpdateService
func NewRpcUserCommonRoleUpdateService(ctx context.Context) *RpcUserCommonRoleUpdateService {
	return &RpcUserCommonRoleUpdateService{ctx: ctx}
}

// Run create note info
func (s *RpcUserCommonRoleUpdateService) Run(req *user.RpcUserCommonRoleUpdateReq) (resp *user.RpcUserCommonRoleUpdateResp, err error) {
	err = req.IsValid()
	if err != nil {
		klog.CtxErrorf(s.ctx, err.Error())
		resp.Status.Code = constants.ParamInvalid
		resp.Status.Message = constants.ParamInvalidMsg
		return
	}
	resp = new(user.RpcUserCommonRoleUpdateResp)
	resp.Status = new(user.BaseResp)

	roleModel := model.UserCommonRole{}
	err = mysql.DB.First(&roleModel, req.Id).Error
	if err != nil {
		klog.CtxErrorf(s.ctx, err.Error())
		resp.Status.Code = constants.UserCommonRoleUpdateFailed
		resp.Status.Message = constants.UserCommonRoleUpdateFailedMsg
		return
	}

	if req.Status != nil {
		roleModel.Status = int32(*req.Status)
	}
	if req.Level != nil {
		roleModel.Level = int32(*req.Level)
	}
	if req.Name != nil {
		roleModel.Name = *req.Name
	}
	if req.Remark != nil {
		roleModel.Remark = *req.Remark
	}
	if req.DataScope != nil {
		roleModel.DataScope = int32(*req.DataScope)
	}

	err = mysql.DB.Save(&roleModel).Error
	if err != nil {
		klog.CtxErrorf(s.ctx, err.Error())
		resp.Status.Code = constants.UserCommonRoleUpdateFailed
		resp.Status.Message = constants.UserCommonRoleUpdateFailedMsg
		return
	}

	resp.Status.Code = constants.UserCommonRoleUpdateSuccess
	resp.Status.Message = constants.UserCommonRoleUpdateSuccessMsg
	return
}
