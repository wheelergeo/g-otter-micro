package service

import (
	"context"

	"github.com/cloudwego/kitex/pkg/klog"
	user "github.com/wheelergeo/g-otter-gen/user"
	"github.com/wheelergeo/g-otter-micro-user/biz/dal/model"
	"github.com/wheelergeo/g-otter-micro-user/biz/dal/mysql"
	"github.com/wheelergeo/g-otter-pkg/constants"
)

type RpcUserCommonRoleRetrieveService struct {
	ctx context.Context
} // NewRpcUserCommonRoleRetrieveService new RpcUserCommonRoleRetrieveService
func NewRpcUserCommonRoleRetrieveService(ctx context.Context) *RpcUserCommonRoleRetrieveService {
	return &RpcUserCommonRoleRetrieveService{ctx: ctx}
}

// Run create note info
func (s *RpcUserCommonRoleRetrieveService) Run(req *user.RpcUserCommonRoleRetrieveReq) (resp *user.RpcUserCommonRoleRetrieveResp, err error) {
	err = req.IsValid()
	if err != nil {
		klog.CtxErrorf(s.ctx, err.Error())
		resp.Status.Code = constants.ParamInvalid
		resp.Status.Message = constants.ParamInvalidMsg
		return
	}
	resp = new(user.RpcUserCommonRoleRetrieveResp)
	resp.Status = new(user.BaseResp)

	roleModels := []model.UserCommonRole{}
	db := mysql.DB
	if req.Status != nil {
		db = db.Where("status = ?", *req.Status)
	}
	if req.Level != nil {
		db = db.Where("level = ?", *req.Level)
	}
	if req.Name != nil {
		db = db.Where("name LIKE ?", "%"+*req.Name+"%")
	}
	if req.Id != nil {
		db = db.Where("id = ?", *req.Id)
	}

	err = mysql.DB.Find(&roleModels).Error
	if err != nil {
		klog.CtxErrorf(s.ctx, err.Error())
		resp.Status.Code = constants.UserCommonRoleRetrieveFailed
		resp.Status.Message = constants.UserCommonRoleRetrieveFailedMsg
		return
	}

	for _, v := range roleModels {
		resp.List = append(resp.List, &user.UserRoleInfo{
			Id:        int32(v.ID),
			Status:    int8(v.Status),
			Level:     int8(v.Level),
			Name:      v.Name,
			Remark:    v.Remark,
			DataScope: int8(v.DataScope),
			CreateAt:  v.CreatedAt.String(),
			UpdateAt:  v.UpdatedAt.String(),
		})
	}

	resp.Status.Code = constants.UserCommonRoleRetrieveSuccess
	resp.Status.Message = constants.UserCommonRoleRetrieveSuccessMsg
	return
}
