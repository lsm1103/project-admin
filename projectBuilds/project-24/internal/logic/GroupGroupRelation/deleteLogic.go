package GroupGroupRelation

import (
	"context"

	"project-admin/projectBuilds/project-24/internal/svc"
	"project-admin/projectBuilds/project-24/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) DeleteLogic {
	return DeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteLogic) Delete(req *types.DeleteReq) error {
	// 自动生成的后台管理接口v1
	err := l.svcCtx.GroupGroupRelationModel.Delete(l.ctx, nil, req.Id)
	if err != nil {
		return err
	}

	return nil
}
