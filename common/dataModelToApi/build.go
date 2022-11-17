package dataModelToApi

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"project-admin/common/listUtils"
	"strings"
	"text/template"

	"github.com/zeromicro/go-zero/core/errorx"
	"github.com/zeromicro/go-zero/core/stringx"
	"github.com/zeromicro/go-zero/tools/goctl/util"
	//"github.com/zeromicro/go-zero/tools/goctl/model/sql/parser"
	"project-admin/common/dataModelToApi/parser"
)

type (
	ServiceInfo struct {
		Title   string
		Desc    string
		Author  string
		Email   string
		Version string

		ProjectName string
		ServiceType string	`json:"service_type,options=admin|mock"`
		Host        string
		Port        string
		DataSource  string
		CacheHost   string
	}
	SqlParseCfg struct {
		Filename string
		Database string
		Strict   bool
	}
	DataModelToApi struct {
		ServiceCfg ServiceInfo
		SqlCfg SqlParseCfg
	}

	adminManage struct {
		Syntax        string
		Info          string
		Types         string
		ModuleHandler string
	}
)

var ignoreColumns = []string{"create_at", "created_at", "create_time", "update_at", "updated_at", "update_time"}
var updateColumns = []string{"id", "state"}

func (m DataModelToApi)StartBuild() {
	//info模版
	infoText, err := LoadTemplate("template/info.tpl")
	if err != nil {
		panic(err)
	}
	infoOutput, err := m.buildCode(infoText, m.ServiceCfg)
	if err != nil {
		panic(err)
	}

	//解析数据库sql文件
	tables, err := parser.Parse(m.SqlCfg.Filename, m.SqlCfg.Database, m.SqlCfg.Strict)
	if err != nil {
		panic(err)
	}
	typesOutputs := []string{}
	moduleHandlerOutputs := []string{}
	//循环表的列表
	for _, item := range tables {
		//types模版
		fieldText, err := LoadTemplate("template/field.tpl")
		if err != nil {
			panic(err)
		}
		allFields, err := m.genFields(fieldText, item.Fields, "", "all")
		if err != nil {
			panic(err)
		}

		createList := []*parser.Field{}
		updateList := []*parser.Field{}
		for _, item := range item.Fields {
			if !stringx.Contains(ignoreColumns, item.Name.Source()) {
				updateList = append(updateList, item)
				if !stringx.Contains(updateColumns, item.Name.Source()) {
					createList = append(createList, item)
				}
			}
		}
		//createFields
		createFields, err := m.genFields(fieldText, createList, "", "create")
		if err != nil {
			panic(err)
		}
		//updateFields
		updateFields, err := m.genFields(fieldText, updateList, ",optional", "update")
		if err != nil {
			panic(err)
		}

		typesData := map[string]string{
			"tableName":    item.Name.ToCamel(),
			"createFields": createFields,
			"updateFields": updateFields,
			"allFields":    allFields,
		}
		typesText, err := LoadTemplate("template/types.tpl")
		if err != nil {
			panic(err)
		}
		typesOutput, err := m.buildCode(typesText, typesData)
		if err != nil {
			panic(err)
		}
		typesOutputs = append(typesOutputs, typesOutput.String())

		//module_handler 模版
		moduleHandlerText, err := LoadTemplate("template/module_handler.tpl")
		if err != nil {
			panic(err)
		}
		moduleHandlerOutput, err := m.buildCode(moduleHandlerText, map[string]string{
			"tableName": item.Name.ToCamel(),
			"serviceType": m.ServiceCfg.ServiceType,
		})
		if err != nil {
			panic(err)
		}
		moduleHandlerOutputs = append(moduleHandlerOutputs, moduleHandlerOutput.String())
	}
	//adminManage根模版
	data := adminManage{
		Syntax:        "v1",
		Info:          infoOutput.String(),
		Types:         strings.Join(typesOutputs, "\n"),
		ModuleHandler: strings.Join(moduleHandlerOutputs, "\n"),
	}
	text, err := LoadTemplate("template/adminManage.tpl")
	if err != nil {
		panic(err)
	}
	output, err := m.buildCode(text, data)
	if err != nil {
		panic(err)
	}

	outDir := path.Join("../../projectBuilds", m.ServiceCfg.ProjectName)
	if err = os.MkdirAll(outDir, os.ModePerm); err != nil {
		fmt.Printf("MkdirAll outDir: %s", err.Error())
	}
	outFile := path.Join(outDir, "service.api")

	//为每个表创建api文件
	err = ioutil.WriteFile(outFile, output.Bytes(), os.ModePerm)
	if err != nil {
		panic(err)
	}
	fmt.Printf("project dir:%s is build done\n", outFile)
}

//生成字段列表
func (m DataModelToApi)genFields(fieldTemplate string, fields []*parser.Field, tag, type_ string) (string, error) {
	var list []string

	for _, field := range fields {
		fieldData := map[string]interface{}{
			"name": util.SafeString(field.Name.ToCamel()),
			//"tag":        fmt.Sprintf(fieldTpl,field.Name.Source(), ""),
			"hasComment": field.Comment != "",
			"comment":    field.Comment,
		}
		fieldName := field.Name.Source()
		if fieldName == "id" {
			if type_ != "all" || m.ServiceCfg.ServiceType != "mock"{
				fieldData["tag"] = fmt.Sprintf("`json:\"%s\"`", fieldName)
			} else {
				fieldData["tag"] = fmt.Sprintf("`json:\"%s\"tag:\"uuid\"`", fieldName)
			}
		} else {
			if type_ != "all" || m.ServiceCfg.ServiceType != "mock"{
				fieldData["tag"] = fmt.Sprintf("`json:\"%s%s\"`", fieldName, tag)
			} else {
				MockTagMap := map[string]string{
					"int64": "tag:\"uint\"min:\"1\"max:\"20\"",
					"string": "tag:\"chineseChar\"fixed_len:\"18|150\"",
					"bool":"tag:\"bool\"",
				}
				MockTagList := []string{
					"phone",
					"email",
					"name",
					"address",
					"bankID",
					"city",
					"idCart",
					"english",
					"orderNo",
					"password",
					"time",
					"timeStamp",
					"date",
					"dateTime",
				}
				TimeMockTag := map[string]string{
					"create_time":"timeStamp",
					"update_time":"timeStamp",
				}
				tag_ := MockTagMap[field.DataType]
				if listUtils.In(fieldName, MockTagList){
					tag_ = fmt.Sprintf("tag:\"%s\"",fieldName)
				}
				if t_,ok := TimeMockTag[fieldName]; ok{
					tag_ = fmt.Sprintf("tag:\"%s\"", t_)
				}
				fieldData["tag"] = fmt.Sprintf("`json:\"%s%s\"%s`", fieldName, tag, tag_)
			}
		}
		if field.DataType == "time.Time" {
			fieldData["type"] = "string"
		} else {
			fieldData["type"] = field.DataType
		}
		result, err := m.buildCode(fieldTemplate, fieldData)
		if err != nil {
			return "", err
		}

		list = append(list, result.String())
	}
	return strings.Join(list, "\n"), nil
}

func (m DataModelToApi)buildCode(strTemplate string, data interface{}) (*bytes.Buffer, error) {
	tpl, err := template.New("templateOne").Parse(strTemplate) // （2）解析模板
	if err != nil {
		panic(err)
	}

	buf := new(bytes.Buffer)
	if err = tpl.Execute(buf, data); err != nil {
		return nil, errorx.Wrapf(err, "template execute error:%+v", tpl)
	}
	return buf, nil
}


// LoadTemplate get template content by the filePath.
func LoadTemplate(filePath string) (string, error) {
	content, err := os.ReadFile(filePath)
	if err != nil {
		return "", err
	}
	return string(content), nil
}
