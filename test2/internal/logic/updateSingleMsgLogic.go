package logic

import (
	"context"
	"github.com/jinzhu/copier"
	"project-admin/dataModel"

	"project-admin/test2/internal/svc"
	"project-admin/test2/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
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
