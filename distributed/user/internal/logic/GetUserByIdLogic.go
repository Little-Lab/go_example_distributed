package logic

import (
	"context"

	"user/internal/svc"
	"user/pb/users"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserByIdLogic {
	return &GetUserByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserByIdLogic) GetUserById(in *users.GetUserByIdReq) (*users.GetUserByIdResp, error) {
	// todo: add your logic here and delete this line

	return &users.GetUserByIdResp{}, nil
}
