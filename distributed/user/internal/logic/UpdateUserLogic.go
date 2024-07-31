package logic

import (
	"context"

	"user/internal/svc"
	"user/pb/users"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateUserLogic {
	return &UpdateUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateUserLogic) UpdateUser(in *users.UpdateUserReq) (*users.UpdateUserResp, error) {
	// todo: add your logic here and delete this line

	return &users.UpdateUserResp{}, nil
}
