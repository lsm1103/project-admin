package Doc

import (
	"context"

	"project-admin/projectBuilds/project1/internal/svc"
	"project-admin/projectBuilds/project1/internal/types"

	"github.com/zeromicro/go-zero/core/logx"

	"github.com/jinzhu/copier"

	dataModel "project-admin/dataModel/project1"
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

func (l *UpdateLogic) Update(req *types.UpdateDocReq) error {
	// 自动生成的后台管理接口
	sqlReq := &dataModel.Doc{}
	err := copier.Copy(sqlReq, req)
	if err != nil {
		return err
	}
	err = l.svcCtx.DocModel.Update(l.ctx, nil, sqlReq)
	if err != nil {
		return err
	}

	return nil
}
