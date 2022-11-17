package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"github.com/bitly/go-simplejson"
	"github.com/zeromicro/go-zero/rest/httpx"
	"github.com/zeromicro/go-zero/tools/goctl/plugin"

	"project-admin/common/dataModelToApi"
	"project-admin/libs/goctl/model/sql/command"
	"project-admin/libs/goctl-swagger/generate"
	"project-admin/libs/goctl/api/gogen"
	"project-admin/libs/goctl/api/parser"
	"project-admin/libs/goctl/util/pathx"
)

func Read2Byte(path string) ([]byte, error) {
	f, err := os.Open(path)
	if err != nil {
		return []byte("read file fail"), err
	}
	defer f.Close()
	fd, err := ioutil.ReadAll(f)
	if err != nil {
		return []byte("read file fail"), err
	}
	return fd, nil
}

const swaggerTemplateV2 = `
<!-- HTML for static distribution bundle build -->
<!DOCTYPE html>
<html lang="en">
  <head>
    <meta charset="UTF-8">
    <title>API documentation</title>
    <link rel="stylesheet" type="text/css" href="{{ .SwaggerHost }}/swagger-ui.css" >
    <link rel="icon" type="image/png" href="{{ .SwaggerHost }}/favicon-32x32.png" sizes="32x32" />
    <link rel="icon" type="image/png" href="{{ .SwaggerHost }}/favicon-16x16.png" sizes="16x16" />
    <style>
      html
      {
        box-sizing: border-box;
        overflow: -moz-scrollbars-vertical;
        overflow-y: scroll;
      }

      *,
      *:before,
      *:after
      {
        box-sizing: inherit;
      }

      body
      {
        margin:0;
        background: #fafafa;
      }
    </style>
  </head>

  <body>
    <div id="swagger-ui"></div>

    <script src="{{ .SwaggerHost }}/swagger-ui-bundle.js"> </script>
    <script src="{{ .SwaggerHost }}/swagger-ui-standalone-preset.js"> </script>
    <script>
    window.onload = function() {
      // Begin Swagger UI call region
      const ui = SwaggerUIBundle({
        "dom_id": "#swagger-ui",
        deepLinking: true,
        presets: [
          SwaggerUIBundle.presets.apis,
          SwaggerUIStandalonePreset
        ],
        plugins: [
          SwaggerUIBundle.plugins.DownloadUrl
        ],
        layout: "StandaloneLayout",
		validatorUrl: null,
        url: "{{ .SpecURL }}",
      })

      // End Swagger UI call region
      window.ui = ui
    }
  </script>
  </body>
</html>`

const html = `<!DOCTYPE html>
<html>
	<head>
	<link type="text/css" rel="stylesheet" href="https://cdn.jsdelivr.net/npm/swagger-ui-dist@3/swagger-ui.css">
	<!-- <link rel="shortcut icon" href="https://fastapi.tiangolo.com/img/favicon.png"> -->
	<title>Swagger UI</title>
	</head>
	<body>
	<div id="swagger-ui">
	</div>
	<script src="https://cdn.jsdelivr.net/npm/swagger-ui-dist@3/swagger-ui-bundle.js"></script>
	<!-- 'SwaggerUIBundle' is now available on the page -->
	<script>
	const ui = SwaggerUIBundle({
		url: '{{.swaggerDoc}}',
		dom_id: '#swagger-ui',
		presets: [
		SwaggerUIBundle.presets.apis,
		SwaggerUIBundle.SwaggerUIStandalonePreset
		],
		layout: "BaseLayout",
		deepLinking: true,
		showExtensions: true,
		showCommonExtensions: true
	})
	</script>
	</body>
</html>`

func Doc(swaggerDocUrl, env string) http.HandlerFunc {
	// permission
	needPermission := false
	if env == "prod" {
		needPermission = true
	}

	return func(w http.ResponseWriter, r *http.Request) {
		//w.Header().Set("Content-Type", "text/plain")
		r.ParseForm()
		projectName := r.Form.Get("name")
		url := r.Form.Get("url")
		if projectName == "" && url != ""{
			swaggerDocUrl = url
		}else if projectName != "" {
			swaggerDocUrl = fmt.Sprintf("/docData?name=%s",projectName)
		} else {
			swaggerDocUrl = "/docData"
		}
		html_ := strings.Replace(html, "{{.swaggerDoc}}", swaggerDocUrl, 1)

		if needPermission {
			//w.WriteHeader(http.StatusOK)
			w.Write([]byte("Swagger not open on prod"))
			return
		}
		fmt.Fprintf(w, html_)
	}
}

func DocData() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		projectName := r.Form.Get("name")
		path := "/Users/xm/Desktop/go_package/project-admin/project-admin/swagger.json"
		if projectName != ""{
			path = fmt.Sprintf("../projectBuilds/%s/swagger.json", projectName)
		}
		buf, err := Read2Byte(path)
		if err != nil {
			httpx.Error(w, err)
			return
		}
		json, err := simplejson.NewJson(buf)
		if err != nil {
			httpx.Error(w, err)
			return
		}
		httpx.OkJson(w, json)
	}
}

func BuildAPP()  {
	//生成数据库curl代码
	pathx.RegisterGoctlHome("/Users/xm/Desktop/go_package/project-admin/libs/template")
	err := command.FromDDL(command.DdlArg{
		Src:           "/Users/xm/Desktop/go_package/project-admin/deploy/init.sql",
		Dir:           "/Users/xm/Desktop/go_package/project-admin/dataModel/project4",
		Cache:         true,
		Strict:        false,
		IgnoreColumns: []string{"create_at", "created_at", "create_time", "update_at", "updated_at", "update_time"},
	})
	if err != nil {
		return
	}

	//生成api文件
	dataModelToApi.DataModelToApi{
		dataModelToApi.ServiceInfo{
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
		}, dataModelToApi.SqlParseCfg{
			Filename: "/Users/xm/Desktop/go_package/project-admin/deploy/init.sql",
			Database: "",
			Strict: false,
		},
	}.StartBuild()

	//生成api服务代码
	apiFile := "/Users/xm/Desktop/go_package/project-admin/projectBuilds/project5/service.api"
	dir := "/Users/xm/Desktop/go_package/project-admin/projectBuilds/project5"
	style := "goZero"
	pathx.RegisterGoctlHome("/Users/xm/Desktop/go_package/project-admin/libs/template")
	err = gogen.DoGenProject(apiFile, dir, style)

	//生成swagger doc文件
	fileName := "swagger.json"
	apiFilePath := "/Users/xm/Desktop/go_package/project-admin/projectBuilds/project10/service.api"
	sp, err := parser.Parse(apiFilePath)
	if err != nil {
		return
	}
	host := sp.Info.Properties["host"]
	port := sp.Info.Properties["port"]
	url := fmt.Sprintf("%s:%s", host, port)
	err = generate.Do(fileName, url, "", &plugin.Plugin{
		Api:         sp,
		ApiFilePath: apiFilePath,
		Style:       "goZero",
		Dir:         "/Users/xm/Desktop/go_package/project-admin/projectBuilds/project10",
	})
	if err != nil {
		return
	}
}