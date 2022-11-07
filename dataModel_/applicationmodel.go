package dataModel_

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ ApplicationModel = (*customApplicationModel)(nil)

type (
	// ApplicationModel is an interface to be customized, add more methods here,
	// and implement the added methods in customApplicationModel.
	ApplicationModel interface {
		applicationModel
	}

	customApplicationModel struct {
		*defaultApplicationModel
	}
)

// NewApplicationModel returns a model for the database table.
func NewApplicationModel(conn sqlx.SqlConn, c cache.CacheConf) ApplicationModel {
	return &customApplicationModel{
		defaultApplicationModel: newApplicationModel(conn, c),
	}
}
