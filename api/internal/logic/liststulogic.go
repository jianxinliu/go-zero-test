package logic

import (
	"app/sql/model"
	"app/util"
	"context"

	"app/api/internal/svc"
	"app/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListStuLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListStuLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListStuLogic {
	return &ListStuLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListStuLogic) ListStu(req *types.ListStuReq) (resp *types.ListStuResp, err error) {
	resp = new(types.ListStuResp)

	list, err := l.svcCtx.StuModel.FindAllBy(l.ctx, req.With, &model.Student{
		Name: req.Name,
		City: util.ToNullString(req.City),
	})
	if err != nil {
		resp.Message = err.Error()
		resp.List = []types.Student{}
		return resp, nil
	}
	// fixme: 怎样才能不做这种类型转换呢
	var retList []types.Student
	for _, stu := range list {
		retList = append(retList, types.Student{
			Name: stu.Name,
			//  fixme: 怎样才能不做这种类型转换呢
			Age:  int(stu.Age.Int64),
			City: stu.City.String,
			Id:   int(stu.Id),
		})
	}
	resp.List = retList

	return
}
