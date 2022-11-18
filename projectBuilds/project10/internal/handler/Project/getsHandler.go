package Project

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"project-admin/projectBuilds/project10/internal/logic/Project"
	"project-admin/projectBuilds/project10/internal/svc"
	"project-admin/projectBuilds/project10/internal/types"
)

func GetsHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetsReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := Project.NewGetsLogic(r.Context(), svcCtx)
		resp, err := l.Gets(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
