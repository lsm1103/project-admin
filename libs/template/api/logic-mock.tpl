package {{.pkgName}}

import (
	{{.imports}}
	{{if .respType}}"project-admin/common/mocks"{{end}}
)

type {{.logic}} struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func New{{.logic}}(ctx context.Context, svcCtx *svc.ServiceContext) {{.logic}} {
	return {{.logic}}{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *{{.logic}}) {{.function}}({{.request}}) {{.responseType}} {
    // 方便前端调试的接口mock
	{{if .respType}}resp = &types.{{.respType}}{}
    mocks.RespMock(resp){{end}}
	{{.returnString}}
}