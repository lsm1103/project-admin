package {{.PkgName}}

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	{{.ImportPackages}}
	"{{.RootPkgName}}/common/result"
	{{if eq .ServiceType "admin"}}{{if eq .HandlerType "gets"}}"{{.RootPkgName}}/common/sqlUtils"{{end}}{{end}}
)

func {{.HandlerName}}(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		{{if .HasRequest}}{{if eq .ServiceType "admin"}}{{if eq .HandlerType "gets"}}
		var req sqlUtils.GetsReq
		var _ types.GetsReq
		{{else}}var req types.{{.RequestType}}{{end}}
		{{else}}var req types.{{.RequestType}}
		{{end}}
		if err := httpx.Parse(r, &req); err != nil {
			// httpx.Error(w, err)
			result.ParamErrorResult(r,w,err)
			return
		}{{end}}

		l := {{.LogicName}}.New{{.LogicType}}(r.Context(), svcCtx)
		{{if .HasResp}}resp, {{end}}err := l.{{.Call}}({{if .HasRequest}}&req{{end}})
		/* 
		if err != nil {
			httpx.Error(w, err)
		} else {
			{{if .HasResp}}httpx.OkJson(w, resp){{else}}httpx.Ok(w){{end}}
		}
		*/
		result.HttpResult(r,w,err,{{if .HasResp}}resp{{else}}nil{{end}})
	}
}
