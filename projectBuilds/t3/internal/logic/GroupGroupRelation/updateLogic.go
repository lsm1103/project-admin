package GroupGroupRelation

import (
	"context"

	"project-admin/projectBuilds/t3/internal/svc"
	"project-admin/projectBuilds/t3/internal/types"

	"github.com/zeromicro/go-zero/core/logx"

	"github.com/jinzhu/copier"
	dataModel "project-admin/dataModel/t3Model"
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

// false request:req *types.UpdateGroupGroupRelationReq
func (l *UpdateLogic) Update(req *types.UpdateGroupGroupRelationReq) error {
	// 自动生成的后台管理接口  req *types.UpdateGroupGroupRelationReq
	sqlReq := &dataModel.GroupGroupRelation{}
	err := copier.Copy(sqlReq, req)
	if err != nil {
		return err
	}
	err = l.svcCtx.GroupGroupRelationModel.Update(l.ctx, nil, sqlReq)
	if err != nil {
		return err
	}

	return nil
}
