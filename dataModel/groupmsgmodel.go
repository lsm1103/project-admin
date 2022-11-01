package dataModel

import (
	"context"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"

	"project-admin/common/sqlUtils"
)

var _ GroupMsgModel = (*customGroupMsgModel)(nil)

type (
	// GroupMsgModel is an interface to be customized, add more methods here,
	// and implement the added methods in customGroupMsgModel.
	GroupMsgModel interface {
		groupMsgModel

		FindAll(in *sqlUtils.GetsReq, resp interface{}) (err error)
		SoftDelete(ctx context.Context, session sqlx.Session, data *GroupMsg) error
	}

	customGroupMsgModel struct {
		*defaultGroupMsgModel
	}
)

// NewGroupMsgModel returns a model for the database table.
func NewGroupMsgModel(conn sqlx.SqlConn, c cache.CacheConf) GroupMsgModel {
	return &customGroupMsgModel{
		defaultGroupMsgModel: newGroupMsgModel(conn, c),
	}
}

func (m *customGroupMsgModel) FindAll(in *sqlUtils.GetsReq, resp interface{}) (err error) {
	queryStr := sqlUtils.NewModelTool().BuildQuery(in, groupMsgListRows, m.table)
	err = m.CachedConn.QueryRowsNoCache(resp, queryStr)
	return
}

func (m *customGroupMsgModel) SoftDelete(ctx context.Context, session sqlx.Session, data *GroupMsg) error {
	data.State = sqlUtils.Del
	return m.Update(ctx, session, data)
}
