package logic

import (
	"app/rpc/app"
	"context"

	"app/api/internal/svc"
	"app/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListStuByRpcLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListStuByRpcLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListStuByRpcLogic {
	return &ListStuByRpcLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListStuByRpcLogic) ListStuByRpc(req *types.ListStuReq) (resp *types.ListStuResp, err error) {
	resp = new(types.ListStuResp)
	resp.List = []*types.Student{}

	l.Logger.Infof("s %s ps %p", req.City, &req.City)

	list, err := l.svcCtx.AppRpc.ListStu(l.ctx, &app.StuListReq{
		Name: req.Name,
		City: &req.City,
	})
	l.Logger.Info("list by rpc;.....")
	if err != nil {
		resp.Message = err.Error()
		return resp, nil
	}
	var retList []*types.Student
	for _, stu := range list.List {
		retList = append(retList, &types.Student{
			Name: stu.Name,
			Age:  int(stu.Age),
			City: stu.City,
		})
	}
	if len(retList) > 0 {
		resp.List = retList
	}
	return
}
