package Doc

import (
	"context"

	"project-admin/projectBuilds/project4/internal/svc"
	"project-admin/projectBuilds/project4/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
	"project-admin/common/mocks"
)

type GetsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetsLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetsLogic {
	return GetsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetsLogic) Gets(req *types.GetsReq) (resp *types.DocList, err error) {
	// 方便前端调试的接口mock
	resp = &types.DocList{}
	mocks.RespMock(resp)
	return
}
