package Application

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"project-admin/projectBuilds/project24/internal/logic/Application"
	"project-admin/projectBuilds/project24/internal/svc"
	"project-admin/projectBuilds/project24/internal/types"

	"project-admin/common/result"
)

func UpdateHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UpdateApplicationReq

		if err := httpx.Parse(r, &req); err != nil {
			// httpx.Error(w, err)
			result.ParamErrorResult(r, w, err)
			return
		}

		l := Application.NewUpdateLogic(r.Context(), svcCtx)
		resp, err := l.Update(&req)
		/*
			if err != nil {
				httpx.Error(w, err)
			} else {
				httpx.OkJson(w, resp)
			}
		*/
		result.HttpResult(r, w, err, resp)
	}
}
