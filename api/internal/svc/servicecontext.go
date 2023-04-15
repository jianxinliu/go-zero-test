package svc

import (
	"app/api/internal/config"
	"app/rpc/app"
	"app/rpc/appserver"
	"app/sql/model"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config   config.Config
	SqlConn  sqlx.SqlConn
	StuModel model.StudentModel

	AppRpc appserver.AppClient
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.DataSource)
	return &ServiceContext{
		Config:   c,
		SqlConn:  conn,
		StuModel: model.NewStudentModel(conn, c.Cache),

		AppRpc: app.NewApp(zrpc.MustNewClient(c.AppRpcConf)),
	}
}
