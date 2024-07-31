package logic

import (
	"context"
	"database/sql"
	"fmt"
	"user/internal/model"

	"github.com/dtm-labs/client/dtmgrpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"user/internal/svc"
	"user/pb/users"

	"github.com/zeromicro/go-zero/core/logx"
)

type TranDelUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewTranDelUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TranDelUserLogic {
	return &TranDelUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *TranDelUserLogic) TranDelUser(in *users.InviteUserReq) (*users.AddUserResp, error) {
	fmt.Println(">>>>>>>>>>>>TranDelUser")

	barrier, err := dtmgrpc.BarrierFromGrpc(l.ctx)
	if err != nil {
		fmt.Println(">>>>>>>>>>>>score del barrier error!")
	}
	if err := barrier.CallWithDB(l.svcCtx.Db, func(tx *sql.Tx) error {
		fmt.Println(">>>>>>>>>>>>start user del tran!")
		user := new(model.User)
		user.Email.String = in.Email
		l.svcCtx.UserModel.TranDelete(tx, user)
		return nil
	}); err != nil {
		return nil, status.Error(codes.Aborted, err.Error())
	}
	return &users.AddUserResp{}, nil
}
