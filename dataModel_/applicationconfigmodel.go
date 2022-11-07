package dataModel_

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ ApplicationConfigModel = (*customApplicationConfigModel)(nil)

type (
	// ApplicationConfigModel is an interface to be customized, add more methods here,
	// and implement the added methods in customApplicationConfigModel.
	ApplicationConfigModel interface {
		applicationConfigModel
	}

	customApplicationConfigModel struct {
		*defaultApplicationConfigModel
	}
)

// NewApplicationConfigModel returns a model for the database table.
func NewApplicationConfigModel(conn sqlx.SqlConn, c cache.CacheConf) ApplicationConfigModel {
	return &customApplicationConfigModel{
		defaultApplicationConfigModel: newApplicationConfigModel(conn, c),
	}
}
