package dataModelToApi

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"text/template"

	"github.com/zeromicro/go-zero/core/errorx"
	"github.com/zeromicro/go-zero/core/stringx"

	"project-admin/common/dataModelToApi/parser"
)

type (
	BuildWeb struct {
		RootPkgPath string
		TableList []*parser.Table
		ServiceCfg ServiceInfo
	}
)

func (m BuildWeb)BuildWebCode() error {
	//typesOutputs := []string{}
	//moduleHandlerOutputs := []string{}
	//循环表的列表
	for _, table := range m.TableList {
		//加载html部分的模版
		//组织好相关数据，生成代码，存为字节流
		htmlText, err := LoadTemplate(fmt.Sprintf("%s/common/dataModelToApi/template/web/html.tpl", m.RootPkgPath))
		if err != nil {
			return err
		}

		//所有字段
		//allFields := item.Fields
		//新增接口相关字段
		createList := []*parser.Field{}
		//修改接口相关字段
		updateList := []*parser.Field{}
		for _, item := range table.Fields {
			if !stringx.Contains(ignoreColumns, item.Name.Source()) {
				updateList = append(updateList, item)
				//编辑页面字段

				//table.Name 表名
				//table.PrimaryKey 主键
				//table.PrimaryKey.DataType 字段数据类型
				//table.UniqueIndex 联合索引
				//table.Fields 表的所有字段
				//item.DataType 字段类型

				if !stringx.Contains(updateColumns, item.Name.Source()) {
					createList = append(createList, item)
					//列表显示字段、新增弹窗字段
					//textTpl := "<el-table-column prop=\"project_id\" label=\"所属项目id\" min-width=\"80\" align=\"center\"/>"
				}
			}
		}

		//for _,item := range createList{
		//	item.Name
		//}
		//
		//addField := "<el-table-column prop=\"en_name\" label=\"英文名称\" min-width=\"80\" align=\"center\" />"
		//
		//updateField := "<el-form-item label=\"英文名称\" prop=\"en_name\">\n  <el-input v-model=\"formData.en_name\" autocomplete=\"off\" />\n</el-form-item>"

		htmlOutput, err := m.buildCode(htmlText, map[string]string{})
		if err != nil {
			return err
		}

		//加载style部分的模版
		//组织好相关数据，生成代码，存为字节流
		styleText, err := LoadTemplate(fmt.Sprintf("%s/common/dataModelToApi/template/web/style.tpl", m.RootPkgPath))
		if err != nil {
			return err
		}
		styleOutput, err := m.buildCode(styleText, map[string]string{

		})
		if err != nil {
			return err
		}

		//加载list.tpl(主要模版）
		//组织好相关数据，生成代码，存为字节流
		listText, err := LoadTemplate(fmt.Sprintf("%s/common/dataModelToApi/template/web/list.tpl", m.RootPkgPath))
		if err != nil {
			return err
		}
		listOutput, err := m.buildCode(listText, map[string]string{
			"H  tml":htmlOutput.String(),
			"St yle":styleOutput.String(),
		})
		if err != nil {
			return err
		}

		//把三部分拼接到一起
		//构建vue文件的存储路径
		//生成vue文件
		//构建文件输出路径
		outDir := path.Join(fmt.Sprintf("%s/projectBuilds/web", m.RootPkgPath), m.ServiceCfg.ProjectName)
		if err = os.MkdirAll(outDir, os.ModePerm); err != nil {
			fmt.Printf("MkdirAll outDir: %s", err.Error())
		}
		outFile := path.Join(outDir, fmt.Sprintf("%sList.vue", table.Name.ToCamel() ) )
		//创建api文件
		err = ioutil.WriteFile(outFile, listOutput.Bytes(), os.ModePerm)
		if err != nil {
			return err
		}
		fmt.Printf("project dir:%s is build done\n", outFile)

	}
	return nil
}

func (m BuildWeb)buildCode(strTemplate string, data interface{}) (*bytes.Buffer, error) {
	tpl, err := template.New("templateOne").Parse(strTemplate) // （2）解析模板
	if err != nil {
		return nil, err
	}

	buf := new(bytes.Buffer)
	if err = tpl.Execute(buf, data); err != nil {
		return nil, errorx.Wrapf(err, "template execute error:%+v", tpl)
	}
	return buf, nil
}