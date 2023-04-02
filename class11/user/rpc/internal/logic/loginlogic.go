package logic

import (
	"context"
	"errors"
	"lanshan/class11/user/rpc/internal/svc"
	"lanshan/class11/user/rpc/pb/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LoginLogic) Login(in *user.LoginReq) (*user.LoginRes, error) {
	userInfo, err := l.svcCtx.UserModel.FindOneByUsername(l.ctx, in.Username)
	if err != nil {
		return nil, err
	}

	if userInfo.Password != in.Password {
		return nil, errors.New("用户密码不正确")
	}

	return &user.LoginRes{
		Code: 0,
		Msg:  "login successfully",
	}, nil
}
