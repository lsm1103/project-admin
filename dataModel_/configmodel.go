package dataModel_

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ ConfigModel = (*customConfigModel)(nil)

type (
	// ConfigModel is an interface to be customized, add more methods here,
	// and implement the added methods in customConfigModel.
	ConfigModel interface {
		configModel
	}

	customConfigModel struct {
		*defaultConfigModel
	}
)

// NewConfigModel returns a model for the database table.
func NewConfigModel(conn sqlx.SqlConn, c cache.CacheConf) ConfigModel {
	return &customConfigModel{
		defaultConfigModel: newConfigModel(conn, c),
	}
}
