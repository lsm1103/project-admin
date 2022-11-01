package msg1

import (
	"context"

	"project-admin/test2/internal/svc"
	"project-admin/test2/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
	"project-admin/common/mocks"
)

type GetsSingleMsgLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetsSingleMsgLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetsSingleMsgLogic {
	return GetsSingleMsgLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetsSingleMsgLogic) GetsSingleMsg(req *types.GetsReq) (resp *types.SingleMsgList, err error) {
	// 方便前端调试的接口mock
	resp = &types.SingleMsgList{}
	mocks.RespMock(resp)
	return
}
