package service

import (
	"context"

	"github.com/cloudwego/kitex/pkg/klog"
	user "github.com/wheelergeo/g-otter-gen/user"
	"github.com/wheelergeo/g-otter-micro-user/biz/dal/model"
	"github.com/wheelergeo/g-otter-micro-user/biz/dal/mysql"
	"github.com/wheelergeo/g-otter-pkg/constants"
)

type RpcUserCommonDeptDeleteService struct {
	ctx context.Context
} // NewRpcUserCommonDeptDeleteService new RpcUserCommonDeptDeleteService
func NewRpcUserCommonDeptDeleteService(ctx context.Context) *RpcUserCommonDeptDeleteService {
	return &RpcUserCommonDeptDeleteService{ctx: ctx}
}

// Run create note info
func (s *RpcUserCommonDeptDeleteService) Run(req *user.RpcUserCommonDeptDeleteReq) (resp *user.RpcUserCommonDeptDeleteResp, err error) {
	// Finish your business logic.
	resp = new(user.RpcUserCommonDeptDeleteResp)
	resp.Status = new(user.BaseResp)

	err = mysql.DB.Delete(&model.UserCommonDept{}, req.Ids).Error
	if err != nil {
		klog.CtxErrorf(s.ctx, err.Error())
		resp.Status.Code = constants.UserCommonDeptDeleteFailed
		resp.Status.Message = constants.UserCommonDeptDeleteFailedMsg
		return
	}

	resp.Status.Code = constants.UserCommonDeptDeleteSuccess
	resp.Status.Message = constants.UserCommonDeptDeleteSuccessMsg
	return
}
