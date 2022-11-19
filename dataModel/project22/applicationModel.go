package project22

import (
	"context"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"

	"project-admin/common/sqlUtils"
)

var _ ApplicationModel = (*customApplicationModel)(nil)

type (
	// ApplicationModel is an interface to be customized, add more methods here,
	// and implement the added methods in customApplicationModel.
	ApplicationModel interface {
		applicationModel

		FindAll(in *sqlUtils.GetsReq, resp interface{}) (err error)
		SoftDelete(ctx context.Context, session sqlx.Session, data *Application) error
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

func (m *customApplicationModel) FindAll(in *sqlUtils.GetsReq, resp interface{}) (err error) {
	queryStr := sqlUtils.NewModelTool().BuildQuery(in, applicationListRows, m.table)
	err = m.CachedConn.QueryRowsNoCache(resp, queryStr)
	return
}

func (m *customApplicationModel) SoftDelete(ctx context.Context, session sqlx.Session, data *Application) error {
	data.State = sqlUtils.Del
	return m.Update(ctx, session, data)
}
