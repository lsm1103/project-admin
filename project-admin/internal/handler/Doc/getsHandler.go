package Doc

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"project-admin/project-admin/internal/logic/Doc"
	"project-admin/project-admin/internal/svc"
	"project-admin/project-admin/internal/types"

	"project-admin/common/result"
	"project-admin/common/sqlUtils"
)

func GetsHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var req sqlUtils.GetsReq
		var _ types.GetsReq

		if err := httpx.Parse(r, &req); err != nil {
			// httpx.Error(w, err)
			result.ParamErrorResult(r, w, err)
			return
		}

		l := Doc.NewGetsLogic(r.Context(), svcCtx)
		resp, err := l.Gets(&req)
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
