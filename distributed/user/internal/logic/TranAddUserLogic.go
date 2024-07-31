package logic

import (
	"context"
	"database/sql"
	"fmt"
	"time"
	"user/internal/model"
	"user/internal/svc"
	"user/pb/users"

	"github.com/bxcodec/faker"

	"github.com/dtm-labs/client/dtmgrpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/zeromicro/go-zero/core/logx"
)

type TranAddUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewTranAddUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TranAddUserLogic {
	return &TranAddUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *TranAddUserLogic) TranAddUser(in *users.InviteUserReq) (*users.AddUserResp, error) {
	//fmt.Println(">>>>>>>>>>>>TranAddUser")
	//var user model.User
	//faker.FakeData(user)
	//user.Email.String = in.Email
	//user.VaildPeroid = time.Now()
	//fmt.Println("=============>", &user)
	//l.svcCtx.UserModel.Insert(l.ctx, &user)
	barrier, err := dtmgrpc.BarrierFromGrpc(l.ctx)
	if err != nil {
		fmt.Println(">>>>>>>>>>>>user inc barrier error!")
	}
	if err := barrier.CallWithDB(l.svcCtx.Db, func(tx *sql.Tx) error {
		fmt.Println(">>>>>>>>>>>>start user inc tran!")
		var user model.User
		faker.FakeData(user)
		user.Email.String = in.Email
		user.VaildPeroid = time.Now()
		fmt.Println("=============>", &user)
		//l.svcCtx.UserModel.Insert(l.ctx, &user)
		l.svcCtx.UserModel.TranInsert(tx, &user)
		return nil
	}); err != nil {
		return nil, status.Error(codes.Aborted, err.Error())
	}
	return &users.AddUserResp{}, nil
}
