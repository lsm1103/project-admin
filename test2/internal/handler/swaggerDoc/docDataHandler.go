package swaggerDoc

import (
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"

	"project-admin/test2/internal/logic/swaggerDoc"
	"project-admin/test2/internal/svc"
)

func DocDataHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		l := swaggerDoc.NewDocDataLogic(r.Context(), svcCtx)
		resp, err := l.DocData()
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
