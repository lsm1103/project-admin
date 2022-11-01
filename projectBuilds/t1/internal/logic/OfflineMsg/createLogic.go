package OfflineMsg

import (
	"context"

	"project-admin/buildProject/t1/internal/svc"
	"project-admin/buildProject/t1/internal/types"

	"github.com/zeromicro/go-zero/core/logx"

	"github.com/jinzhu/copier"
	"project-admin/common/uniqueid"
	"project-admin/dataModel"
)

type CreateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) CreateLogic {
	return CreateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateLogic) Create(req *types.CreateOfflineMsgReq) error {
	// 自动生成的后台管理接口  req *types.CreateOfflineMsgReq
	sqlReq := &dataModel.OfflineMsg{}
	err := copier.Copy(sqlReq, req)
	if err != nil {
		return err
	}
	sqlReq.Id = uniqueid.GenId()
	_, err = l.svcCtx.OfflineMsgModel.Insert(l.ctx, nil, sqlReq)
	if err != nil {
		return err
	}

	return nil
}
