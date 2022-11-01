package handler

import (
	"net/http"
	"project-admin/common/sqlUtils"

	"github.com/zeromicro/go-zero/rest/httpx"
	"project-admin/common/result"
	"project-admin/test2/internal/logic"
	"project-admin/test2/internal/svc"
)

func GetsSingleMsgHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req sqlUtils.GetsReq
		if err := httpx.Parse(r, &req); err != nil {
			// httpx.Error(w, err)
			result.ParamErrorResult(r, w, err)
			return
		}

		l := logic.NewGetsSingleMsgLogic(r.Context(), svcCtx)
		resp, err := l.GetsSingleMsg(&req)
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
