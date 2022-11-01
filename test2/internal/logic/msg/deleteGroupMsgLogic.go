package msg

import (
	"context"

	"project-admin/test2/internal/svc"
	"project-admin/test2/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteGroupMsgLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteGroupMsgLogic(ctx context.Context, svcCtx *svc.ServiceContext) DeleteGroupMsgLogic {
	return DeleteGroupMsgLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteGroupMsgLogic) DeleteGroupMsg(req *types.DeleteReq) error {
	// 自动生成的后台管理接口
	err := l.svcCtx.GroupMsgModel.Delete(l.ctx, nil, req.Id)
	if err != nil {
		return err
	}

	return nil
}
