package logic

import (
	"context"
	"fmt"
	"score/internal/model"
	"score/internal/svc"
	"score/pb/scores"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddScoreLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddScoreLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddScoreLogic {
	return &AddScoreLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// -----------------------score-----------------------
func (l *AddScoreLogic) AddScore(in *scores.AddScoreReq) (*scores.AddScoreResp, error) {
	score := new(model.Score)
	scoreM, _ := l.svcCtx.ScoreModel.FindByUserId(l.ctx, in.UserId)
	if scoreM == nil {
		score.Score = 0
		score.UserId = in.UserId
		score.Score = in.Score
		l.svcCtx.ScoreModel.Insert(l.ctx, score)
	} else {
		score.UserId = in.UserId
		score.Score = scoreM.Score + in.Score
		err := l.svcCtx.ScoreModel.Update(l.ctx, score)
		if err != nil {
			return nil, fmt.Errorf("积分更新失败 err : %v , score:%+v \n", err, score)
		}
	}
	return &scores.AddScoreResp{}, nil
}
