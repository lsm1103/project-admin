package Application

import (
	"context"
	"fmt"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"project-admin/common/buildCode"
	"project-admin/common/xerr"

	"github.com/zeromicro/go-zero/core/logx"
	"project-admin/project-admin/internal/svc"
	"project-admin/project-admin/internal/types"
)

type BuildLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewBuildLogic(ctx context.Context, svcCtx *svc.ServiceContext) BuildLogic {
	return BuildLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *BuildLogic) Build(req *types.BuildReq, id int64) error {
	//app := &types.Application{}
	//err := l.svcCtx.ApplicationModel.FindOne(l.ctx, nil, id, app)
	//if err != nil {
	//	return errors.Wrapf(xerr.NewErrCode(xerr.USER_OPERATION_ERR),"获取数据失败：%s", err.Error())
	//}
	//
	build := buildCode.BuildCode{
		RootPkgPath: l.svcCtx.RootPkgPath,
		//Info:        buildCode.BuildAppInfo{
		//	Title:        app.ZnName,
		//	Desc:         app.Info,
		//	Author:       "",
		//	Email:        "",
		//	Version:      "",
		//	ProjectName:  app.EnName,
		//	ServiceType:  "",
		//	Host:         "",
		//	Port:         "",
		//	DataSource:   "",
		//	CacheHost:    "",
		//	Style:        "",
		//	TemplatePath: "",
		//	Database:     "",
		//	DdlArg:       buildCode.DdlArg{},
		//},
	}
	err := copier.Copy(&build.Info, req)
	if err != nil {
		return errors.Wrapf(xerr.NewErrCode(xerr.USER_OPERATION_ERR),
			"数据格式转化失败：%s", err.Error())
	}

	if build.Info.TemplatePath == ""{
		build.Info.TemplatePath = fmt.Sprintf("%s/libs/template", build.RootPkgPath)
	}
	if build.Info.Src == ""{
		build.Info.Src = fmt.Sprintf("%s/deploy/init.sql", build.RootPkgPath)
	}
	if len(build.Info.IgnoreColumns) == 0{
		build.Info.IgnoreColumns = []string{"create_at", "created_at", "create_time", "update_at", "updated_at", "update_time"}
	}

	buildType := "buildAll"
	switch buildType {
	case "buildApiFile":
		//生成api文件
		err = build.BuildApiFile()
		if err != nil {
			return errors.Wrapf(xerr.NewErrCode(xerr.USER_OPERATION_ERR),
				"生成api文件失败：%s", err.Error())
		}
	case "buildDataModel":
		//生成数据库curl代码
		err = build.BuildDataModel()
		if err != nil {
			return errors.Wrapf(xerr.NewErrCode(xerr.USER_OPERATION_ERR),
				"生成数据库curl代码失败：%s", err.Error())
		}
	case "buildApiService":
		//生成api服务代码
		err = build.BuildApiService()
		if err != nil {
			return errors.Wrapf(xerr.NewErrCode(xerr.USER_OPERATION_ERR),
				"生成api服务代码失败：%s", err.Error())
		}
	case "buildSwaggerDoc":
		//生成swagger doc文件
		err = build.BuildSwaggerDoc()
		if err != nil {
			return errors.Wrapf(xerr.NewErrCode(xerr.USER_OPERATION_ERR),
				"生成swagger doc文件失败：%s", err.Error())
		}
	case "buildAll":
		//生成api文件
		err = build.BuildApiFile()
		if err != nil {
			return errors.Wrapf(xerr.NewErrCode(xerr.USER_OPERATION_ERR),
				"生成api文件失败：%s", err.Error())
		}
		//生成数据库curl代码
		err = build.BuildDataModel()
		if err != nil {
			return errors.Wrapf(xerr.NewErrCode(xerr.USER_OPERATION_ERR),
				"生成数据库curl代码失败：%s", err.Error())
		}
		//生成api服务代码
		err = build.BuildApiService()
		if err != nil {
			return errors.Wrapf(xerr.NewErrCode(xerr.USER_OPERATION_ERR),
				"生成api服务代码失败：%s", err.Error())
		}
		//生成swagger doc文件
		err = build.BuildSwaggerDoc()
		if err != nil {
			return errors.Wrapf(xerr.NewErrCode(xerr.USER_OPERATION_ERR),
				"生成swagger doc文件失败：%s", err.Error())
		}
	}
	return nil
}