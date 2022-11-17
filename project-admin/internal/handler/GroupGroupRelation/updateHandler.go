package GroupGroupRelation

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"project-admin/project-admin/internal/logic/GroupGroupRelation"
	"project-admin/project-admin/internal/svc"
	"project-admin/project-admin/internal/types"

	"project-admin/common/result"
)

func UpdateHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UpdateGroupGroupRelationReq

		if err := httpx.Parse(r, &req); err != nil {
			// httpx.Error(w, err)
			result.ParamErrorResult(r, w, err)
			return
		}

		l := GroupGroupRelation.NewUpdateLogic(r.Context(), svcCtx)
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
