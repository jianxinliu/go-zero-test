package logic

import (
	"app/api/internal/svc"
	"app/api/internal/types"
	"app/sql/model"
	"app/util"
	"context"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddStuLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddStuLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddStuLogic {
	return &AddStuLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddStuLogic) AddStu(req *types.AddStuReq) (resp *types.AddStuResp, err error) {
	resp = new(types.AddStuResp)

	_, err = l.svcCtx.StuModel.Insert(l.ctx, &model.Student{
		Name: req.Name,
		Age:  util.ToNullInt(req.Age),
		City: util.ToNullString(req.City),
	})
	if err != nil {
		resp.Success = false
		resp.Message = err.Error()
		return resp, nil
	}
	resp.Success = true
	return
}
