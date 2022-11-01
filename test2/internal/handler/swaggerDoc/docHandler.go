package swaggerDoc

import (
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"

	"project-admin/test2/internal/logic/swaggerDoc"
	"project-admin/test2/internal/svc"
)

func DocHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		l := swaggerDoc.NewDocLogic(r.Context(), svcCtx)
		err := l.Doc(w)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.Ok(w)
		}
	}
}
