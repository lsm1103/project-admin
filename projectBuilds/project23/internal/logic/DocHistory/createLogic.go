package DocHistory

import (
	"context"

	"project-admin/projectBuilds/project23/internal/svc"
	"project-admin/projectBuilds/project23/internal/types"

	"github.com/zeromicro/go-zero/core/logx"

	"github.com/jinzhu/copier"
	"project-admin/common/uniqueid"

	dataModel "project-admin/dataModel/project23"
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

func (l *CreateLogic) Create(req *types.CreateDocHistoryReq) (resp *types.DocHistory, err error) {
	// 自动生成的后台管理接口
	sqlReq := &dataModel.DocHistory{}
	err := copier.Copy(sqlReq, req)
	if err != nil {
		return err
	}
	sqlReq.Id = uniqueid.GenId()
	_, err = l.svcCtx.DocHistoryModel.Insert(l.ctx, nil, sqlReq)
	if err != nil {
		return err
	}

	return
}
