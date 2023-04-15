package handler

import (
	"net/http"

	"app/api/internal/logic"
	"app/api/internal/svc"
	"app/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func UpdateStuHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UpdateStuReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewUpdateStuLogic(r.Context(), svcCtx)
		resp, err := l.UpdateStu(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
