package Application

import (
	"context"

	"project-admin/projectBuilds/project10/internal/svc"
	"project-admin/projectBuilds/project10/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetsLogic {
	return &GetsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetsLogic) Gets(req *types.GetsReq) (resp *types.ApplicationList, err error) {
	// todo: add your logic here and delete this line

	return
}
