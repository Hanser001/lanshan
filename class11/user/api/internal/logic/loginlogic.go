package logic

import (
	"context"
	"errors"
	"lanshan/class11/user/rpc/pb/user"
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

	_, err = l.svcCtx.SysRpcClient.Login(l.ctx, &user.LoginReq{
		Username: req.Username,
		Password: req.Password,
	})

	if err != nil {
		return nil, err
	}

	return &types.LoginRes{
		Code: 0,
		Msg:  "ok",
	}, nil
}
