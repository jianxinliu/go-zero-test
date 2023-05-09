package logic

import (
	"context"

	"app/rpc/appserver"
	"app/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListStuLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewListStuLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListStuLogic {
	return &ListStuLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ListStuLogic) ListStu(in *appserver.StuListReq) (*appserver.StuListResp, error) {
	l.Logger.Info("list all stu by rpc.....")
	if *in.City == "beijing" {
		l.Logger.Infof("sss %s ps %p", *in.City+"__as", in.City)
	}
	l.Logger.Infof("City: %s %v", in.City, in.City)
	return &appserver.StuListResp{
		List: []*appserver.Student{},
	}, nil
}
