package Application

import (
	"context"
	"fmt"
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
		return err
	}
	var sh string
	switch req.RunType {
	case "http":
		sh = fmt.Sprintf("cd %s/projectBuilds/%s-%s && nohup ./project &", l.svcCtx.RootPkgPath, app.EnName, req.Version)
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
	err = cmd.Run()
	fmt.Printf("[fc.execScript-pid:%d,err:%+v],\nsh:%s\n", cmd.Process.Pid, err, sh)
	data,err := cmd.CombinedOutput()
	fmt.Println(string(data))

	//go func (){
	//	cmd := exec.Command("/bin/sh", "-c", sh)
	//	err = cmd.Run()
	//	fmt.Printf("[fc.execScript-pid:%d,err:%+v],\nsh:%s\n", cmd.Process.Pid, err, sh)
	//	data,err := cmd.CombinedOutput()
	//	fmt.Println(string(data), err)
	//}()

	return err
}
