package svc

import (
	{{.configImport}}
    {{if eq .serviceType "admin"}}
    "github.com/zeromicro/go-zero/core/stores/sqlx"
	dataModel "{{.rootPkgName}}/dataModel/{{.projectName}}"{{end}}
)

type ServiceContext struct {
	Config {{.config}}
	{{.middleware}}
	{{.typeFields}}
}

func NewServiceContext(c {{.config}}) *ServiceContext {
	return &ServiceContext{
		Config: c, 
		{{.middlewareAssignment}}
		{{.valueFields}}
	}
}
