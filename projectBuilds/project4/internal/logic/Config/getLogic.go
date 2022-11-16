package Config

import (
	"context"

	"project-admin/projectBuilds/project4/internal/svc"
	"project-admin/projectBuilds/project4/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
	"project-admin/common/mocks"
)

type GetLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetLogic {
	return GetLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetLogic) Get(req *types.GetReq) (resp *types.Config, err error) {
	// 方便前端调试的接口mock
	resp = &types.Config{}
	mocks.RespMock(resp)
	return
}
