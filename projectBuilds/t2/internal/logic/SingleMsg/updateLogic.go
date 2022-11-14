package SingleMsg

import (
	"context"

	"project-admin/projectBuilds/t2/internal/svc"
	"project-admin/projectBuilds/t2/internal/types"

	"github.com/zeromicro/go-zero/core/logx"

	"github.com/jinzhu/copier"
	"project-admin/dataModel"
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

func (l *UpdateLogic) Update(req *types.UpdateSingleMsgReq) error {
	// 自动生成的后台管理接口  req *types.UpdateSingleMsgReq
	sqlReq := &dataModel.SingleMsg{}
	err := copier.Copy(sqlReq, req)
	if err != nil {
		return err
	}
	err = l.svcCtx.SingleMsgModel.Update(l.ctx, nil, sqlReq)
	if err != nil {
		return err
	}

	return nil
}
