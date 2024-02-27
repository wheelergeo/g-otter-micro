package service

import (
	"context"

	"github.com/cloudwego/kitex/pkg/klog"
	user "github.com/wheelergeo/g-otter-gen/user"
	"github.com/wheelergeo/g-otter-micro-user/biz/dal/model"
	"github.com/wheelergeo/g-otter-micro-user/biz/dal/mysql"
	"github.com/wheelergeo/g-otter-pkg/constants"
)

type RpcUserCommonDeptRetrieveTreeService struct {
	ctx context.Context
} // NewRpcUserCommonDeptRetrieveTreeService new RpcUserCommonDeptRetrieveTreeService
func NewRpcUserCommonDeptRetrieveTreeService(ctx context.Context) *RpcUserCommonDeptRetrieveTreeService {
	return &RpcUserCommonDeptRetrieveTreeService{ctx: ctx}
}

// Run create note info
func (s *RpcUserCommonDeptRetrieveTreeService) Run(req *user.RpcUserCommonDeptRetrieveTreeReq) (resp *user.RpcUserCommonDeptRetrieveTreeResp, err error) {
	err = req.IsValid()
	if err != nil {
		klog.CtxErrorf(s.ctx, err.Error())
		resp.Status.Code = constants.ParamInvalid
		resp.Status.Message = constants.ParamInvalidMsg
		return
	}
	resp = new(user.RpcUserCommonDeptRetrieveTreeResp)
	resp.Status = new(user.BaseResp)

	deptModels := []model.UserCommonDept{}
	db := mysql.DB.Where("parent_id = ?", 0)
	if req.AncestorId != nil {
		db = db.Where("id = ?", *req.AncestorId)
	}

	err = db.Find(&deptModels).Error
	if err != nil {
		klog.CtxErrorf(s.ctx, err.Error())
		resp.Status.Code = constants.UserCommonDeptRetrieveTreeFailed
		resp.Status.Message = constants.UserCommonDeptRetrieveTreeFailedMsg
		return
	}

	for _, v := range deptModels {
		tree, _ := deptTree(v.ID)
		resp.List = append(resp.List, &tree)
	}

	resp.Status.Code = constants.UserCommonDeptRetrieveTreeSuccess
	resp.Status.Message = constants.UserCommonDeptRetrieveTreeSuccessMsg
	return
}

func deptTree(id int64) (t user.UserDeptTree, err error) {
	deptModel := model.UserCommonDept{}
	err = mysql.DB.Where("id = ?", id).First(&deptModel).Error
	t.Parent = &user.UserDeptInfo{
		Id:         int32(deptModel.ID),
		ParentId:   int32(deptModel.ParentID),
		AncestorId: int32(deptModel.AncestorID),
		DeptName:   deptModel.DeptName,
		Leader:     deptModel.Leader,
		PhoneNum:   deptModel.PhoneNumber,
		Email:      deptModel.Email,
		Status:     int8(deptModel.Status),
		CreateAt:   deptModel.CreatedAt.String(),
		UpdateAt:   deptModel.UpdatedAt.String(),
	}
	if err != nil {
		klog.Errorf(err.Error())
		return
	}

	deptModels := []model.UserCommonDept{}
	err = mysql.DB.Where("parent_id = ?", id).Find(&deptModels).Error
	if err != nil {
		klog.Errorf(err.Error())
		return
	}

	if len(deptModels) == 0 {
		return t, nil
	}

	for _, v := range deptModels {
		t2, err := deptTree(v.ID)
		if err != nil {
			continue
		}
		t.Children = append(t.Children, &t2)
	}
	return t, nil
}
