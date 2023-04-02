package logic

import (
	"context"
	"errors"
	"lanshan/class11/user/api/internal/svc"
	"lanshan/class11/user/api/internal/types"
	"lanshan/class11/user/rpc/pb/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegisterLogic) Register(req *types.RegisterReq) (resp *types.RegisterRes, err error) {
	if req.Username == "" || req.Password == "" {
		return &types.RegisterRes{
			Code: 400,
			Msg:  "",
		}, errors.New("username or password can not be null")
	}

	_, err = l.svcCtx.SysRpcClient.Register(l.ctx, &user.RegisterReq{
		Username: req.Username,
		Password: req.Password,
	})

	if err != nil {
		return nil, err
	}

	return &types.RegisterRes{
		Code: 0,
		Msg:  "ok",
	}, nil
}
