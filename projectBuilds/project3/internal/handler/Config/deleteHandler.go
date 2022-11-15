package Config

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"project-admin/projectBuilds/project3/internal/logic/Config"
	"project-admin/projectBuilds/project3/internal/svc"
	"project-admin/projectBuilds/project3/internal/types"

	"project-admin/common/result"
)

func DeleteHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.DeleteReq

		if err := httpx.Parse(r, &req); err != nil {
			// httpx.Error(w, err)
			result.ParamErrorResult(r, w, err)
			return
		}

		l := Config.NewDeleteLogic(r.Context(), svcCtx)
		err := l.Delete(&req)
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
