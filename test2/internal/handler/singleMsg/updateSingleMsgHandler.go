package singleMsg

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"project-admin/test2/internal/logic/singleMsg"
	"project-admin/test2/internal/svc"
	"project-admin/test2/internal/types"

	"project-admin/common/result"
)

func UpdateSingleMsgHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UpdateSingleMsg
		if err := httpx.Parse(r, &req); err != nil {
			// httpx.Error(w, err)
			result.ParamErrorResult(r, w, err)
			return
		}

		l := singleMsg.NewUpdateSingleMsgLogic(r.Context(), svcCtx)
		err := l.UpdateSingleMsg(&req)
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
