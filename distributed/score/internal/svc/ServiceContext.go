package svc

import (
	"database/sql"
	"score/internal/config"
	"score/internal/model"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config     config.Config
	Db         *sql.DB
	ScoreModel model.ScoreModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	db, _ := sqlx.NewMysql(c.DB.DataSource).RawDB()
	return &ServiceContext{
		Config:     c,
		Db:         db,
		ScoreModel: model.NewScoreModel(sqlx.NewMysql(c.DB.DataSource)),
	}
}
