package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"

	"project-admin/dataModel"
	"project-admin/projectBuilds/t2/internal/config"
)

type ServiceContext struct {
	Config config.Config

	GroupModel      dataModel.GroupModel
	UserGroupModel  dataModel.UserGroupModel
	FriendModel     dataModel.FriendModel
	OfflineMsgModel dataModel.OfflineMsgModel
	SingleMsgModel  dataModel.SingleMsgModel
	GroupMsgModel   dataModel.GroupMsgModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,

		GroupModel:      dataModel.NewGroupModel(sqlx.NewMysql(c.DB.DataSource), c.Cache),
		UserGroupModel:  dataModel.NewUserGroupModel(sqlx.NewMysql(c.DB.DataSource), c.Cache),
		FriendModel:     dataModel.NewFriendModel(sqlx.NewMysql(c.DB.DataSource), c.Cache),
		OfflineMsgModel: dataModel.NewOfflineMsgModel(sqlx.NewMysql(c.DB.DataSource), c.Cache),
		SingleMsgModel:  dataModel.NewSingleMsgModel(sqlx.NewMysql(c.DB.DataSource), c.Cache),
		GroupMsgModel:   dataModel.NewGroupMsgModel(sqlx.NewMysql(c.DB.DataSource), c.Cache),
	}
}
