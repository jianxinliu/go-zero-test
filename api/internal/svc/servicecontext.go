package svc

import (
	"app/api/internal/config"
	"app/sql/model"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config   config.Config
	SqlConn  sqlx.SqlConn
	StuModel model.StudentModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.DataSource)
	return &ServiceContext{
		Config:   c,
		SqlConn:  conn,
		StuModel: model.NewStudentModel(conn, c.Cache),
	}
}
