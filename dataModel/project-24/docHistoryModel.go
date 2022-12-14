package project_24

import (
	"context"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"

	"project-admin/common/sqlUtils"
)

var _ DocHistoryModel = (*customDocHistoryModel)(nil)

type (
	// DocHistoryModel is an interface to be customized, add more methods here,
	// and implement the added methods in customDocHistoryModel.
	DocHistoryModel interface {
		docHistoryModel

		FindAll(in *sqlUtils.GetsReq, resp interface{}) (err error)
		SoftDelete(ctx context.Context, session sqlx.Session, data *DocHistory) error
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

func (m *customDocHistoryModel) FindAll(in *sqlUtils.GetsReq, resp interface{}) (err error) {
	queryStr := sqlUtils.NewModelTool().BuildQuery(in, docHistoryListRows, m.table)
	err = m.CachedConn.QueryRowsNoCache(resp, queryStr)
	return
}

func (m *customDocHistoryModel) SoftDelete(ctx context.Context, session sqlx.Session, data *DocHistory) error {
	return sqlUtils.ErrNotState
}
