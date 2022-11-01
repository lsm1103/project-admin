package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"project-admin/common/sqlUtils"
	"project-admin/dataModel"
	"project-admin/test2/internal/config"
)

type ServiceContext struct {
	Config config.Config

	ModelHandle    sqlUtils.ModelHandle
	JoinTableQuery sqlUtils.JoinTableQuery
	SingleMsgModel dataModel.SingleMsgModel
	GroupMsgModel dataModel.GroupMsgModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,

		ModelHandle:    sqlUtils.NewModelHandle(sqlx.NewMysql(c.DB.DataSource ), c.Cache),
		JoinTableQuery: sqlUtils.NewJoinTableQuery(sqlx.NewMysql(c.DB.DataSource ), c.Cache),
		SingleMsgModel: dataModel.NewSingleMsgModel(sqlx.NewMysql(c.DB.DataSource ), c.Cache),
		GroupMsgModel: dataModel.NewGroupMsgModel(sqlx.NewMysql(c.DB.DataSource ), c.Cache),
	}
}
