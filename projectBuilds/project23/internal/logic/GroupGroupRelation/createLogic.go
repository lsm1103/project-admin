package GroupGroupRelation

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

func (l *CreateLogic) Create(req *types.CreateGroupGroupRelationReq) (resp *types.GroupGroupRelation, err error) {
	// 自动生成的后台管理接口v1
	sqlReq := &dataModel.GroupGroupRelation{}
	err = copier.Copy(sqlReq, req)
	if err != nil {
		return
	}
	sqlReq.Id = uniqueid.GenId()

	insertR, err := l.svcCtx.GroupGroupRelationModel.Insert(l.ctx, nil, sqlReq)
	if err != nil {
		return
	}
	sqlReq.Id, err = insertR.LastInsertId()
	if err != nil {
		return
	}
	resp = &types.GroupGroupRelation{}
	err = copier.Copy(resp, sqlReq)
	if err != nil {
		return
	}

	return
}
