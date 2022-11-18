package main

import (
	"fmt"
	"github.com/bitly/go-simplejson"
	"github.com/zeromicro/go-zero/rest/httpx"
	"github.com/zeromicro/go-zero/tools/goctl/config"
	"github.com/zeromicro/go-zero/tools/goctl/plugin"
	"github.com/zeromicro/go-zero/tools/goctl/util/console"
	"golang.org/x/xerrors"
	"io/ioutil"
	"net/http"
	"os"
	"project-admin/common/result"
	"project-admin/common/xerr"
	"strings"

	"project-admin/common/dataModelToApi"
	gogen "project-admin/libs/goctl_/api-gogen"
	parser "project-admin/libs/goctl_/api-parser"
	"project-admin/libs/goctl_/goctl-swagger-generate"
	gen "project-admin/libs/goctl_/sql-gen"
	util "project-admin/libs/goctl_/sql-util"
	pathx "project-admin/libs/goctl_/util-pathx"
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


type DdlArg struct {
	Src 	string
	Cache 	bool
	Strict 	bool
	IgnoreColumns 	[]string
}
type BuildAppInfo struct {
	Title       string `json:"title,content=项目标题"`
	Desc        string `json:"desc,content=项目说明"`
	Author      string `json:"author,content=项目作者"`
	Email       string `json:"email,content=联系邮箱"`
	Version     string `json:"version,content=版本号"`

	ProjectName string `json:"projectName,content=项目英文名称"`
	ServiceType string `json:"service_type,options=admin|mock,content=项目生成类型"`
	Host        string `json:"host,content=域名"`
	Port        string `json:"port,content=端口"`
	DataSource  string `json:"dataSource,content=数据源"`
	CacheHost   string `json:"cacheHost,content=缓存域名"`
	Style 		string `json:"style,content=项目代码风格"`
	TemplatePath string `json:"templatePath,content=模版地址"`

	Database 	string `json:"database,content=数据库名"`
	DdlArg  		   `json:"ddlArg,content=生成数据库curl代码配置"`
}
var RootPkgPath = "/Users/xm/Desktop/go_package/project-admin"

//生成数据库curl代码
func BuildDataModel(info *BuildAppInfo) error {
	log := console.NewConsole(true)
	src := strings.TrimSpace(info.Src)
	if len(src) == 0 {
		return xerrors.New("expected path or path globbing patterns, but nothing found")

	}
	files, err := util.MatchFiles(src)
	if err != nil {
		return err
	}
	if len(files) == 0 {
		return xerrors.New("sql not matched")
	}

	generator, err := gen.NewDefaultGenerator(
		fmt.Sprintf("%s/dataModel/%s", RootPkgPath, info.ProjectName),
		&config.Config{NamingFormat: info.Style },
		gen.WithConsoleOption(log), gen.WithIgnoreColumns(info.IgnoreColumns))
	if err != nil {
		return err
	}

	for _, file := range files {
		err = generator.StartFromDDL(file, info.Cache, info.Strict, info.Database)
		if err != nil {
			return err
		}
	}
	return nil
}

//生成api服务代码
func BuildApiService(info *BuildAppInfo) error {
	apiFile := fmt.Sprintf("%s/projectBuilds/%s/service.api", RootPkgPath, info.ProjectName)
	dir := fmt.Sprintf("%s/projectBuilds/%s", RootPkgPath, info.ProjectName)
	pathx.RegisterGoctlHome(info.TemplatePath)
	return gogen.DoGenProject(apiFile, dir, info.Style)
}

func BuildSwaggerDoc(info *BuildAppInfo) error {
	apiFilePath := fmt.Sprintf("%s/projectBuilds/%s/service.api", RootPkgPath, info.ProjectName)
	sp, err := parser.Parse(apiFilePath)
	if err != nil {
		fmt.Printf("err:%+v", err)
	}
	url := fmt.Sprintf("%s:%s", info.Host, info.Port)
	return generate.Do("swagger.json", url, "", &plugin.Plugin{
		Api:         sp,
		ApiFilePath: apiFilePath,
		Style:       info.Style,
		Dir:         fmt.Sprintf("%s/projectBuilds/%s", RootPkgPath, info.ProjectName),
	})
}

func BuildAPP() http.HandlerFunc {
	info := BuildAppInfo{
		Title:   "项目管理服务",
		Desc:    "对研发项目进行管理，包括代码生成、mock服务生成、cicd等；如通过api设计文档自动生成服务，并根据api文件配置的字段mock规则进行mock生成结果",
		Author:  "lsm",
		Email:   "18370872400@163.com",
		Version: "v0.1.1",

		ProjectName: "project10",
		ServiceType: "admin",
		Host:        "0.0.0.0",
		Port:        "820",
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
	}
	return func(w http.ResponseWriter, r *http.Request){
		pathx.RegisterGoctlHome(info.TemplatePath)
		//生成数据库curl代码
		err := BuildDataModel(&info)
		if err != nil {
			httpx.OkJson(w, xerr.NewErrCode(xerr.USER_OPERATION_ERR))
			return
		}

		//生成api文件
		dataModelToApi.DataModelToApi{
			RootPkgPath,
			dataModelToApi.ServiceInfo{
				Title:   info.Title,
				Desc:    info.Desc,
				Author:  info.Author,
				Email:   info.Email,
				Version: info.Version,

				ProjectName: info.ProjectName,
				ServiceType: info.ServiceType,
				Host:        info.Host,
				Port:        info.Port,
				DataSource: info.DataSource,
				CacheHost: info.CacheHost,
			}, dataModelToApi.SqlParseCfg{
				Filename: info.DdlArg.Src,
				Database: info.Database,
				Strict: info.Strict,
			},
		}.StartBuild()

		//生成api服务代码
		err = BuildApiService(&info)
		if err != nil {
			httpx.OkJson(w, xerr.NewErrCode(xerr.USER_OPERATION_ERR))
			return
		}

		//生成swagger doc文件
		err = BuildSwaggerDoc(&info)
		if err != nil {
			httpx.OkJson(w, xerr.NewErrCode(xerr.USER_OPERATION_ERR))
			return
		}
		result.HttpResult(r, w, nil, nil)
	}
}