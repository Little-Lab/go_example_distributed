package logic

import (
	"context"

	"score/internal/svc"
	"score/pb/scores"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelScoreLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDelScoreLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelScoreLogic {
	return &DelScoreLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DelScoreLogic) DelScore(in *scores.DelScoreReq) (*scores.DelScoreResp, error) {
	// todo: add your logic here and delete this line

	return &scores.DelScoreResp{}, nil
}
