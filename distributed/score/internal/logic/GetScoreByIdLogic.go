package logic

import (
	"context"

	"score/internal/svc"
	"score/pb/scores"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetScoreByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetScoreByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetScoreByIdLogic {
	return &GetScoreByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetScoreByIdLogic) GetScoreById(in *scores.GetScoreByIdReq) (*scores.GetScoreByIdResp, error) {
	// todo: add your logic here and delete this line

	return &scores.GetScoreByIdResp{}, nil
}
