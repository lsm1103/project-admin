package GroupGroupRelation

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"project-admin/projectBuilds/project1/internal/logic/GroupGroupRelation"
	"project-admin/projectBuilds/project1/internal/svc"
	"project-admin/projectBuilds/project1/internal/types"

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

		l := GroupGroupRelation.NewGetLogic(r.Context(), svcCtx)
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
