package service

import (
	"context"
	"errors"
	"time"

	user "github.com/wheelergeo/g-otter-gen/user"
	"github.com/wheelergeo/g-otter-micro-user/biz/dal/model"
	"github.com/wheelergeo/g-otter-micro-user/biz/dal/mysql"
	"github.com/wheelergeo/g-otter-pkg/constants"
	"gorm.io/gorm"
)

type RpcUpdateUserOnlineService struct {
	ctx context.Context
} // NewRpcUpdateUserOnlineService new RpcUpdateUserOnlineService
func NewRpcUpdateUserOnlineService(ctx context.Context) *RpcUpdateUserOnlineService {
	return &RpcUpdateUserOnlineService{ctx: ctx}
}

// Run create note info
func (s *RpcUpdateUserOnlineService) Run(req *user.RpcUpdateUserOnlineReq) (resp *user.RpcUpdateUserOnlineResp, err error) {
	resp = new(user.RpcUpdateUserOnlineResp)
	online := model.UserCommonOnline{}
	err = mysql.DB.Where("token", req.Token).
		Limit(1).Find(&online).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		resp.Status.Code = constants.UpdateUserOnlineFailed
		resp.Status.Message = constants.UpdateUserOnlineFailedMsg
	}
	expiredAt, _ := time.Parse("2006-01-02 15:04:05", req.TokenExpiredAt)
	loginAt, _ := time.Parse("2006-01-02 15:04:05", req.LoginAt)
	online.TokenExpiredAt = expiredAt
	online.LoginAt = loginAt

	err = mysql.DB.Save(&online).Error
	if err != nil {
		resp.Status.Code = constants.UpdateUserOnlineFailed
		resp.Status.Message = constants.UpdateUserOnlineFailedMsg
		return
	}

	resp.Status.Code = constants.UpdateUserOnlineSuccess
	resp.Status.Message = constants.UpdateUserOnlineSuccessMsg

	return
}
