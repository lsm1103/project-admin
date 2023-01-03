package DocHistory

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"project-admin/projectBuilds/project-24/internal/logic/DocHistory"
	"project-admin/projectBuilds/project-24/internal/svc"
	"project-admin/projectBuilds/project-24/internal/types"

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

		l := DocHistory.NewGetsLogic(r.Context(), svcCtx)
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