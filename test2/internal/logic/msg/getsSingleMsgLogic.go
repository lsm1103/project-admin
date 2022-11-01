package msg

import (
	"context"

	"project-admin/test2/internal/svc"
	"project-admin/test2/internal/types"

	"github.com/zeromicro/go-zero/core/logx"

	"project-admin/common/sqlUtils"
)

type GetsSingleMsgLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetsSingleMsgLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetsSingleMsgLogic {
	return GetsSingleMsgLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetsSingleMsgLogic) GetsSingleMsg(req *sqlUtils.GetsReq) (resp *types.SingleMsgList, err error) {
	// 自动生成的后台管理接口  req *types.GetsReq
	resp = &types.SingleMsgList{Current: req.Current, PageSize: req.PageSize}
	err = l.svcCtx.SingleMsgModel.FindAll(req, &resp.List)
	if err != nil {
		return nil, err
	}
	if int64(len(resp.List)) > req.PageSize {
		resp.IsNext = true
		resp.List = resp.List[:req.PageSize]
	}
	return
}
