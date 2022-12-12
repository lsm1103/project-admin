package svc

import (
	"project-admin/projectBuilds/project24/internal/config"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
	dataModel "project-admin/dataModel/project24"
)

type ServiceContext struct {
	Config config.Config

	GroupModel              dataModel.GroupModel
	UserGroupModel          dataModel.UserGroupModel
	GroupGroupRelationModel dataModel.GroupGroupRelationModel
	ConfigModel             dataModel.ConfigModel
	ProjectModel            dataModel.ProjectModel
	ApplicationModel        dataModel.ApplicationModel
	ApplicationInfoModel    dataModel.ApplicationInfoModel
	DocModel                dataModel.DocModel
	DocHistoryModel         dataModel.DocHistoryModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,

		GroupModel:              dataModel.NewGroupModel(sqlx.NewMysql(c.DB.DataSource), c.Cache),
		UserGroupModel:          dataModel.NewUserGroupModel(sqlx.NewMysql(c.DB.DataSource), c.Cache),
		GroupGroupRelationModel: dataModel.NewGroupGroupRelationModel(sqlx.NewMysql(c.DB.DataSource), c.Cache),
		ConfigModel:             dataModel.NewConfigModel(sqlx.NewMysql(c.DB.DataSource), c.Cache),
		ProjectModel:            dataModel.NewProjectModel(sqlx.NewMysql(c.DB.DataSource), c.Cache),
		ApplicationModel:        dataModel.NewApplicationModel(sqlx.NewMysql(c.DB.DataSource), c.Cache),
		ApplicationInfoModel:    dataModel.NewApplicationInfoModel(sqlx.NewMysql(c.DB.DataSource), c.Cache),
		DocModel:                dataModel.NewDocModel(sqlx.NewMysql(c.DB.DataSource), c.Cache),
		DocHistoryModel:         dataModel.NewDocHistoryModel(sqlx.NewMysql(c.DB.DataSource), c.Cache),
	}
}
