package project10

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ DocHistoryModel = (*customDocHistoryModel)(nil)

type (
	// DocHistoryModel is an interface to be customized, add more methods here,
	// and implement the added methods in customDocHistoryModel.
	DocHistoryModel interface {
		docHistoryModel
	}

	customDocHistoryModel struct {
		*defaultDocHistoryModel
	}
)

// NewDocHistoryModel returns a model for the database table.
func NewDocHistoryModel(conn sqlx.SqlConn, c cache.CacheConf) DocHistoryModel {
	return &customDocHistoryModel{
		defaultDocHistoryModel: newDocHistoryModel(conn, c),
	}
}
