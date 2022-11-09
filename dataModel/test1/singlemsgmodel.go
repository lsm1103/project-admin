package test1

import (
	"context"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"

	"project-admin/common/sqlUtils"
)

var _ SingleMsgModel = (*customSingleMsgModel)(nil)

type (
	// SingleMsgModel is an interface to be customized, add more methods here,
	// and implement the added methods in customSingleMsgModel.
	SingleMsgModel interface {
		singleMsgModel

		FindAll(in *sqlUtils.GetsReq, resp interface{}) (err error)
		SoftDelete(ctx context.Context, session sqlx.Session, data *SingleMsg) error
	}

	customSingleMsgModel struct {
		*defaultSingleMsgModel
	}
)

// NewSingleMsgModel returns a model for the database table.
func NewSingleMsgModel(conn sqlx.SqlConn, c cache.CacheConf) SingleMsgModel {
	return &customSingleMsgModel{
		defaultSingleMsgModel: newSingleMsgModel(conn, c),
	}
}

func (m *customSingleMsgModel) FindAll(in *sqlUtils.GetsReq, resp interface{}) (err error) {
	queryStr := sqlUtils.NewModelTool().BuildQuery(in, singleMsgListRows, m.table)
	err = m.CachedConn.QueryRowsNoCache(resp, queryStr)
	return
}

func (m *customSingleMsgModel) SoftDelete(ctx context.Context, session sqlx.Session, data *SingleMsg) error {
	data.State = sqlUtils.Del
	return m.Update(ctx, session, data)
}
