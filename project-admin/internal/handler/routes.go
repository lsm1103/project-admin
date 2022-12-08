// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	Application "project-admin/project-admin/internal/handler/Application"
	ApplicationInfo "project-admin/project-admin/internal/handler/ApplicationInfo"
	Config "project-admin/project-admin/internal/handler/Config"
	Doc "project-admin/project-admin/internal/handler/Doc"
	DocHistory "project-admin/project-admin/internal/handler/DocHistory"
	Group "project-admin/project-admin/internal/handler/Group"
	GroupGroupRelation "project-admin/project-admin/internal/handler/GroupGroupRelation"
	Project "project-admin/project-admin/internal/handler/Project"
	UserGroup "project-admin/project-admin/internal/handler/UserGroup"
	"project-admin/project-admin/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/",
				Handler: Group.CreateHandler(serverCtx),
			},
			{
				Method:  http.MethodPut,
				Path:    "/",
				Handler: Group.UpdateHandler(serverCtx),
			},
			{
				Method:  http.MethodDelete,
				Path:    "/",
				Handler: Group.DeleteHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/:id",
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
				Path:    "/",
				Handler: UserGroup.CreateHandler(serverCtx),
			},
			{
				Method:  http.MethodPut,
				Path:    "/",
				Handler: UserGroup.UpdateHandler(serverCtx),
			},
			{
				Method:  http.MethodDelete,
				Path:    "/",
				Handler: UserGroup.DeleteHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/:id",
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
				Path:    "/",
				Handler: GroupGroupRelation.CreateHandler(serverCtx),
			},
			{
				Method:  http.MethodPut,
				Path:    "/",
				Handler: GroupGroupRelation.UpdateHandler(serverCtx),
			},
			{
				Method:  http.MethodDelete,
				Path:    "/",
				Handler: GroupGroupRelation.DeleteHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/:id",
				Handler: GroupGroupRelation.GetHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/gets",
				Handler: GroupGroupRelation.GetsHandler(serverCtx),
			},
		},
		rest.WithPrefix("/admin/GroupGroupRelation/v1"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/",
				Handler: Config.CreateHandler(serverCtx),
			},
			{
				Method:  http.MethodPut,
				Path:    "/",
				Handler: Config.UpdateHandler(serverCtx),
			},
			{
				Method:  http.MethodDelete,
				Path:    "/",
				Handler: Config.DeleteHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/:id",
				Handler: Config.GetHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/gets",
				Handler: Config.GetsHandler(serverCtx),
			},
		},
		rest.WithPrefix("/admin/Config/v1"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/",
				Handler: Project.CreateHandler(serverCtx),
			},
			{
				Method:  http.MethodPut,
				Path:    "/",
				Handler: Project.UpdateHandler(serverCtx),
			},
			{
				Method:  http.MethodDelete,
				Path:    "/",
				Handler: Project.DeleteHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/:id",
				Handler: Project.GetHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/gets",
				Handler: Project.GetsHandler(serverCtx),
			},
		},
		rest.WithPrefix("/admin/Project/v1"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/",
				Handler: Application.CreateHandler(serverCtx),
			},
			{
				Method:  http.MethodPut,
				Path:    "/",
				Handler: Application.UpdateHandler(serverCtx),
			},
			{
				Method:  http.MethodDelete,
				Path:    "/",
				Handler: Application.DeleteHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/:id",
				Handler: Application.GetHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/gets",
				Handler: Application.GetsHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/build",
				Handler: Application.BuildHandler(serverCtx),
			},
		},
		rest.WithPrefix("/admin/Application/v1"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/",
				Handler: ApplicationInfo.CreateHandler(serverCtx),
			},
			{
				Method:  http.MethodPut,
				Path:    "/",
				Handler: ApplicationInfo.UpdateHandler(serverCtx),
			},
			{
				Method:  http.MethodDelete,
				Path:    "/",
				Handler: ApplicationInfo.DeleteHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/:id",
				Handler: ApplicationInfo.GetHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/gets",
				Handler: ApplicationInfo.GetsHandler(serverCtx),
			},
		},
		rest.WithPrefix("/admin/ApplicationInfo/v1"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/",
				Handler: Doc.CreateHandler(serverCtx),
			},
			{
				Method:  http.MethodPut,
				Path:    "/",
				Handler: Doc.UpdateHandler(serverCtx),
			},
			{
				Method:  http.MethodDelete,
				Path:    "/",
				Handler: Doc.DeleteHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/:id",
				Handler: Doc.GetHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/gets",
				Handler: Doc.GetsHandler(serverCtx),
			},
		},
		rest.WithPrefix("/admin/Doc/v1"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/",
				Handler: DocHistory.CreateHandler(serverCtx),
			},
			{
				Method:  http.MethodPut,
				Path:    "/",
				Handler: DocHistory.UpdateHandler(serverCtx),
			},
			{
				Method:  http.MethodDelete,
				Path:    "/",
				Handler: DocHistory.DeleteHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/:id",
				Handler: DocHistory.GetHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/gets",
				Handler: DocHistory.GetsHandler(serverCtx),
			},
		},
		rest.WithPrefix("/admin/DocHistory/v1"),
	)
}
