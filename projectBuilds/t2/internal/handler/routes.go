// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	Friend "project-admin/projectBuilds/t2/internal/handler/Friend"
	Group "project-admin/projectBuilds/t2/internal/handler/Group"
	GroupMsg "project-admin/projectBuilds/t2/internal/handler/GroupMsg"
	OfflineMsg "project-admin/projectBuilds/t2/internal/handler/OfflineMsg"
	SingleMsg "project-admin/projectBuilds/t2/internal/handler/SingleMsg"
	UserGroup "project-admin/projectBuilds/t2/internal/handler/UserGroup"
	"project-admin/projectBuilds/t2/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/create",
				Handler: Group.CreateHandler(serverCtx),
			},
			{
				Method:  http.MethodPut,
				Path:    "/update",
				Handler: Group.UpdateHandler(serverCtx),
			},
			{
				Method:  http.MethodDelete,
				Path:    "/delete",
				Handler: Group.DeleteHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/get",
				Handler: Group.GetHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/gets",
				Handler: Group.GetsHandler(serverCtx),
			},
		},
		rest.WithPrefix("/admin/Group/v1"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/create",
				Handler: UserGroup.CreateHandler(serverCtx),
			},
			{
				Method:  http.MethodPut,
				Path:    "/update",
				Handler: UserGroup.UpdateHandler(serverCtx),
			},
			{
				Method:  http.MethodDelete,
				Path:    "/delete",
				Handler: UserGroup.DeleteHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/get",
				Handler: UserGroup.GetHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/gets",
				Handler: UserGroup.GetsHandler(serverCtx),
			},
		},
		rest.WithPrefix("/admin/UserGroup/v1"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/create",
				Handler: Friend.CreateHandler(serverCtx),
			},
			{
				Method:  http.MethodPut,
				Path:    "/update",
				Handler: Friend.UpdateHandler(serverCtx),
			},
			{
				Method:  http.MethodDelete,
				Path:    "/delete",
				Handler: Friend.DeleteHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/get",
				Handler: Friend.GetHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/gets",
				Handler: Friend.GetsHandler(serverCtx),
			},
		},
		rest.WithPrefix("/admin/Friend/v1"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/create",
				Handler: OfflineMsg.CreateHandler(serverCtx),
			},
			{
				Method:  http.MethodPut,
				Path:    "/update",
				Handler: OfflineMsg.UpdateHandler(serverCtx),
			},
			{
				Method:  http.MethodDelete,
				Path:    "/delete",
				Handler: OfflineMsg.DeleteHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/get",
				Handler: OfflineMsg.GetHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/gets",
				Handler: OfflineMsg.GetsHandler(serverCtx),
			},
		},
		rest.WithPrefix("/admin/OfflineMsg/v1"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/create",
				Handler: SingleMsg.CreateHandler(serverCtx),
			},
			{
				Method:  http.MethodPut,
				Path:    "/update",
				Handler: SingleMsg.UpdateHandler(serverCtx),
			},
			{
				Method:  http.MethodDelete,
				Path:    "/delete",
				Handler: SingleMsg.DeleteHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/get",
				Handler: SingleMsg.GetHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/gets",
				Handler: SingleMsg.GetsHandler(serverCtx),
			},
		},
		rest.WithPrefix("/admin/SingleMsg/v1"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/create",
				Handler: GroupMsg.CreateHandler(serverCtx),
			},
			{
				Method:  http.MethodPut,
				Path:    "/update",
				Handler: GroupMsg.UpdateHandler(serverCtx),
			},
			{
				Method:  http.MethodDelete,
				Path:    "/delete",
				Handler: GroupMsg.DeleteHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/get",
				Handler: GroupMsg.GetHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/gets",
				Handler: GroupMsg.GetsHandler(serverCtx),
			},
		},
		rest.WithPrefix("/admin/GroupMsg/v1"),
	)
}
