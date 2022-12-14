package Application

import (
	"context"
	"fmt"
	"os/exec"
	"strconv"
	"strings"

	// "github.com/jinzhu/copier"
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

/*
- 应用版本问题，每个版本做为一个数据库记录，增加版本号字段；
- 应用表记录该应用信息，版本等其他信息写在config表；
- sql文件存储，按版本存储
- api文件存储，按版本存储
*/
func (l *BuildLogic) Build(req *types.BuildReq) error {
	app := &types.Application{}
	err := l.svcCtx.ApplicationModel.FindOne(l.ctx, nil, req.ApplicationId, app)
	if err != nil {
		return errors.Wrapf(xerr.NewErrCode(xerr.USER_OPERATION_ERR), "获取数据失败：%s", err.Error())
	}
	appConf := &[]types.Config{}
	err = l.svcCtx.JoinTableQuery.FindApplicationJoinConfig(l.ctx, req.UserId, req.ApplicationId, req.Version, appConf)
	if err != nil {
		return errors.Wrap(xerr.NewErrCodeMsg(xerr.DATA_NOT_FIND, "找不到该版本应用的配置"), err.Error())
	}
	conf := map[string]string{}
	for _,item := range *appConf{
		conf[item.Key] = item.Value
	}
	isCache, _ := strconv.ParseBool(conf["cache"])
	isStrict,_ := strconv.ParseBool(conf["strict"])

	build := buildCode.BuildCode{
		RootPkgPath: l.svcCtx.RootPkgPath,
		Info: buildCode.BuildAppInfo{
			Title:        app.ZnName,
			Desc:         app.Info,
			Author:       conf["author"],
			Email:        conf["email"],
			Version:      conf["version"],
			ProjectName:  fmt.Sprintf("%s-%s", app.EnName, conf["version"]),
			ServiceType:  conf["serviceType"],
			Host:         conf["host"],
			Port:         conf["port"],
			DataSource:   conf["dataSource"],
			CacheHost:    conf["cacheHost"],
			Style:        conf["style"],
			TemplatePath: conf["templatePath"],
			Database:     conf["database"],
			DdlArg:       buildCode.DdlArg{
				Src:           conf["src"],
				Cache:         isCache,
				Strict:        isStrict,
				IgnoreColumns: strings.Split(conf["ignoreColumns"], ","),
			},
		},
	}
	if build.Info.TemplatePath == "" {
		build.Info.TemplatePath = fmt.Sprintf("%s/libs/template", build.RootPkgPath)
	}
	if build.Info.Src == "" {
		build.Info.Src = fmt.Sprintf("%s/deploy/init.sql", build.RootPkgPath)
	}
	if len(build.Info.IgnoreColumns) == 0 {
		build.Info.IgnoreColumns = []string{"create_at", "created_at", "create_time", "update_at", "updated_at", "update_time"}
	}

	switch req.BuildType {
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
	cmd := exec.Command("/bin/sh", "-c", fmt.Sprintf("cd %s/projectBuilds/%s && go build project.go"), l.svcCtx.RootPkgPath, build.Info.ProjectName )
	err = cmd.Run()
	return nil
}
