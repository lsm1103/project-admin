package GroupMsg

import (
	"context"

	"project-admin/projectBuilds/t2/internal/svc"
	"project-admin/projectBuilds/t2/internal/types"

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

func (l *GetsLogic) Gets(req *sqlUtils.GetsReq) (resp *types.GroupMsgList, err error) {
	// 自动生成的后台管理接口  req *types.GetsReq
	resp = &types.GroupMsgList{Current: req.Current, PageSize: req.PageSize}
	err = l.svcCtx.GroupMsgModel.FindAll(req, &resp.List)
	if err != nil {
		return nil, err
	}
	if int64(len(resp.List)) > req.PageSize {
		resp.IsNext = true
		resp.List = resp.List[:req.PageSize]
	}
	return
}
