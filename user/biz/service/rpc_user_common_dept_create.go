package service

import (
	"context"

	"github.com/cloudwego/kitex/pkg/klog"
	user "github.com/wheelergeo/g-otter-gen/user"
	"github.com/wheelergeo/g-otter-micro-user/biz/dal/model"
	"github.com/wheelergeo/g-otter-micro-user/biz/dal/mysql"
	"github.com/wheelergeo/g-otter-pkg/constants"
)

type RpcUserCommonDeptCreateService struct {
	ctx context.Context
} // NewRpcUserCommonDeptCreateService new RpcUserCommonDeptCreateService
func NewRpcUserCommonDeptCreateService(ctx context.Context) *RpcUserCommonDeptCreateService {
	return &RpcUserCommonDeptCreateService{ctx: ctx}
}

// Run create note info
func (s *RpcUserCommonDeptCreateService) Run(req *user.RpcUserCommonDeptCreateReq) (resp *user.RpcUserCommonDeptCreateResp, err error) {
	// Finish your business logic.
	resp = new(user.RpcUserCommonDeptCreateResp)
	resp.Status = new(user.BaseResp)
	deptModel := model.UserCommonDept{
		ParentID:   int64(req.ParentId),
		AncestorID: int64(req.AncestorId),
		DeptName:   req.DeptName,
		Status:     int32(req.Status),
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

	err = mysql.DB.Create(&deptModel).Error
	if err != nil {
		klog.CtxErrorf(s.ctx, err.Error())
		resp.Status.Code = constants.UserCommonDeptCreateFailed
		resp.Status.Message = constants.UserCommonDeptCreateFailedMsg
		return
	}

	resp.Status.Code = constants.UserCommonDeptCreateSuccess
	resp.Status.Message = constants.UserCommonDeptCreateSuccessMsg
	return
}
