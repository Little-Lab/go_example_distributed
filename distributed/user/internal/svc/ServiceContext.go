package svc

import (
	"database/sql"
	"user/internal/config"
	"user/internal/model"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config    config.Config
	Db        *sql.DB
	UserModel model.UserModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	db, _ := sqlx.NewMysql(c.DB.DataSource).RawDB()
	return &ServiceContext{
		Config:    c,
		Db:        db,
		UserModel: model.NewUserModel(sqlx.NewMysql(c.DB.DataSource)),
	}
}
