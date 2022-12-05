package project23

import (
	"context"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"

	"project-admin/common/sqlUtils"
)

var _ DocModel = (*customDocModel)(nil)

type (
	// DocModel is an interface to be customized, add more methods here,
	// and implement the added methods in customDocModel.
	DocModel interface {
		docModel

		FindAll(in *sqlUtils.GetsReq, resp interface{}) (err error)
		SoftDelete(ctx context.Context, session sqlx.Session, data *Doc) error
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

func (m *customDocModel) FindAll(in *sqlUtils.GetsReq, resp interface{}) (err error) {
	queryStr := sqlUtils.NewModelTool().BuildQuery(in, docListRows, m.table)
	err = m.CachedConn.QueryRowsNoCache(resp, queryStr)
	return
}

func (m *customDocModel) SoftDelete(ctx context.Context, session sqlx.Session, data *Doc) error {
	data.State = sqlUtils.Del
	err := m.Update(ctx, session, data)
	return err
}
