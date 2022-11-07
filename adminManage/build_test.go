package adminManage

import (
	"fmt"
	"os"
	"testing"
)

//测试解析选择
func TestStartBuild(t *testing.T) {
	fmt.Print(os.Getwd())
	StartBuild(ServiceInfo{
		Title:   "项目1",
		Desc:    "通过api设计文档自动生成服务，并根据api文件配置的字段mock规则进行mock生成结果",
		Author:  "lsm",
		Email:   "18370872400@163.com",
		Version: "v0.1.1",

		Host: "0.0.0.0",
		Port: "801",
		CommonPkgPath: "project-admin",
		DataSource: "root:pujian123@tcp(172.16.10.183:4306)/im-center",
		CacheHost: "172.16.10.183:6379",
	},SqlParseCfg{
		filename: "/Users/xm/Desktop/go_package/project-admin/deploy/init.sql",
		database: "",
		strict:   false,
	},
	"service3.api")

	t.Logf("r")
}