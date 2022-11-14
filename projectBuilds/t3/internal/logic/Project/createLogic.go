package Project

import (
	"context"

	"project-admin/projectBuilds/t3/internal/svc"
	"project-admin/projectBuilds/t3/internal/types"

	"github.com/zeromicro/go-zero/core/logx"

	"github.com/jinzhu/copier"
	"project-admin/common/uniqueid"
	dataModel "project-admin/dataModel/t3Model"
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

// false request:req *types.CreateProjectReq
func (l *CreateLogic) Create(req *types.CreateProjectReq) error {
	// 自动生成的后台管理接口  req *types.CreateProjectReq
	sqlReq := &dataModel.Project{}
	err := copier.Copy(sqlReq, req)
	if err != nil {
		return err
	}
	sqlReq.Id = uniqueid.GenId()
	_, err = l.svcCtx.ProjectModel.Insert(l.ctx, nil, sqlReq)
	if err != nil {
		return err
	}

	return nil
}
