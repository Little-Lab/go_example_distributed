package logic

import (
	"context"

	"user/internal/svc"
	"user/pb/users"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDelUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelUserLogic {
	return &DelUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DelUserLogic) DelUser(in *users.DelUserReq) (*users.DelUserResp, error) {
	// todo: add your logic here and delete this line

	return &users.DelUserResp{}, nil
}
