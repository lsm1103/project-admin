package ApplicationConfig

import (
	"context"

	"project-admin/projectBuilds/project3/internal/svc"
	"project-admin/projectBuilds/project3/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) CreateLogic {
	return CreateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateLogic) Create(req *types.CreateApplicationConfigReq) error {
	// 方便前端调试的接口mock

	return nil
}
