package msg

import (
	"context"

	"project-admin/test2/internal/svc"
	"project-admin/test2/internal/types"

	"github.com/zeromicro/go-zero/core/logx"

	"github.com/jinzhu/copier"
	"project-admin/common/uniqueid"
	"project-admin/dataModel"
)

type CreateGroupMsgLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateGroupMsgLogic(ctx context.Context, svcCtx *svc.ServiceContext) CreateGroupMsgLogic {
	return CreateGroupMsgLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateGroupMsgLogic) CreateGroupMsg(req *types.CreateGroupMsgReq) error {
	// 自动生成的后台管理接口
	sqlReq := &dataModel.GroupMsg{}
	err := copier.Copy(sqlReq, req)
	if err != nil {
		return err
	}
	sqlReq.Id = uniqueid.GenId()
	_, err = l.svcCtx.GroupMsgModel.Insert(l.ctx, nil, sqlReq)
	if err != nil {
		return err
	}

	return nil
}
