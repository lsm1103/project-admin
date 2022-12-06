package Project

import (
	"context"

	"project-admin/project-admin/internal/svc"
	"project-admin/project-admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"

	"project-admin/common/sqlUtils"
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

func (l *GetsLogic) Gets(req *sqlUtils.GetsReq) (resp *types.ProjectList, err error) {
	// 自动生成的后台管理接口v1
	resp = &types.ProjectList{Current: req.Current, PageSize: req.PageSize}
	err = l.svcCtx.ProjectModel.FindAll(req, &resp.List)
	if err != nil {
		return nil, err
	}
	if int64(len(resp.List)) > req.PageSize {
		resp.IsNext = true
		resp.List = resp.List[:req.PageSize]
	}
	return
}
