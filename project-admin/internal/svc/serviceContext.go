package svc

import (
	"os"
	"project-admin/common/sqlUtils"
	"project-admin/project-admin/internal/config"
	"strings"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
	dataModel "project-admin/project-admin/model"
)

type ServiceContext struct {
	Config      config.Config
	RootPkgPath string

	JoinTableQuery          sqlUtils.JoinTableQuery
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
	pwd, err := os.Getwd()
	if err != nil {
		return nil
	}
	return &ServiceContext{
		Config:      c,
		RootPkgPath: strings.Replace(pwd, "/project-admin", "", 1),

		JoinTableQuery:          sqlUtils.NewJoinTableQuery(sqlx.NewMysql(c.DB.DataSource), c.Cache),
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
