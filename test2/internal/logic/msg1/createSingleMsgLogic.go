package msg1

import (
	"context"

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
	// 方便前端调试的接口mock

	return nil
}
