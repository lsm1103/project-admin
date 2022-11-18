package generate

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/zeromicro/go-zero/tools/goctl/api/parser"
	"github.com/zeromicro/go-zero/tools/goctl/plugin"
	"testing"
)

func TestDo(t *testing.T) {
	fileName := "swagger.json"
	apiFilePath := "/Users/xm/Desktop/go_package/project-admin/projectBuilds/project10/service.api"
	sp, err := parser.Parse(apiFilePath)
	assert.Nil(t, err)
	host := sp.Info.Properties["host"]
	port := sp.Info.Properties["port"]
	url := fmt.Sprintf("%s:%s", host, port)
	err = Do(fileName, url, "", &plugin.Plugin{
		Api:         sp,
		ApiFilePath: apiFilePath,
		Style:       "goZero",
		Dir:         "/Users/xm/Desktop/go_package/project-admin/projectBuilds/project10",
	})
	t.Logf("err:%+v", err)
}