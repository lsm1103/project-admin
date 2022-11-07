package dataModel_

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ DocModel = (*customDocModel)(nil)

type (
	// DocModel is an interface to be customized, add more methods here,
	// and implement the added methods in customDocModel.
	DocModel interface {
		docModel
	}

	customDocModel struct {
		*defaultDocModel
	}
)

// NewDocModel returns a model for the database table.
func NewDocModel(conn sqlx.SqlConn, c cache.CacheConf) DocModel {
	return &customDocModel{
		defaultDocModel: newDocModel(conn, c),
	}
}
