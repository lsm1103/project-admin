package GroupGroupRelation

import (
	"context"

	"project-admin/projectBuilds/project24/internal/svc"
	"project-admin/projectBuilds/project24/internal/types"

	"github.com/zeromicro/go-zero/core/logx"

	"github.com/jinzhu/copier"

	dataModel "project-admin/dataModel/project24"
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

func (l *UpdateLogic) Update(req *types.UpdateGroupGroupRelationReq) (resp *types.GroupGroupRelation, err error) {
	// 自动生成的后台管理接口v1
	sqlReq := &dataModel.GroupGroupRelation{}
	err = copier.Copy(sqlReq, req)
	if err != nil {
		return
	}

	err = l.svcCtx.GroupGroupRelationModel.Update(l.ctx, nil, sqlReq)
	if err != nil {
		return
	}
	resp = &types.GroupGroupRelation{}
	err = copier.Copy(resp, sqlReq)
	if err != nil {
		return
	}

	return
}
