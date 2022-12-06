package ApplicationConfig

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"project-admin/project-admin/internal/logic/ApplicationConfig"
	"project-admin/project-admin/internal/svc"
	"project-admin/project-admin/internal/types"

	"project-admin/common/result"
)

func CreateHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CreateApplicationConfigReq

		if err := httpx.Parse(r, &req); err != nil {
			// httpx.Error(w, err)
			result.ParamErrorResult(r, w, err)
			return
		}

		l := ApplicationConfig.NewCreateLogic(r.Context(), svcCtx)
		resp, err := l.Create(&req)
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
