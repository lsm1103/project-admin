package UserGroup

import (
	"context"

	"project-admin/projectBuilds/project24/internal/svc"
	"project-admin/projectBuilds/project24/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
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

func (l *GetLogic) Get(req *types.GetReq) (resp *types.UserGroup, err error) {
	// 自动生成的后台管理接口v1
	resp = &types.UserGroup{}
	err = l.svcCtx.UserGroupModel.FindOne(l.ctx, nil, req.Id, resp)
	if err != nil {
		return nil, err
	}

	return
}
