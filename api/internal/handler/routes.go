// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	"app/api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/add",
				Handler: AddStuHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/update/:id",
				Handler: UpdateStuHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/list",
				Handler: ListStuHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/list_rpc",
				Handler: ListStuByRpcHandler(serverCtx),
			},
		},
		rest.WithPrefix("/app/v1/stu"),
	)
}
