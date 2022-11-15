package UserGroup

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"project-admin/projectBuilds/project2/internal/logic/UserGroup"
	"project-admin/projectBuilds/project2/internal/svc"
	"project-admin/projectBuilds/project2/internal/types"

	"project-admin/common/result"
)

func UpdateHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UpdateUserGroupReq

		if err := httpx.Parse(r, &req); err != nil {
			// httpx.Error(w, err)
			result.ParamErrorResult(r, w, err)
			return
		}

		l := UserGroup.NewUpdateLogic(r.Context(), svcCtx)
		err := l.Update(&req)
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
