package Application

import (
	"bytes"
	"context"
	"fmt"
	"github.com/pkg/errors"
	"os"
	"os/exec"
	"project-admin/common/xerr"

	"project-admin/project-admin/internal/svc"
	"project-admin/project-admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RunServiceLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRunServiceLogic(ctx context.Context, svcCtx *svc.ServiceContext) RunServiceLogic {
	return RunServiceLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RunServiceLogic) RunService(req *types.RunServiceReq) error {
	app := &types.Application{}
	err := l.svcCtx.ApplicationModel.FindOne(l.ctx, nil, req.ApplicationId, app)
	if err != nil {
		return errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "查询该应用失败：%s", err.Error())
	}
	var sh string
	switch req.RunType {
	case "http":
		appPath := fmt.Sprintf("%s/projectBuilds/%s-%s", l.svcCtx.RootPkgPath, app.EnName, req.Version)
		cmd_ := exec.Command("/bin/sh", "-c", "/usr/local/go/bin/go build project.go")
		cmd_.Dir = appPath

		var out bytes.Buffer
		cmd_.Stdout = &out
		cmd_.Stderr = os.Stderr

		err = cmd_.Start()
		if err != nil {
			l.Errorf("err:%v", err)
		}
		err = cmd_.Wait()
		l.Info("out: ",out.String())
		if err != nil{
			l.Errorf("err:%v", err)
			return errors.Wrap(xerr.NewErrCodeMsg(xerr.SERVER_COMMON_ERROR, "服务编译失败"),err.Error())
		}
		sh = fmt.Sprintf("cd %s/projectBuilds/%s-%s && ./project", l.svcCtx.RootPkgPath, app.EnName, req.Version)
	case "rpc":
		return xerr.NewErrCodeMsg(xerr.USER_OPERATION_ERR, "rpc服务功能开发中，静候佳音🐱！")
	case "websocket":
		return xerr.NewErrCodeMsg(xerr.USER_OPERATION_ERR, "websocket服务功能开发中，静候佳音🐱！")
	case "tcp":
		return xerr.NewErrCodeMsg(xerr.USER_OPERATION_ERR, "tcp服务功能开发中，静候佳音🐱！")
	case "mqtt":
		return xerr.NewErrCodeMsg(xerr.USER_OPERATION_ERR, "mqtt服务功能开发中，静候佳音🐱！")
	}
	cmd := exec.Command("/bin/sh", "-c", sh)
	err = cmd.Start()
	l.Info("sh:%s\n[fc.execScript-pid:%d], err:%v\n", sh, cmd.Process.Pid, err)
	if err != nil{
		return xerr.NewErrCodeMsg(xerr.SERVER_COMMON_ERROR, "服务运行失败")
	}
	return nil
}
