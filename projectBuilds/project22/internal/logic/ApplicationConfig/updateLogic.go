package ApplicationConfig

import (
	"context"

	"project-admin/projectBuilds/project22/internal/svc"
	"project-admin/projectBuilds/project22/internal/types"

	"github.com/zeromicro/go-zero/core/logx"

	"github.com/jinzhu/copier"

	dataModel "project-admin/dataModel/project22"
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

func (l *UpdateLogic) Update(req *types.UpdateApplicationConfigReq) error {
	// 自动生成的后台管理接口
	sqlReq := &dataModel.ApplicationConfig{}
	err := copier.Copy(sqlReq, req)
	if err != nil {
		return err
	}
	err = l.svcCtx.ApplicationConfigModel.Update(l.ctx, nil, sqlReq)
	if err != nil {
		return err
	}

	return nil
}
