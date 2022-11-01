package msg

import (
	"context"

	"project-admin/test2/internal/svc"
	"project-admin/test2/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetGroupMsgLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetGroupMsgLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetGroupMsgLogic {
	return GetGroupMsgLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetGroupMsgLogic) GetGroupMsg(req *types.GetReq) (resp *types.GroupMsg, err error) {
	// 自动生成的后台管理接口
	err = l.svcCtx.GroupMsgModel.FindOne(l.ctx, nil, req.Id, resp)
	if err != nil {
		return nil, err
	}

	return
}
