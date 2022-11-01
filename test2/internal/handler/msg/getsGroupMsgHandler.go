package msg

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"project-admin/test2/internal/logic/msg"
	"project-admin/test2/internal/svc"
	"project-admin/test2/internal/types"

	"project-admin/common/result"
)

func GetsGroupMsgHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetsReq
		if err := httpx.Parse(r, &req); err != nil {
			// httpx.Error(w, err)
			result.ParamErrorResult(r, w, err)
			return
		}

		l := msg.NewGetsGroupMsgLogic(r.Context(), svcCtx)
		resp, err := l.GetsGroupMsg(&req)
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
