package service

import (
	"context"

	"github.com/cloudwego/kitex/pkg/klog"
	user "github.com/wheelergeo/g-otter-gen/user"
	"github.com/wheelergeo/g-otter-micro-user/biz/dal/model"
	"github.com/wheelergeo/g-otter-micro-user/biz/dal/mysql"
	"github.com/wheelergeo/g-otter-pkg/constants"
)

type RpcUserCommonDeptRetrieveService struct {
	ctx context.Context
} // NewRpcUserCommonDeptRetrieveService new RpcUserCommonDeptRetrieveService
func NewRpcUserCommonDeptRetrieveService(ctx context.Context) *RpcUserCommonDeptRetrieveService {
	return &RpcUserCommonDeptRetrieveService{ctx: ctx}
}

// Run create note info
func (s *RpcUserCommonDeptRetrieveService) Run(req *user.RpcUserCommonDeptRetrieveReq) (resp *user.RpcUserCommonDeptRetrieveResp, err error) {
	err = req.IsValid()
	if err != nil {
		klog.CtxErrorf(s.ctx, err.Error())
		resp.Status.Code = constants.ParamInvalid
		resp.Status.Message = constants.ParamInvalidMsg
		return
	}
	resp = new(user.RpcUserCommonDeptRetrieveResp)
	resp.Status = new(user.BaseResp)

	deptModels := []model.UserCommonDept{}
	db := mysql.DB
	if req.ParentId != nil {
		db = db.Where("parent_id = ?", *req.ParentId)
	}
	if req.AncestorId != nil {
		db = db.Where("ancestor_id = ?", *req.AncestorId)
	}
	if req.DeptName != nil {
		db = db.Where("dept_name LIKE ?", "%"+*req.DeptName+"%")
	}
	if req.Status != nil {
		db = db.Where("status = ?", *req.Status)
	}
	if req.Id != nil {
		db = db.Where("id = ?", *req.Id)
	}

	err = mysql.DB.Find(&deptModels).Error
	if err != nil {
		klog.CtxErrorf(s.ctx, err.Error())
		resp.Status.Code = constants.UserCommonDeptRetrieveFailed
		resp.Status.Message = constants.UserCommonDeptRetrieveFailedMsg
		return
	}

	for _, v := range deptModels {
		resp.List = append(resp.List, &user.UserDeptInfo{
			Id:         int32(v.ID),
			ParentId:   int32(v.ParentID),
			AncestorId: int32(v.AncestorID),
			DeptName:   v.DeptName,
			Leader:     v.Leader,
			PhoneNum:   v.PhoneNumber,
			Email:      v.Email,
			Status:     int8(v.Status),
			CreateAt:   v.CreatedAt.String(),
			UpdateAt:   v.UpdatedAt.String(),
		})
	}

	resp.Status.Code = constants.UserCommonDeptRetrieveSuccess
	resp.Status.Message = constants.UserCommonDeptRetrieveSuccessMsg
	return
}
