package logic

import (
	"context"
	"database/sql"
	"fmt"
	"user/internal/svc"
	"user/pb/users"

	"github.com/jinzhu/copier"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserByCodeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserByCodeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserByCodeLogic {
	return &GetUserByCodeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserByCodeLogic) GetUserByCode(in *users.GetUserByCodeReq) (*users.GetUserByIdResp, error) {
	fmt.Println(">>>>>>>>>>>>GetUserByCode")
	var code sql.NullString
	code.String = in.Code
	user, _ := l.svcCtx.UserModel.FindOneByCode(l.ctx, code)
	resp := &users.GetUserByIdResp{
		User: &users.User{},
	}
	copier.Copy(resp.User, user)
	return resp, nil
}
