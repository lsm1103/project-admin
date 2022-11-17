package Config

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"project-admin/project-admin/internal/logic/Config"
	"project-admin/project-admin/internal/svc"
	"project-admin/project-admin/internal/types"

	"project-admin/common/result"
)

func GetHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetReq

		if err := httpx.Parse(r, &req); err != nil {
			// httpx.Error(w, err)
			result.ParamErrorResult(r, w, err)
			return
		}

		l := Config.NewGetLogic(r.Context(), svcCtx)
		resp, err := l.Get(&req)
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
