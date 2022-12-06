package UserGroup

import (
	"context"

	"project-admin/project-admin/internal/svc"
	"project-admin/project-admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"

	"github.com/jinzhu/copier"
	"project-admin/common/uniqueid"

	dataModel "project-admin/project-admin/model"
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

func (l *CreateLogic) Create(req *types.CreateUserGroupReq) (resp *types.UserGroup, err error) {
	// 自动生成的后台管理接口v1
	sqlReq := &dataModel.UserGroup{}
	err = copier.Copy(sqlReq, req)
	if err != nil {
		return
	}
	sqlReq.Id = uniqueid.GenId()
	insertR, err := l.svcCtx.UserGroupModel.Insert(l.ctx, nil, sqlReq)
	if err != nil {
		return
	}

	id, err := insertR.LastInsertId()
	if err != nil {
		return
	}
	if id != 0 {
		sqlReq.Id = id
	}
	resp = &types.UserGroup{}
	err = copier.Copy(resp, sqlReq)
	if err != nil {
		return
	}

	return
}
