package logic

import (
	"context"
	"errors"
	"lanshan/class11/user/model"

	"lanshan/class11/user/rpc/internal/svc"
	"lanshan/class11/user/rpc/pb/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RegisterLogic) Register(in *user.RegisterReq) (*user.RegisterRes, error) {

	_, err := l.svcCtx.UserModel.FindOneByUsername(l.ctx, in.Username)

	switch err {
	case nil:
		return &user.RegisterRes{
			Code: 400,
			Msg:  "user has benn existed",
		}, errors.New("user has been existed")

	case model.ErrNotFound:
		return &user.RegisterRes{
			Code: 0,
			Msg:  "register successfully",
		}, nil
	default:

		return &user.RegisterRes{
			Code: 400,
			Msg:  "unknown err",
		}, errors.New("unknown err")
	}
}
