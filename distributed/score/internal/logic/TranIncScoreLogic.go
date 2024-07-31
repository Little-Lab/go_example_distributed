package logic

import (
	"context"
	"database/sql"
	"fmt"
	"score/internal/model"
	"score/internal/svc"
	"score/pb/scores"

	"github.com/dtm-labs/client/dtmgrpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/zeromicro/go-zero/core/logx"
)

type TranIncScoreLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewTranIncScoreLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TranIncScoreLogic {
	return &TranIncScoreLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *TranIncScoreLogic) TranIncScore(in *scores.AddScoreReq) (*scores.AddScoreResp, error) {
	fmt.Println(">>>>>>>>>>>>score inc log")
	barrier, err := dtmgrpc.BarrierFromGrpc(l.ctx)
	if err != nil {
		fmt.Println(">>>>>>>>>>>>score inc barrier error!")
	}
	if err := barrier.CallWithDB(l.svcCtx.Db, func(tx *sql.Tx) error {
		fmt.Println(">>>>>>>>>>>>start score inc tran!")
		score := new(model.Score)
		scoreM, _ := l.svcCtx.ScoreModel.FindByUserId(l.ctx, in.UserId)
		if scoreM == nil {
			score.Score = 0
			score.UserId = in.UserId
			score.Score = in.Score
			l.svcCtx.ScoreModel.TranInsert(tx, score)
		} else {
			score.UserId = in.UserId
			score.Score = scoreM.Score + in.Score
			_, err := l.svcCtx.ScoreModel.TranUpdate(tx, score)
			if err != nil {
				return fmt.Errorf("积分更新失败 err : %v , score:%+v \n", err, score)
			}
		}
		return nil
	}); err != nil {
		return nil, status.Error(codes.Aborted, err.Error())
	}
	return &scores.AddScoreResp{}, nil
}
