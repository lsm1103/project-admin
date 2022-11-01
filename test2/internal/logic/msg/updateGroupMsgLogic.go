package msg

import (
	"context"

	"project-admin/test2/internal/svc"
	"project-admin/test2/internal/types"

	"github.com/zeromicro/go-zero/core/logx"

	"github.com/jinzhu/copier"
	"project-admin/dataModel"
)

type UpdateGroupMsgLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateGroupMsgLogic(ctx context.Context, svcCtx *svc.ServiceContext) UpdateGroupMsgLogic {
	return UpdateGroupMsgLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateGroupMsgLogic) UpdateGroupMsg(req *types.UpdateGroupMsgReq) error {
	// 自动生成的后台管理接口
	sqlReq := &dataModel.GroupMsg{}
	err := copier.Copy(sqlReq, req)
	if err != nil {
		return err
	}
	err = l.svcCtx.GroupMsgModel.Update(l.ctx, nil, sqlReq)
	if err != nil {
		return err
	}

	return nil
}
