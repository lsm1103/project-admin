package dataModelToApi

import (
	"fmt"
	"project-admin/common/dataModelToApi/parser"
	"testing"
)

func TestBuildWebCode(t *testing.T) {
	RootPkgPath := "/Users/xm/Desktop/go_package/project-admin"
	//解析数据库sql文件
	tables, err := parser.Parse(fmt.Sprintf("%s/deploy/init.sql", RootPkgPath),"", false)
	if err != nil {
		t.Error(err)
	}

	BuildWeb{
		RootPkgPath: RootPkgPath,
		TableList:   tables,
		ServiceCfg:  ServiceInfo{},
	}.BuildWebCode()
	
	t.Logf("r")
}