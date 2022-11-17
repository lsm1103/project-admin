// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	Application "project-admin/project-admin/internal/handler/Application"
	ApplicationConfig "project-admin/project-admin/internal/handler/ApplicationConfig"
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
				Handler: GroupGroupRelation.CreateHandler(serverCtx),
			},
			{
				Method:  http.MethodPut,
				Path:    "/update",
				Handler: GroupGroupRelation.UpdateHandler(serverCtx),
			},
			{
				Method:  http.MethodDelete,
				Path:    "/delete",
				Handler: GroupGroupRelation.DeleteHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/get",
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
				Path:    "/create",
				Handler: Config.CreateHandler(serverCtx),
			},
			{
				Method:  http.MethodPut,
				Path:    "/update",
				Handler: Config.UpdateHandler(serverCtx),
			},
			{
				Method:  http.MethodDelete,
				Path:    "/delete",
				Handler: Config.DeleteHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/get",
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
				Path:    "/create",
				Handler: Project.CreateHandler(serverCtx),
			},
			{
				Method:  http.MethodPut,
				Path:    "/update",
				Handler: Project.UpdateHandler(serverCtx),
			},
			{
				Method:  http.MethodDelete,
				Path:    "/delete",
				Handler: Project.DeleteHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/get",
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
				Path:    "/create",
				Handler: Application.CreateHandler(serverCtx),
			},
			{
				Method:  http.MethodPut,
				Path:    "/update",
				Handler: Application.UpdateHandler(serverCtx),
			},
			{
				Method:  http.MethodDelete,
				Path:    "/delete",
				Handler: Application.DeleteHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/get",
				Handler: Application.GetHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/gets",
				Handler: Application.GetsHandler(serverCtx),
			},
		},
		rest.WithPrefix("/admin/Application/v1"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/create",
				Handler: ApplicationConfig.CreateHandler(serverCtx),
			},
			{
				Method:  http.MethodPut,
				Path:    "/update",
				Handler: ApplicationConfig.UpdateHandler(serverCtx),
			},
			{
				Method:  http.MethodDelete,
				Path:    "/delete",
				Handler: ApplicationConfig.DeleteHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/get",
				Handler: ApplicationConfig.GetHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/gets",
				Handler: ApplicationConfig.GetsHandler(serverCtx),
			},
		},
		rest.WithPrefix("/admin/ApplicationConfig/v1"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/create",
				Handler: Doc.CreateHandler(serverCtx),
			},
			{
				Method:  http.MethodPut,
				Path:    "/update",
				Handler: Doc.UpdateHandler(serverCtx),
			},
			{
				Method:  http.MethodDelete,
				Path:    "/delete",
				Handler: Doc.DeleteHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/get",
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
				Path:    "/create",
				Handler: DocHistory.CreateHandler(serverCtx),
			},
			{
				Method:  http.MethodPut,
				Path:    "/update",
				Handler: DocHistory.UpdateHandler(serverCtx),
			},
			{
				Method:  http.MethodDelete,
				Path:    "/delete",
				Handler: DocHistory.DeleteHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/get",
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
