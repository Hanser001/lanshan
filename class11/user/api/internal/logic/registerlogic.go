package logic

import (
	"context"
	"errors"
	"lanshan/class11/user/model"

	"lanshan/class11/user/api/internal/svc"
	"lanshan/class11/user/api/internal/types"

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

	_, err = l.svcCtx.UserModel.FindOneByUsername(l.ctx, req.Username)
	switch err {
	case nil:
		return &types.RegisterRes{
			Code: 400,
			Msg:  "user has been existed",
		}, errors.New("user has been existed")

	case model.ErrNotFound:
		newUser := model.User{
			Username: req.Username,
			Password: req.Password,
		}
		l.svcCtx.UserModel.Insert(l.ctx, &newUser)
		return &types.RegisterRes{
			Code: 0,
			Msg:  "register successfully",
		}, nil

	default:
		return &types.RegisterRes{
			Code: 400,
			Msg:  "unknown error",
		}, errors.New("unknown error")
	}
}
