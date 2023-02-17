package dataModelToApi

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strings"
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
	//循环表的列表
	for _, table := range m.TableList {
		//加载html部分的模版
		htmlText, err := LoadTemplate(fmt.Sprintf("%s/common/dataModelToApi/template/web/html.tpl", m.RootPkgPath))
		if err != nil {
			return err
		}

		primaryKey := table.PrimaryKey.Name.Source()
		primaryKeyTpl := "<el-table-column prop=\"%s\" label=\"主键\" min-width=\"60\" :show-overflow-tooltip=\" true\" align=\"center\"/>"
		primaryKeyTextOutPut := fmt.Sprintf(primaryKeyTpl, primaryKey)
	    viewTextOutPuts := []string{}
		dialogTextOutPuts := []string{}

		for _, item := range table.Fields {
			showOverflowTooltip := ""
			minWidth := "80"
			if item.DataType == "string"{
				showOverflowTooltip = ":show-overflow-tooltip=\" true\""
				minWidth = "100"
			}

			if !stringx.Contains(ignoreColumns, item.Name.Source()) {
				if !stringx.Contains(updateColumns, item.Name.Source()) {
					//列表显示-所属用户
					if stringx.Contains([]string{"create_user", "user_id"}, item.Name.Source()){
						createUserFieldTpl := "<el-table-column\n          column-key=\"%s\"\n          :filters=\"nameList\"\n          :filter-method=\"filterHandler\"\n          prop=\"%s\" label=\"%s\" min-width=\"%s\" align=\"center\" %s/>"
						viewTextOutPuts = append(viewTextOutPuts, fmt.Sprintf(
							createUserFieldTpl,
							item.Name.Source(),
							item.Name.Source(),
							item.Comment,
							minWidth,
							showOverflowTooltip ))
					} else {
						//列表显示-其他字段
						fieldTpl := "      <el-table-column prop=\"%s\" label=\"%s\" min-width=\"%s\" align=\"center\" %s/>"
						if item.DataType == "int64"{
							fieldTpl = "      <el-table-column prop=\"%s\" label=\"%s\" min-width=\"%s\" align=\"center\" %s/>"
						}
						viewTextOutPuts = append(viewTextOutPuts, fmt.Sprintf(
							fieldTpl,
							item.Name.Source(),
							item.Comment,
							minWidth,
							showOverflowTooltip ))
					}
					//弹窗-其他字段
					dialogFieldTpl := "        <el-form-item label=\"%s\" prop=\"%s\" :show-overflow-tooltip=\"true\">\n          <el-input v-model=\"formData.%s\" autocomplete=\"off\" />\n        </el-form-item>"
					if item.DataType == "int64"{
						dialogFieldTpl = "        <el-form-item label=\"%s\" prop=\"%s\" min-width=\"100\">\n          <el-input-number v-model=\"formData.%s\" autocomplete=\"off\" />\n        </el-form-item>"
					}
					dialogTextOutPuts = append(dialogTextOutPuts, fmt.Sprintf(
						dialogFieldTpl,
						item.Comment,
						item.Name.Source(),
						item.Name.Source()))
				}
				//state字段
				if item.Name.Source() == "state"{
					stateViewOutPut := "      <el-table-column label=\"状态\" min-width=\"40\" align=\"center\">\n        <template #default=\"scope\">\n          <el-switch\n              v-model=\"scope.row.state\"\n              inline-prompt\n              :active-value=\"1\"\n              :inactive-value=\"-1\"\n              @change=\"handleStateChange(scope.row, v)\"\n          />\n        </template>\n      </el-table-column>"
					viewTextOutPuts = append(viewTextOutPuts, stateViewOutPut)

					stateFieldOutPut := "        <el-form-item label=\"状态\" prop=\"state\">\n          <el-switch\n              v-model=\"formData.state\"\n              inline-prompt\n              :active-value=\"1\"\n              :inactive-value=\"-1\"\n          />\n        </el-form-item>"
					dialogTextOutPuts = append(dialogTextOutPuts, stateFieldOutPut)
				}
			} else{
				//弹窗-创建/更新时间
				createOrUpdateTimeFieldTpl := "        <el-form-item label=\"%s\" prop=\"%s\" v-if=\"dialogType === 'edit'\" >\n          <el-input v-model=\"formData.%s\" autocomplete=\"off\"/>\n        </el-form-item>"
				dialogTextOutPuts = append(dialogTextOutPuts, fmt.Sprintf(
					createOrUpdateTimeFieldTpl,
					item.Comment,
					item.Name.Source(),
					item.Name.Source()))
			}
		}

		//组织好相关数据，生成代码，存为字节流
		htmlOutput, err := m.buildCode(htmlText, map[string]string{
			"primaryKey": primaryKey,
			"primaryKeyTextOutPut": primaryKeyTextOutPut,
			"viewTextOutPut": strings.Join(viewTextOutPuts, "\n"),
			"dialogTextOutPut": strings.Join(dialogTextOutPuts, "\n"),
			"dialogTitle": "<h4 :id=\"titleId\" :class=\"titleClass\">{{ dialogTitle }}</h4>",
		})
		if err != nil {
			return err
		}

		//加载style部分的模版
		styleText, err := LoadTemplate(fmt.Sprintf("%s/common/dataModelToApi/template/web/style.tpl", m.RootPkgPath))
		if err != nil {
			return err
		}
		styleOutput, err := m.buildCode(styleText, map[string]string{})
		if err != nil {
			return err
		}

		//加载list.tpl(主要模版）
		listText, err := LoadTemplate(fmt.Sprintf("%s/common/dataModelToApi/template/web/list.tpl", m.RootPkgPath))
		if err != nil {
			return err
		}
		//构建vue文件的存储路径
		listOutput, err := m.buildCode(listText, map[string]string{
			"ModuleName":table.Name.ToCamel(),
			"PrimaryKey":primaryKey,
			"Html":htmlOutput.String(),
			"Style":styleOutput.String(),
		})
		if err != nil {
			return err
		}

		//构建vue文件的存储路径
		outDir := path.Join(fmt.Sprintf("%s/libs/gin-vue-admin/web/src/view/demo", m.RootPkgPath), m.ServiceCfg.ProjectName)
		//outDir := path.Join(fmt.Sprintf("%s/projectBuilds/web", m.RootPkgPath), m.ServiceCfg.ProjectName)
		if err = os.MkdirAll(outDir, os.ModePerm); err != nil {
			fmt.Printf("MkdirAll outDir: %s", err.Error())
		}
		//构建文件输出路径
		outFile := path.Join(outDir, fmt.Sprintf("%sList.vue", table.Name.ToCamel() ) )
		//生成vue文件
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

//table.Name 表名
//table.PrimaryKey 主键 row-key="id"
//table.PrimaryKey.DataType 字段数据类型
//table.UniqueIndex 联合索引 暂时没用到
//table.Fields 表的所有字段
//item.DataType 字段类型