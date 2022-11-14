package svc

import (
    "github.com/zeromicro/go-zero/core/stores/sqlx"

	{{.configImport}}
	dataModel "{{.rootPkgName}}/dataModel/{{.projectName}}"
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
