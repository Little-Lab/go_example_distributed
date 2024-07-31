package logic

import (
	"context"
	"database/sql"
	"time"
	"user/internal/model"

	"user/internal/svc"
	"user/pb/users"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddUserLogic {
	return &AddUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// -----------------------user-----------------------
func (l *AddUserLogic) AddUser(in *users.AddUserReq) (*users.AddUserResp, error) {
	user := new(model.User)
	user.Email.String = in.Email
	user.Avatar = sql.NullString{"1", true}
	user.AlipayNickname = "a"
	user.VaildPeroid = time.Now()
	l.svcCtx.UserModel.Insert(l.ctx, user)
	//if err := barrier.CallWithDB(l.svcCtx.Db, fun
	return &users.AddUserResp{}, nil
}
