package logic

import (
	"app/util"
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/sqlx"

	"app/api/internal/svc"
	"app/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateStuLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateStuLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateStuLogic {
	return &UpdateStuLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateStuLogic) UpdateStu(req *types.UpdateStuReq) (resp *types.UpdateStuResp, err error) {
	resp = new(types.UpdateStuResp)
	one, err := l.svcCtx.StuModel.FindOne(l.ctx, int64(req.Id))
	if err != nil || err == sqlx.ErrNotFound {
		resp.Success = false
		resp.Message = fmt.Sprintf("student %d not exist, update exit!", req.Id)
		return resp, nil
	}

	one.Age = util.ToNullInt(req.Age)
	one.City = util.ToNullString(req.City)
	one.Name = req.Name

	err = l.svcCtx.StuModel.Update(l.ctx, one)
	if err != nil {
		resp.Success = false
		resp.Message = fmt.Sprintf("update %d failed", req.Id)
		return resp, nil
	}
	resp.Success = true
	return
}
