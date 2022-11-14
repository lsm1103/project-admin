package Doc

import (
	"context"

	"project-admin/projectBuilds/t3/internal/svc"
	"project-admin/projectBuilds/t3/internal/types"

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

// false request:req *types.GetReq
func (l *GetLogic) Get(req *types.GetReq) (resp *types.Doc, err error) {
	// 自动生成的后台管理接口  req *types.GetReq
	err = l.svcCtx.DocModel.FindOne(l.ctx, nil, req.Id, resp)
	if err != nil {
		return nil, err
	}

	return
}
