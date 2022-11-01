package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"

	"project-admin/buildProject/t1/internal/config"
	"project-admin/dataModel"
)

type ServiceContext struct {
	Config config.Config

	OfflineMsgModel dataModel.OfflineMsgModel
	SingleMsgModel  dataModel.SingleMsgModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,

		OfflineMsgModel: dataModel.NewOfflineMsgModel(sqlx.NewMysql(c.DB.DataSource), c.Cache),
		SingleMsgModel:  dataModel.NewSingleMsgModel(sqlx.NewMysql(c.DB.DataSource), c.Cache),
	}
}
