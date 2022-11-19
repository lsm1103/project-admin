package Application

import (
	"context"
	"fmt"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"project-admin/common/buildCode"
	"project-admin/common/dataModelToApi"
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

func (l *BuildLogic) Build(req *types.BuildReq) error {
	build := buildCode.BuildCode{RootPkgPath: l.svcCtx.RootPkgPath}
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
	//生成数据库curl代码
	err = build.BuildDataModel()
	if err != nil {
		return errors.Wrapf(xerr.NewErrCode(xerr.USER_OPERATION_ERR),
			"生成数据库curl代码失败：%s", err.Error())
	}
	//生成api文件
	dataModelToApi.DataModelToApi{
		build.RootPkgPath,
		dataModelToApi.ServiceInfo{
			Title:   build.Info.Title,
			Desc:    build.Info.Desc,
			Author:  build.Info.Author,
			Email:   build.Info.Email,
			Version: build.Info.Version,

			ProjectName: build.Info.ProjectName,
			ServiceType: build.Info.ServiceType,
			Host:        build.Info.Host,
			Port:        build.Info.Port,
			DataSource: build.Info.DataSource,
			CacheHost: build.Info.CacheHost,
		}, dataModelToApi.SqlParseCfg{
			Filename: build.Info.DdlArg.Src,
			Database: build.Info.Database,
			Strict: build.Info.Strict,
		},
	}.StartBuild()
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
	return nil
}
