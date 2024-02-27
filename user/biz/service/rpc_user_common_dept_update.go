package service

import (
	"context"

	"github.com/cloudwego/kitex/pkg/klog"
	user "github.com/wheelergeo/g-otter-gen/user"
	"github.com/wheelergeo/g-otter-micro-user/biz/dal/model"
	"github.com/wheelergeo/g-otter-micro-user/biz/dal/mysql"
	"github.com/wheelergeo/g-otter-pkg/constants"
)

type RpcUserCommonDeptUpdateService struct {
	ctx context.Context
} // NewRpcUserCommonDeptUpdateService new RpcUserCommonDeptUpdateService
func NewRpcUserCommonDeptUpdateService(ctx context.Context) *RpcUserCommonDeptUpdateService {
	return &RpcUserCommonDeptUpdateService{ctx: ctx}
}

// Run create note info
func (s *RpcUserCommonDeptUpdateService) Run(req *user.RpcUserCommonDeptUpdateReq) (resp *user.RpcUserCommonDeptUpdateResp, err error) {
	err = req.IsValid()
	if err != nil {
		klog.CtxErrorf(s.ctx, err.Error())
		resp.Status.Code = constants.ParamInvalid
		resp.Status.Message = constants.ParamInvalidMsg
		return
	}
	resp = new(user.RpcUserCommonDeptUpdateResp)
	resp.Status = new(user.BaseResp)

	deptModel := model.UserCommonDept{}
	err = mysql.DB.First(&deptModel, req.Id).Error
	if err != nil {
		klog.CtxErrorf(s.ctx, err.Error())
		resp.Status.Code = constants.UserCommonDeptUpdateFailed
		resp.Status.Message = constants.UserCommonDeptUpdateFailedMsg
		return
	}

	if req.ParentId != nil {
		deptModel.ParentID = int64(*req.ParentId)
	}
	if req.AncestorId != nil {
		deptModel.AncestorID = int64(*req.AncestorId)
	}
	if req.DeptName != nil {
		deptModel.DeptName = *req.DeptName
	}
	if req.Leader != nil {
		deptModel.Leader = *req.Leader
	}
	if req.PhoneNum != nil {
		deptModel.PhoneNumber = *req.PhoneNum
	}
	if req.Email != nil {
		deptModel.Email = *req.Email
	}
	if req.Status != nil {
		deptModel.Status = int32(*req.Status)
	}

	err = mysql.DB.Save(&deptModel).Error
	if err != nil {
		klog.CtxErrorf(s.ctx, err.Error())
		resp.Status.Code = constants.UserCommonDeptUpdateFailed
		resp.Status.Message = constants.UserCommonDeptUpdateFailedMsg
		return
	}

	resp.Status.Code = constants.UserCommonDeptUpdateSuccess
	resp.Status.Message = constants.UserCommonDeptUpdateSuccessMsg
	return
}
