package logic

import (
	"context"

	"project-admin/test2/internal/svc"
	"project-admin/test2/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
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
	resp = &types.SingleMsg{}
	//mocks.RespMock(resp)

	err = l.svcCtx.SingleMsgModel.FindOne(l.ctx, nil, req.Id, resp)
	if err != nil {
		return nil, err
	}

	return
}
