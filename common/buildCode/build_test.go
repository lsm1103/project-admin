package buildCode

import (
	"fmt"
	"testing"
)

func TestBuildDataModel(t *testing.T) {
	t.Log("")
	RootPkgPath := "/Users/xm/Desktop/go_package/project-admin"
	err := BuildCode(BuildCode{
		RootPkgPath: RootPkgPath,
		Info:        BuildAppInfo{
			Title:   "项目11",
			Desc:    "对研发项目进行管理，包括代码生成、mock服务生成、cicd等；如通过api设计文档自动生成服务，并根据api文件配置的字段mock规则进行mock生成结果",
			Author:  "lsm",
			Email:   "18370872400@163.com",
			Version: "v0.1.1",

			ProjectName: "project11",
			ServiceType: "admin",
			Host:        "0.0.0.0",
			Port:        "811",
			DataSource:  "root:pujian123@tcp(172.16.10.183:4306)/project-admin",
			CacheHost: 	 "172.16.10.183:6379",
			//DataSource: "root:lsm.1018@tcp(127.0.0.1:3306)/project-admin",
			//CacheHost:  "127.0.0.1:6379",
			Style:        "goZero",
			TemplatePath: fmt.Sprintf("%s/libs/template", RootPkgPath),
			Database:     "",
			DdlArg:       DdlArg{
				Src:           fmt.Sprintf("%s/deploy/init.sql", RootPkgPath),
				Cache:         true,
				Strict:        false,
				IgnoreColumns: []string{"create_at", "created_at", "create_time", "update_at", "updated_at", "update_time"},
			},
		},
	}).BuildDataModel()
	if err != nil {
		t.Log(err)
	}
}

//func BuildAPP() http.HandlerFunc {
//	info := BuildAppInfo{
//		Title:   "项目管理服务",
//		Desc:    "对研发项目进行管理，包括代码生成、mock服务生成、cicd等；如通过api设计文档自动生成服务，并根据api文件配置的字段mock规则进行mock生成结果",
//		Author:  "lsm",
//		Email:   "18370872400@163.com",
//		Version: "v0.1.1",
//
//		ProjectName: "project10",
//		ServiceType: "admin",
//		Host:        "0.0.0.0",
//		Port:        "820",
//		DataSource:  "root:pujian123@tcp(172.16.10.183:4306)/project-admin",
//		CacheHost: 	 "172.16.10.183:6379",
//		//DataSource: "root:lsm.1018@tcp(127.0.0.1:3306)/project-admin",
//		//CacheHost:  "127.0.0.1:6379",
//		Style:        "goZero",
//		TemplatePath: fmt.Sprintf("%s/libs/template", RootPkgPath),
//		Database:     "",
//		DdlArg:       DdlArg{
//			Src:           fmt.Sprintf("%s/deploy/init.sql", RootPkgPath),
//			Cache:         true,
//			Strict:        false,
//			IgnoreColumns: []string{"create_at", "created_at", "create_time", "update_at", "updated_at", "update_time"},
//		},
//	}
//	return func(w http.ResponseWriter, r *http.Request){
//		pathx.RegisterGoctlHome(info.TemplatePath)
//		//生成数据库curl代码
//		err := BuildDataModel(&info)
//		if err != nil {
//			httpx.OkJson(w, xerr.NewErrCode(xerr.USER_OPERATION_ERR))
//			return
//		}
//
//		//生成api文件
//		dataModelToApi.DataModelToApi{
//			RootPkgPath,
//			dataModelToApi.ServiceInfo{
//				Title:   info.Title,
//				Desc:    info.Desc,
//				Author:  info.Author,
//				Email:   info.Email,
//				Version: info.Version,
//
//				ProjectName: info.ProjectName,
//				ServiceType: info.ServiceType,
//				Host:        info.Host,
//				Port:        info.Port,
//				DataSource: info.DataSource,
//				CacheHost: info.CacheHost,
//			}, dataModelToApi.SqlParseCfg{
//				Filename: info.DdlArg.Src,
//				Database: info.Database,
//				Strict: info.Strict,
//			},
//		}.StartBuild()
//
//		//生成api服务代码
//		err = BuildApiService(&info)
//		if err != nil {
//			httpx.OkJson(w, xerr.NewErrCode(xerr.USER_OPERATION_ERR))
//			return
//		}
//
//		//生成swagger doc文件
//		err = BuildSwaggerDoc(&info)
//		if err != nil {
//			httpx.OkJson(w, xerr.NewErrCode(xerr.USER_OPERATION_ERR))
//			return
//		}
//		result.HttpResult(r, w, nil, nil)
//	}
//}
