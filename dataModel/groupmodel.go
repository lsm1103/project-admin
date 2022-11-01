package dataModel

import (
	"context"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"

	"project-admin/common/sqlUtils"
)

var _ GroupModel = (*customGroupModel)(nil)

type (
	// GroupModel is an interface to be customized, add more methods here,
	// and implement the added methods in customGroupModel.
	GroupModel interface {
		groupModel

		FindAll(in *sqlUtils.GetsReq, resp interface{}) (err error)
		SoftDelete(ctx context.Context, session sqlx.Session, data *Group) error
	}

	customGroupModel struct {
		*defaultGroupModel
	}
)

// NewGroupModel returns a model for the database table.
func NewGroupModel(conn sqlx.SqlConn, c cache.CacheConf) GroupModel {
	return &customGroupModel{
		defaultGroupModel: newGroupModel(conn, c),
	}
}

func (m *customGroupModel) FindAll(in *sqlUtils.GetsReq, resp interface{}) (err error) {
	queryStr := sqlUtils.NewModelTool().BuildQuery(in, groupListRows, m.table)
	err = m.CachedConn.QueryRowsNoCache(resp, queryStr)
	return
}

func (m *customGroupModel) SoftDelete(ctx context.Context, session sqlx.Session, data *Group) error {
	data.State = sqlUtils.Del
	return m.Update(ctx, session, data)
}
