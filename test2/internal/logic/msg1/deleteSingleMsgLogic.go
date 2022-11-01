package msg1

import (
	"context"

	"project-admin/test2/internal/svc"
	"project-admin/test2/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteSingleMsgLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteSingleMsgLogic(ctx context.Context, svcCtx *svc.ServiceContext) DeleteSingleMsgLogic {
	return DeleteSingleMsgLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteSingleMsgLogic) DeleteSingleMsg(req *types.DeleteReq) error {
	// 方便前端调试的接口mock

	return nil
}
