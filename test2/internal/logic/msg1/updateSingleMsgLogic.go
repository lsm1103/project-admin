package msg1

import (
	"context"

	"project-admin/test2/internal/svc"
	"project-admin/test2/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateSingleMsgLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateSingleMsgLogic(ctx context.Context, svcCtx *svc.ServiceContext) UpdateSingleMsgLogic {
	return UpdateSingleMsgLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateSingleMsgLogic) UpdateSingleMsg(req *types.UpdateSingleMsgReq) error {
	// 方便前端调试的接口mock

	return nil
}
