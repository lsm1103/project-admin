package Project

import (
	"context"

	"project-admin/projectBuilds/project-24/internal/svc"
	"project-admin/projectBuilds/project-24/internal/types"

	"github.com/zeromicro/go-zero/core/logx"

	"github.com/jinzhu/copier"

	dataModel "project-admin/dataModel/project-24"
)

type UpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) UpdateLogic {
	return UpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateLogic) Update(req *types.UpdateProjectReq) (resp *types.Project, err error) {
	// 自动生成的后台管理接口v1
	sqlReq := &dataModel.Project{}
	err = copier.Copy(sqlReq, req)
	if err != nil {
		return
	}

	err = l.svcCtx.ProjectModel.Update(l.ctx, nil, sqlReq)
	if err != nil {
		return
	}
	resp = &types.Project{}
	err = copier.Copy(resp, sqlReq)
	if err != nil {
		return
	}

	return
}
