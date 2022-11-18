package logic

import (
	"context"
	"github.com/jinzhu/copier"
	"project-admin/common/uniqueid"
	"project-admin/dataModel"
	"project-admin/test2/internal/svc"
	"project-admin/test2/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateSingleMsgLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateSingleMsgLogic(ctx context.Context, svcCtx *svc.ServiceContext) CreateSingleMsgLogic {
	return CreateSingleMsgLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateSingleMsgLogic) CreateSingleMsg(req *types.CreateSingleMsgReq) error {
	sqlReq := &dataModel.SingleMsg{}
	err := copier.Copy(sqlReq, req)
	if err != nil {
		return err
	}
	sqlReq.Id = uniqueid.GenId()
	_, err = l.svcCtx.SingleMsgModel.Insert(l.ctx, nil, sqlReq)
	if err != nil {
		return err
	}

	return nil
}
