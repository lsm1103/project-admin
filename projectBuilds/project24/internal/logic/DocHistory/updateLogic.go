package DocHistory

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

func (l *UpdateLogic) Update(req *types.UpdateDocHistoryReq) (resp *types.DocHistory, err error) {
	// 自动生成的后台管理接口v1
	sqlReq := &dataModel.DocHistory{}
	err = copier.Copy(sqlReq, req)
	if err != nil {
		return
	}

	err = l.svcCtx.DocHistoryModel.Update(l.ctx, nil, sqlReq)
	if err != nil {
		return
	}
	resp = &types.DocHistory{}
	err = copier.Copy(resp, sqlReq)
	if err != nil {
		return
	}

	return
}
