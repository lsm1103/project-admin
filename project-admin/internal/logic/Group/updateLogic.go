package Group

import (
	"context"

	"project-admin/project-admin/internal/svc"
	"project-admin/project-admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"

	"github.com/jinzhu/copier"

	dataModel "project-admin/project-admin/model"
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

func (l *UpdateLogic) Update(req *types.UpdateGroupReq) error {
	// 自动生成的后台管理接口
	sqlReq := &dataModel.Group{}
	err := copier.Copy(sqlReq, req)
	if err != nil {
		return err
	}
	err = l.svcCtx.GroupModel.Update(l.ctx, nil, sqlReq)
	if err != nil {
		return err
	}

	return nil
}
