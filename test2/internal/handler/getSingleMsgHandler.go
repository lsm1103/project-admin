package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"project-admin/test2/internal/logic"
	"project-admin/test2/internal/svc"
	"project-admin/test2/internal/types"

	"project-admin/common/result"
)

func GetSingleMsgHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetSingleMsg
		if err := httpx.Parse(r, &req); err != nil {
			// httpx.Error(w, err)
			result.ParamErrorResult(r, w, err)
			return
		}

		l := logic.NewGetSingleMsgLogic(r.Context(), svcCtx)
		resp, err := l.GetSingleMsg(&req)
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
