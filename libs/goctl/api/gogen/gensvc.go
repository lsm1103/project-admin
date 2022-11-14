package gogen

import (
	"bytes"
	_ "embed"
	"fmt"
	"strings"
	"text/template" //add

	"github.com/zeromicro/go-zero/tools/goctl/api/spec"
	"github.com/zeromicro/go-zero/tools/goctl/config"
	"github.com/zeromicro/go-zero/tools/goctl/util/format"
	"github.com/zeromicro/go-zero/tools/goctl/util/pathx"
	"github.com/zeromicro/go-zero/tools/goctl/vars"
)

const contextFilename = "service_context"

//go:embed svc.tpl
var contextTemplate string

func genServiceContext(dir, rootPkg string, cfg *config.Config, api *spec.ApiSpec) error {
	filename, err := format.FileNamingFormat(cfg.NamingFormat, contextFilename)
	if err != nil {
		return err
	}

	var middlewareStr string
	var middlewareAssignment string
	middlewares := getMiddleware(api)

	for _, item := range middlewares {
		middlewareStr += fmt.Sprintf("%s rest.Middleware\n", item)
		name := strings.TrimSuffix(item, "Middleware") + "Middleware"
		middlewareAssignment += fmt.Sprintf("%s: %s,\n", item,
			fmt.Sprintf("middleware.New%s().%s", strings.Title(name), "Handle"))
	}

	configImport := "\"" + pathx.JoinPackages(rootPkg, configDir) + "\""
	if len(middlewareStr) > 0 {
		configImport += "\n\t\"" + pathx.JoinPackages(rootPkg, middlewareDir) + "\""
		configImport += fmt.Sprintf("\n\t\"%s/rest\"", vars.ProjectOpenSourceURL)
	}
	// ==========add==========
	var typeFieldList []string
	var valueFieldList []string
	for _, item := range api.Service.Groups {
		moduleName := item.Annotation.Properties["module"]
		typeFieldTpl := "{{.moduleName}}Model dataModel.{{.moduleName}}Model"
		t := template.Must(template.New("typeFieldTpl").Parse(typeFieldTpl))
		buffer := new(bytes.Buffer)
		err = t.Execute(buffer, map[string]string{"moduleName": moduleName})
		if err != nil {
			return err
		}
		typeFieldList = append(typeFieldList, buffer.String())

		valueFieldTpl := "{{.moduleName}}Model: dataModel.New{{.moduleName}}Model(sqlx.NewMysql(c.DB.DataSource ), c.Cache),"
		t_ := template.Must(template.New("valueFieldTpl").Parse(valueFieldTpl))
		buffer_ := new(bytes.Buffer)
		err = t_.Execute(buffer_, map[string]string{"moduleName": moduleName})
		if err != nil {
			return err
		}
		valueFieldList = append(valueFieldList, buffer_.String())
	}
	typeFields := strings.Join(typeFieldList, "\n")
	valueFields := strings.Join(valueFieldList, "\n")

	rootPkgName := rootPkg
	if strings.Contains(rootPkg, "/") {
		rootPkgName = strings.Split(rootPkg, "/")[0]
	}

	dirList := strings.Split(dir, "/")
	projectName := dirList[len(dirList)-1]
	// ==========add==========
	return genFile(fileGenConfig{
		dir:             dir,
		subdir:          contextDir,
		filename:        filename + ".go",
		templateName:    "contextTemplate",
		category:        category,
		templateFile:    contextTemplateFile,
		builtinTemplate: contextTemplate,
		data: map[string]string{
			"configImport":         configImport,
			"rootPkgName":          rootPkgName,
			"projectName":          projectName,
			"config":               "config.Config",
			"middleware":           middlewareStr,
			"middlewareAssignment": middlewareAssignment,
			"typeFields":           typeFields,
			"valueFields":          valueFields,
		},
	})
}
