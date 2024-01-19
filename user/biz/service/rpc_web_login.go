package service

import (
	"context"
	"errors"
	"time"

	user "github.com/wheelergeo/g-otter-gen/user"
	"github.com/wheelergeo/g-otter-micro-user/biz/dal/model"
	"github.com/wheelergeo/g-otter-micro-user/biz/dal/mysql"
	"github.com/wheelergeo/g-otter-micro-user/conf"
	"github.com/wheelergeo/g-otter-pkg/constants"
	"github.com/wheelergeo/g-otter-pkg/token"
	"github.com/wheelergeo/g-otter-pkg/utils"
	"gorm.io/gorm"
)

type RpcWebLoginService struct {
	ctx context.Context
} // NewRpcWebLoginService new RpcWebLoginService
func NewRpcWebLoginService(ctx context.Context) *RpcWebLoginService {
	return &RpcWebLoginService{ctx: ctx}
}

// Run create note info
func (s *RpcWebLoginService) Run(req *user.RpcWebLoginReq) (resp *user.RpcWebLoginResp, err error) {
	//klog.CtxInfof(s.ctx, "baggage:%s", baggage.FromContext(s.ctx).String())
	resp = new(user.RpcWebLoginResp)
	resp.Status = new(user.BaseResp)
	adminInfo := model.UserAdminInfo{}
	t := time.Now()
	loginLog := model.UserCommonLogin{
		PhoneNumber: req.PhoneNum,
		IP:          req.Ip,
		Location:    req.Location,
		Browser:     req.Browser,
		Os:          req.Os,
		LoginAt:     t,
		Client:      "web",
	}

	defer func() {
		err = mysql.DB.Create(&loginLog).Error
	}()

	// check user is existed
	err = mysql.DB.Where("phone_number", req.PhoneNum).
		Limit(1).Find(&adminInfo).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		resp.Status.Code = constants.UserNotExist
		resp.Status.Message = constants.UserNotExistMsg
		loginLog.Msg = constants.UserNotExistMsg
		loginLog.Status = 1
		return
	}
	loginLog.UUID = adminInfo.UUID
	loginLog.UserName = adminInfo.UserName

	// check password is right
	b := utils.VerifyPwd(adminInfo.UserPassword, req.Password)
	if !b {
		resp.Status.Code = constants.WrongPassword
		resp.Status.Message = constants.WrongPasswordMsg
		err = errors.New(constants.WrongPasswordMsg)
		loginLog.Msg = constants.WrongPasswordMsg
		loginLog.Status = 1
		return
	}

	// generate token
	auth, expiredAt, err := token.TokenAuth().New(
		conf.GetConf().Paseto.AdminRefresh,
		conf.GetConf().Paseto.AdminTimeout,
		token.ClaimData{},
	)
	if err != nil {
		resp.Status.Code = constants.GenTokenErr
		resp.Status.Message = constants.GenTokenErrMsg
		loginLog.Status = 1
		loginLog.Msg = constants.GenTokenErrMsg
		return
	}

	loginLog.Status = 0
	loginLog.Msg = constants.LoginSuccessMsg
	resp.Token = auth
	resp.Status.Code = constants.LoginSuccess
	resp.Status.Message = constants.LoginSuccessMsg

	// update user online
	newOnline := model.UserCommonOnline{
		LoginAt:        t,
		UserName:       adminInfo.UserName,
		PhoneNumber:    adminInfo.PhoneNumber,
		Token:          auth,
		TokenExpiredAt: expiredAt,
		IP:             req.Ip,
		Explorer:       req.Browser,
		Os:             req.Os,
		Client:         "web",
	}

	// multi-point login of the same client is not support
	oldOnline := model.UserCommonOnline{}
	err = mysql.DB.Where("client=web AND phone_number=?", req.PhoneNum).
		Limit(1).Find(&oldOnline).Error
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		newOnline.ID = oldOnline.ID
		token.TokenAuth().Delete(oldOnline.Token)
	}
	err = mysql.DB.Save(&newOnline).Error

	return
}
