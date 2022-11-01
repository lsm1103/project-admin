package msg1

import (
	"context"

	"project-admin/test2/internal/svc"
	"project-admin/test2/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
	"project-admin/common/mocks"
)

type GetSingleMsgLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetSingleMsgLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetSingleMsgLogic {
	return GetSingleMsgLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetSingleMsgLogic) GetSingleMsg(req *types.GetReq) (resp *types.SingleMsg, err error) {
	// 方便前端调试的接口mock
	resp = &types.SingleMsg{}
	mocks.RespMock(resp)
	return
}
