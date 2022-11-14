package svc

import (
    "github.com/zeromicro/go-zero/core/stores/sqlx"

	{{.configImport}}
	dataModel "{{.commonPkgPath}}/dataModel/{{.projectName}}Model"
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
