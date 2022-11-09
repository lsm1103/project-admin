package {{.pkgName}}

import (
	{{.imports}}

	{{if eq .handlerType "create"}}"project-admin/common/uniqueid"{{end}}
	{{if eq .handlerType "create"}}"project-admin/dataModel"{{end}}
	{{if eq .handlerType "create"}}"github.com/jinzhu/copier"{{end}}
	{{if eq .handlerType "update"}}dataModel "project-admin/dataModel/{{projectName}}Model"{{end}}
	{{if eq .handlerType "update"}}"github.com/jinzhu/copier"{{end}}
	{{if eq .handlerType "gets"}}"project-admin/common/sqlUtils"{{end}}
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

func (l *{{.logic}}) {{.function}}(req *sqlUtils.GetsReq) {{.responseType}} {
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
    {{ else if eq .handlerType "get" }}err = l.svcCtx.{{.moduleName}}Model.FindOne(l.ctx, nil, req.Id, resp)
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