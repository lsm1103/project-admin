package dataModelToApi

import (
	"fmt"
	"os"
	"testing"
)

//测试解析选择
func TestStartBuild(t *testing.T) {
	fmt.Println(os.Getwd())
	//DataModelToApi{
	//	ServiceInfo{
	//		Title:   "项目4",
	//		Desc:    "测试项目4；通过api设计文档自动生成服务，并根据api文件配置的字段mock规则进行mock生成结果",
	//		Author:  "lsm",
	//		Email:   "18370872400@163.com",
	//		Version: "v0.1.1",
	//
	//		ProjectName: "project4",
	//		ServiceType: "mock",
	//		Host:        "0.0.0.0",
	//		Port:        "814",
	//		DataSource: "root:pujian123@tcp(172.16.10.183:4306)/project-admin",
	//		CacheHost: "172.16.10.183:6379",
	//		//DataSource: "root:lsm.1018@tcp(127.0.0.1:3306)/project-admin",
	//		//CacheHost:  "127.0.0.1:6379",
	//	}, SqlParseCfg{
	//		filename: "/Users/xm/Desktop/go_package/project-admin/deploy/init.sql",
	//		database: "",
	//		strict:   false,
	//	},
	//}.StartBuild()

	DataModelToApi{
		"/Users/xm/Desktop/go_package/project-admin",
		ServiceInfo{
			Title:   "项目管理服务",
			Desc:    "对研发项目进行管理，包括代码生成、mock服务生成、cicd等；如通过api设计文档自动生成服务，并根据api文件配置的字段mock规则进行mock生成结果",
			Author:  "lsm",
			Email:   "18370872400@163.com",
			Version: "v0.1.1",

			ProjectName: "project-admin",
			ServiceType: "admin",
			Host:        "0.0.0.0",
			Port:        "814",
			DataSource: "root:pujian123@tcp(172.16.10.183:4306)/project-admin",
			CacheHost: "172.16.10.183:6379",
			//DataSource: "root:lsm.1018@tcp(127.0.0.1:3306)/project-admin",
			//CacheHost:  "127.0.0.1:6379",
		}, SqlParseCfg{
			Filename: "/Users/xm/Desktop/go_package/project-admin/deploy/init.sql",
			Database: "",
			Strict: false,
		},
	}.StartBuild()

	t.Logf("r")
}
