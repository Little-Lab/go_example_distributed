package logic

import (
	"context"
	"fmt"
	"score/pb/scores"
	"user/pb/users"

	"github.com/dtm-labs/client/dtmgrpc"

	"bff/internal/svc"
	"bff/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type InviteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewInviteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *InviteLogic {
	return &InviteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// var dtmServer = "etcd://127.0.0.1:2379/dtmservice"
var dtmServer = "etcd://localhost:2379/dtmservice"

func (l *InviteLogic) Invite(req *types.InviteReq) (resp *types.InviteResp, err error) {
	//dtm 事务id
	gid := dtmgrpc.MustGenGid(dtmServer)
	fmt.Println("------------------->dtm:", gid)
	//选择一种分布式事务的实现方式: saga
	//实例化grpc协议方式的saga句柄
	sagaGrpc := dtmgrpc.NewSagaGrpc(dtmServer, gid)
	//用户微服
	u, _ := l.svcCtx.UserSrv.GetUserByCode(l.ctx, &users.GetUserByCodeReq{
		Code: req.Code,
	})
	fmt.Println("----->u:", u)
	userSrv, _ := l.svcCtx.Config.UserSrv.BuildTarget()
	sagaGrpc.Add(userSrv+"/users.user/TranAddUser", userSrv+"/users.user/TranDelUser", &users.InviteUserReq{
		Email: req.Email,
	})
	//积分微服
	scoreSrv, _ := l.svcCtx.Config.ScoreSrv.BuildTarget()
	sagaGrpc.Add(scoreSrv+"/scores.score/TranIncScore", scoreSrv+"/scores.score/TranDecScore", &scores.AddScoreReq{
		UserId: u.User.Id,
		Score:  1,
	})

	//提交事务
	err = sagaGrpc.Submit()
	if err != nil {
		return nil, fmt.Errorf("submit data to  dtm-server err  : %+v \n", err)
	}
	return
}
