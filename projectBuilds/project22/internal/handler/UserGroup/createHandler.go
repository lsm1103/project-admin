package UserGroup

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"project-admin/projectBuilds/project22/internal/logic/UserGroup"
	"project-admin/projectBuilds/project22/internal/svc"
	"project-admin/projectBuilds/project22/internal/types"

	"project-admin/common/result"
)

func CreateHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CreateUserGroupReq

		if err := httpx.Parse(r, &req); err != nil {
			// httpx.Error(w, err)
			result.ParamErrorResult(r, w, err)
			return
		}

		l := UserGroup.NewCreateLogic(r.Context(), svcCtx)
		err := l.Create(&req)
		/*
			if err != nil {
				httpx.Error(w, err)
			} else {
				httpx.Ok(w)
			}
		*/
		result.HttpResult(r, w, err, nil)
	}
}
