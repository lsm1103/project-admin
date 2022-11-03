package main

import (
	"bytes"
	"fmt"
	"github.com/zeromicro/go-zero/core/errorx"
	"github.com/zeromicro/go-zero/core/stringx"
	"github.com/zeromicro/go-zero/tools/goctl/util"
	"io/ioutil"
	"os"
	"project-admin/adminManage/parser"
	"strings"
	"text/template"
)

type (
	adminManage struct {
		Syntax string
		Info string
		Types string
		ModuleHandler string
	}

	ServiceInfo struct {
		Title string
		Desc string
		Author string
		Email string
		Version string
		Host string
		Port string
		CommonPkgPath string
		DataSource string
		CacheHost string
	}
)

var ignoreColumns = []string{"create_at", "created_at", "create_time", "update_at", "updated_at", "update_time"}
var updateColumns = []string{"id", "state",}

func main() {
	//info模版
	serviceInfo := ServiceInfo{
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
	}
	infoText,err := LoadTemplate("template/info.tpl")
	if err != nil {
		panic(err)
	}
	infoOutput, err := build(infoText, serviceInfo)
	if err != nil {
		panic(err)
	}

	//解析数据库sql文件
	sqlFile := "/Users/xm/Desktop/go_package/project-admin/deploy/init.sql"
	database := ""
	strict := false
	tables, err := parser.Parse(sqlFile, database, strict)

	typesOutputs := []string{}
	moduleHandlerOutputs := []string{}
	//循环表的列表
	for _,item := range tables{
		//types模版
		fieldText,err := LoadTemplate("template/field.tpl")
		if err != nil {
			panic(err)
		}
		allFields, err := genFields(fieldText, item.Fields, "")
		if err != nil {
			panic(err)
		}

		createList := []*parser.Field{}
		updateList := []*parser.Field{}
		for _,item := range item.Fields{
			if !stringx.Contains(ignoreColumns, item.Name.Source()){
				updateList = append(updateList, item)
				if !stringx.Contains(updateColumns, item.Name.Source()){
					createList = append(createList, item)
				}
			}
		}
		//createFields
		createFields, err := genFields(fieldText, createList, "")
		if err != nil {
			panic(err)
		}
		//updateFields
		updateFields, err := genFields(fieldText, updateList, ",optional")
		if err != nil {
			panic(err)
		}

		typesData := map[string]string{
			"tableName":item.Name.ToCamel(),
			"createFields":createFields,
			"updateFields":updateFields,
			"allFields":allFields,
		}
		typesText, err := LoadTemplate("template/types.tpl")
		if err != nil {
			panic(err)
		}
		typesOutput, err := build(typesText, typesData)
		if err != nil {
			panic(err)
		}
		typesOutputs = append(typesOutputs, typesOutput.String())

		//module_handler 模版
		moduleHandlerText, err := LoadTemplate("template/module_handler.tpl")
		if err != nil {
			panic(err)
		}
		moduleHandlerOutput, err := build(moduleHandlerText, map[string]string{
			"tableName":item.Name.ToCamel(),
		})
		if err != nil {
			panic(err)
		}
		moduleHandlerOutputs = append(moduleHandlerOutputs, moduleHandlerOutput.String())
	}
	//adminManage根模版
	data := adminManage{
		Syntax:         "v1",
		Info:           infoOutput.String(),
		Types:          strings.Join(typesOutputs, "\n"),
		ModuleHandler: strings.Join(moduleHandlerOutputs, "\n"),
	}
	text,err := LoadTemplate("template/adminManage.tpl")
	if err != nil {
		panic(err)
	}
	output, err := build(text, data)
	if err != nil {
		panic(err)
	}
	//fmt.Printf("output:%s", output.String())

	//为每个表创建api文件
	err = ioutil.WriteFile("service2.api", output.Bytes(), os.ModePerm)
	if err != nil {
		panic(err)
	}
	fmt.Printf("sqlFile:%s is build done\n", sqlFile)
}

func genFields(fieldTemplate string, fields []*parser.Field, tag string) (string, error) {
	var list []string

	for _, field := range fields {
		fieldData := map[string]interface{}{
			"name":       util.SafeString(field.Name.ToCamel()),
			//"tag":        fmt.Sprintf(fieldTpl,field.Name.Source(), ""),
			"hasComment": field.Comment != "",
			"comment":    field.Comment,
		}
		if field.Name.Source() != "id"{
			fieldData["tag"] = fmt.Sprintf("`json:\"%s%s\"`",field.Name.Source(), tag)
		} else {
			fieldData["tag"] = fmt.Sprintf("`json:\"%s\"`",field.Name.Source() )
		}
		if field.DataType == "time.Time"{
			fieldData["type"] = "string"
		} else{
			fieldData["type"] = field.DataType
		}
		result, err := build(fieldTemplate, fieldData)
		if err != nil {
			return "", err
		}

		list = append(list, result.String())
	}
	return strings.Join(list, "\n"), nil
}

// LoadTemplate get template content by the filePath.
func LoadTemplate(filePath string) (string, error) {
	content, err := os.ReadFile(filePath)
	if err != nil {
		return "", err
	}
	return string(content), nil
}

func build(strTemplate string, data interface{}) (*bytes.Buffer, error) {
	tpl, err := template.New("templateOne").Parse(strTemplate)                                                               // （2）解析模板
	if err != nil {
		panic(err)
	}
	//err = tpl.Execute(os.Stdout, data) //（3）数据驱动模板，将name的值填充到模板中
	//if err != nil {
	//	panic(err)
	//}

	buf := new(bytes.Buffer)
	if err = tpl.Execute(buf, data); err != nil {
		return nil, errorx.Wrapf(err, "template execute error:%s", tpl)
	}

	//formatOutput, err := goformat.Source(buf.Bytes())
	//if err != nil {
	//	return nil, errorx.Wrapf(err, "go format error:%s", tpl)
	//}
	//buf.Reset()
	//buf.Write(formatOutput)
	return buf, nil
}
