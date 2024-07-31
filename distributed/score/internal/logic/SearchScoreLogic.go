package logic

import (
	"context"

	"score/internal/svc"
	"score/pb/scores"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchScoreLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSearchScoreLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchScoreLogic {
	return &SearchScoreLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SearchScoreLogic) SearchScore(in *scores.SearchScoreReq) (*scores.SearchScoreResp, error) {
	// todo: add your logic here and delete this line

	return &scores.SearchScoreResp{}, nil
}
