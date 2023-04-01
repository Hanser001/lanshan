package logic

import (
	"context"
	"errors"
	"lanshan/class11/user/model"
	"strings"

	"lanshan/class11/user/api/internal/svc"
	"lanshan/class11/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginReq) (resp *types.LoginRes, err error) {
	if len(strings.TrimSpace(req.Username)) == 0 || len(strings.TrimSpace(req.Password)) == 0 {
		return nil, errors.New("username or password can not be null")
	}

	userInfo, err := l.svcCtx.UserModel.FindOneByUsername(l.ctx, req.Username)
	switch err {
	case nil:
	case model.ErrNotFound:
		return nil, errors.New("用户名不存在")
	default:
		return nil, err
	}

	if userInfo.Password != req.Password {
		return nil, errors.New("用户密码不正确")
	}

	return &types.LoginRes{
		Code: 0,
		Msg:  "login successful",
	}, nil
}
