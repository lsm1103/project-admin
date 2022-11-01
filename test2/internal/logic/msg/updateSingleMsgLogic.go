package msg

import (
	"context"

	"project-admin/test2/internal/svc"
	"project-admin/test2/internal/types"

	"github.com/zeromicro/go-zero/core/logx"

	"github.com/jinzhu/copier"
	"project-admin/dataModel"
)

type UpdateSingleMsgLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateSingleMsgLogic(ctx context.Context, svcCtx *svc.ServiceContext) UpdateSingleMsgLogic {
	return UpdateSingleMsgLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateSingleMsgLogic) UpdateSingleMsg(req *types.UpdateSingleMsgReq) error {
	// 自动生成的后台管理接口
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
