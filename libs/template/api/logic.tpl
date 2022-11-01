package {{.pkgName}}

import (
	{{.imports}}

	{{if eq .handlerType "create"}}"{{.rootPkgName}}/common/uniqueid"{{end}}
	{{if eq .handlerType "create"}}"{{.rootPkgName}}/dataModel"{{end}}
	{{if eq .handlerType "create"}}"github.com/jinzhu/copier"{{end}}
	{{if eq .handlerType "update"}}"{{.rootPkgName}}/dataModel"{{end}}
	{{if eq .handlerType "update"}}"github.com/jinzhu/copier"{{end}}
	{{if eq .handlerType "get"}}"{{.rootPkgName}}/common/xerr"
	"github.com/zeromicro/go-zero/core/stores/sqlx"{{end}}
	{{if eq .handlerType "gets"}}"{{.rootPkgName}}/common/sqlUtils"{{end}}
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

func (l *{{.logic}}) {{.function}}({{if eq .handlerType "gets"}}req *sqlUtils.GetsReq{{else}}{{.request}}{{end}}) {{.responseType}} {
    // 自动生成的后台管理接口  {{.request}}
    {{ if eq .handlerType "create" }}sqlReq := &dataModel.{{.moduleName}}{}
    err := copier.Copy(sqlReq, req)
    if err != nil {
        return err
    }
    sqlReq.Id = uniqueid.GenId()
    _, err = l.svcCtx.{{.moduleName}}Model.Insert(l.ctx, nil, sqlReq)
    if err != nil {
        return err
    }
    {{ else if eq .handlerType "delete" }}err := l.svcCtx.{{.moduleName}}Model.Delete(l.ctx, nil, req.Id)
    if err != nil {
        return err
    }
    {{ else if eq .handlerType "update" }}sqlReq := &dataModel.{{.moduleName}}{}
    err := copier.Copy(sqlReq, req)
    if err != nil {
        return err
    }
    err = l.svcCtx.{{.moduleName}}Model.Update(l.ctx, nil, sqlReq)
    if err != nil {
        return err
    }
    {{ else if eq .handlerType "get" }}resp  = &types.{{.moduleName}}{}
    err = l.svcCtx.{{.moduleName}}Model.FindOne(l.ctx, nil, req.Id, resp)
    if err == sqlx.ErrNotFound {
        return nil, xerr.NewErrCode(xerr.DATA_NOT_FIND)
    }
    if err != nil {
        return nil, err
    }
    {{ else if eq .handlerType "gets" }}resp = &types.{{.respType}}{ Current: req.Current, PageSize: req.PageSize}
    err = l.svcCtx.{{.moduleName}}Model.FindAll(req, &resp.List)
    if err != nil {
        return nil, err
    }
    if int64(len(resp.List)) > req.PageSize {
        resp.IsNext = true
        resp.List = resp.List[:req.PageSize]
    }{{ end }}
    {{.returnString}}
}