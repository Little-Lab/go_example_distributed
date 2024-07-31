package logic

import (
	"context"
	"database/sql"
	"fmt"
	"score/internal/model"

	"github.com/dtm-labs/client/dtmgrpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"score/internal/svc"
	"score/pb/scores"

	"github.com/zeromicro/go-zero/core/logx"
)

type TranDecScoreLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewTranDecScoreLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TranDecScoreLogic {
	return &TranDecScoreLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *TranDecScoreLogic) TranDecScore(in *scores.AddScoreReq) (*scores.AddScoreResp, error) {

	fmt.Println(">>>>>>>>>>>>start score delete tran! userId is:", in.UserId)
	scoreM, _ := l.svcCtx.ScoreModel.FindByUserId(l.ctx, in.UserId)
	var scoreNum int64 = 0
	fmt.Sprintf("------------>4", scoreM)
	if scoreM != nil {
		scoreNum = scoreM.Score + scoreNum
	}

	barrier, err := dtmgrpc.BarrierFromGrpc(l.ctx)
	if err != nil {
		fmt.Println(">>>>>>>>>>>>score dec barrier error!")
	}
	if err := barrier.CallWithDB(l.svcCtx.Db, func(tx *sql.Tx) error {
		score := new(model.Score)
		score.UserId = in.UserId
		score.Score = scoreNum - in.Score
		fmt.Println("-------------->score:", score.Score)
		_, err := l.svcCtx.ScoreModel.TranUpdate(tx, score)
		if err != nil {
			return fmt.Errorf("积分更新失败 err : %v , score:%+v \n", err, score)
		}
		return nil
	}); err != nil {
		return nil, status.Error(codes.Aborted, err.Error())
	}
	return &scores.AddScoreResp{}, nil
}
