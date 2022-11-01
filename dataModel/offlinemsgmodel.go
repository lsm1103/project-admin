package dataModel

import (
	"context"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"

	"project-admin/common/sqlUtils"
)

var _ OfflineMsgModel = (*customOfflineMsgModel)(nil)

type (
	// OfflineMsgModel is an interface to be customized, add more methods here,
	// and implement the added methods in customOfflineMsgModel.
	OfflineMsgModel interface {
		offlineMsgModel

		FindAll(in *sqlUtils.GetsReq, resp interface{}) (err error)
		SoftDelete(ctx context.Context, session sqlx.Session, data *OfflineMsg) error
	}

	customOfflineMsgModel struct {
		*defaultOfflineMsgModel
	}
)

// NewOfflineMsgModel returns a model for the database table.
func NewOfflineMsgModel(conn sqlx.SqlConn, c cache.CacheConf) OfflineMsgModel {
	return &customOfflineMsgModel{
		defaultOfflineMsgModel: newOfflineMsgModel(conn, c),
	}
}

func (m *customOfflineMsgModel) FindAll(in *sqlUtils.GetsReq, resp interface{}) (err error) {
	queryStr := sqlUtils.NewModelTool().BuildQuery(in, offlineMsgListRows, m.table)
	err = m.CachedConn.QueryRowsNoCache(resp, queryStr)
	return
}

func (m *customOfflineMsgModel) SoftDelete(ctx context.Context, session sqlx.Session, data *OfflineMsg) error {
	data.State = sqlUtils.Del
	return m.Update(ctx, session, data)
}
