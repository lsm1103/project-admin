package swaggerDoc

import (
	"context"
	"io/ioutil"
	"os"

	"project-admin/test2/internal/svc"

	"github.com/bitly/go-simplejson"
	"github.com/zeromicro/go-zero/core/logx"
)

type DocDataLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDocDataLogic(ctx context.Context, svcCtx *svc.ServiceContext) DocDataLogic {
	return DocDataLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

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

func (l *DocDataLogic) DocData() (interface{}, error) {
	//path, _ := os.Getwd()
	buf, err := Read2Byte("/Users/xm/Desktop/go_package/project-admin/projectBuilds/project3/swagger.json")
	if err != nil {
		return "获取数据失败", err
	}
	return simplejson.NewJson(buf)
}
