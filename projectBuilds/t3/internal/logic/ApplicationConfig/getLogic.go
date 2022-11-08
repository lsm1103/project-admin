package ApplicationConfig

import (
	"context"

	"project-admin/projectBuilds/t3/internal/svc"
	"project-admin/projectBuilds/t3/internal/types"

	"github.com/zeromicro/go-zero/core/logx"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"project-admin/common/xerr"
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

func (l *GetLogic) Get(req *types.GetReq) (resp *types.ApplicationConfig, err error) {
	// 自动生成的后台管理接口  req *types.GetReq
	resp = &types.ApplicationConfig{}
	err = l.svcCtx.ApplicationConfigModel.FindOne(l.ctx, nil, req.Id, resp)
	if err == sqlx.ErrNotFound {
		return nil, xerr.NewErrCode(xerr.DATA_NOT_FIND)
	}
	if err != nil {
		return nil, err
	}

	return
}