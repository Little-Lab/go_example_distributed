package logic

import (
	"context"

	"score/internal/svc"
	"score/pb/scores"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateScoreLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateScoreLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateScoreLogic {
	return &UpdateScoreLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateScoreLogic) UpdateScore(in *scores.UpdateScoreReq) (*scores.UpdateScoreResp, error) {
	// todo: add your logic here and delete this line

	return &scores.UpdateScoreResp{}, nil
}
