package buildCode

import (
	"fmt"
	"github.com/zeromicro/go-zero/tools/goctl/config"
	"github.com/zeromicro/go-zero/tools/goctl/plugin"
	"github.com/zeromicro/go-zero/tools/goctl/util/console"
	"github.com/zeromicro/go-zero/tools/goctl/util/pathx"
	"golang.org/x/xerrors"
	"project-admin/common/dataModelToApi"
	"strings"

	gogen "project-admin/libs/goctl_/api-gogen"
	parser "project-admin/libs/goctl_/api-parser"
	"project-admin/libs/goctl_/goctl-swagger-generate"
	gen "project-admin/libs/goctl_/sql-gen"
	util "project-admin/libs/goctl_/sql-util"
	//pathx "project-admin/libs/goctl_/util-pathx"
)

type BuildCode struct {
	RootPkgPath string
	Info BuildAppInfo
}
type DdlArg struct {
	Src 	string
	Cache 	bool
	Strict 	bool
	IgnoreColumns 	[]string
}
type BuildAppInfo struct {
	Title       string `json:"title,content=应用标题"`
	Desc        string `json:"desc,content=应用说明"`
	Author      string `json:"author,content=应用作者"`
	Email       string `json:"email,content=联系邮箱"`
	Version     string `json:"version,content=版本号"`

	ProjectName string `json:"projectName,content=应用英文名称"`
	ServiceType string `json:"service_type,options=admin|mock,content=应用生成类型"`
	Host        string `json:"host,content=域名"`
	Port        string `json:"port,content=端口"`
	DataSource  string `json:"dataSource,content=数据源"`
	CacheHost   string `json:"cacheHost,content=缓存域名"`
	Style 		string `json:"style,content=项目代码风格"`
	TemplatePath string `json:"templatePath,content=模版地址"`

	Database 	string `json:"database,content=数据库名"`
	DdlArg  		   `json:"ddlArg,content=生成数据库curl代码配置"`
}

//生成api文件
func (m BuildCode)BuildApiFile() error {
	return dataModelToApi.DataModelToApi{
		m.RootPkgPath,
		dataModelToApi.ServiceInfo{
			Title:   m.Info.Title,
			Desc:    m.Info.Desc,
			Author:  m.Info.Author,
			Email:   m.Info.Email,
			Version: m.Info.Version,

			ProjectName: m.Info.ProjectName,
			ServiceType: m.Info.ServiceType,
			Host:        m.Info.Host,
			Port:        m.Info.Port,
			DataSource: m.Info.DataSource,
			CacheHost: m.Info.CacheHost,
		}, dataModelToApi.SqlParseCfg{
			Filename: m.Info.DdlArg.Src,
			Database: m.Info.Database,
			Strict: m.Info.Strict,
		},
	}.StartBuild()
}

//生成数据库curl代码
func (m BuildCode)BuildDataModel() error {
	log := console.NewConsole(true)
	src := strings.TrimSpace(m.Info.Src)
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

	pathx.RegisterGoctlHome(m.Info.TemplatePath)
	generator, err := gen.NewDefaultGenerator(
		fmt.Sprintf("%s/dataModel/%s", m.RootPkgPath, m.Info.ProjectName),
		&config.Config{NamingFormat: m.Info.Style },
		gen.WithConsoleOption(log), gen.WithIgnoreColumns(m.Info.IgnoreColumns))
	if err != nil {
		return err
	}

	for _, file := range files {
		err = generator.StartFromDDL(file, m.Info.Cache, m.Info.Strict, m.Info.Database)
		if err != nil {
			return err
		}
	}
	return nil
}

//生成api服务代码
func (m BuildCode)BuildApiService() error {
	apiFile := fmt.Sprintf("%s/projectBuilds/%s/service.api", m.RootPkgPath, m.Info.ProjectName)
	dir := fmt.Sprintf("%s/projectBuilds/%s", m.RootPkgPath, m.Info.ProjectName)
	pathx.RegisterGoctlHome(m.Info.TemplatePath)
	return gogen.DoGenProject(apiFile, dir, m.Info.Style)
}

//生成swagger doc文件
func (m BuildCode)BuildSwaggerDoc() error {
	apiFilePath := fmt.Sprintf("%s/projectBuilds/%s/service.api", m.RootPkgPath, m.Info.ProjectName)
	sp, err := parser.Parse(apiFilePath)
	if err != nil {
		fmt.Printf("err:%+v", err)
	}
	url := fmt.Sprintf("%s:%s", m.Info.Host, m.Info.Port)
	return generate.Do("swagger.json", url, "", &plugin.Plugin{
		Api:         sp,
		ApiFilePath: apiFilePath,
		Style:       m.Info.Style,
		Dir:         fmt.Sprintf("%s/projectBuilds/%s", m.RootPkgPath, m.Info.ProjectName),
	})
}