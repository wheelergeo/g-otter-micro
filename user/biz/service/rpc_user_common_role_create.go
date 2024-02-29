package service

import (
	"context"

	"github.com/cloudwego/kitex/pkg/klog"
	user "github.com/wheelergeo/g-otter-gen/user"
	"github.com/wheelergeo/g-otter-micro-user/biz/dal/model"
	"github.com/wheelergeo/g-otter-micro-user/biz/dal/mysql"
	"github.com/wheelergeo/g-otter-pkg/constants"
)

type RpcUserCommonRoleCreateService struct {
	ctx context.Context
} // NewRpcUserCommonRoleCreateService new RpcUserCommonRoleCreateService
func NewRpcUserCommonRoleCreateService(ctx context.Context) *RpcUserCommonRoleCreateService {
	return &RpcUserCommonRoleCreateService{ctx: ctx}
}

// Run create note info
func (s *RpcUserCommonRoleCreateService) Run(req *user.RpcUserCommonRoleCreateReq) (resp *user.RpcUserCommonRoleCreateResp, err error) {
	err = req.IsValid()
	if err != nil {
		klog.CtxErrorf(s.ctx, err.Error())
		resp.Status.Code = constants.ParamInvalid
		resp.Status.Message = constants.ParamInvalidMsg
		return
	}

	resp = new(user.RpcUserCommonRoleCreateResp)
	resp.Status = new(user.BaseResp)
	roleModel := model.UserCommonRole{
		Status:    int32(req.Status),
		Level:     int32(req.Level),
		Name:      req.Name,
		DataScope: int32(req.DataScope),
	}
	if req.Remark != nil {
		roleModel.Remark = *req.Remark
	}

	err = mysql.DB.Create(&roleModel).Error
	if err != nil {
		klog.CtxErrorf(s.ctx, err.Error())
		resp.Status.Code = constants.UserCommonRoleCreateFailed
		resp.Status.Message = constants.UserCommonRoleCreateFailedMsg
		return
	}

	resp.Status.Code = constants.UserCommonRoleCreateSuccess
	resp.Status.Message = constants.UserCommonRoleCreateSuccessMsg
	return
}
