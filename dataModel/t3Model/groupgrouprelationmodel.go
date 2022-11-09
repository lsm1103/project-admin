package t3Model

import (
	"context"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"

	"project-admin/common/sqlUtils"
)

var _ GroupGroupRelationModel = (*customGroupGroupRelationModel)(nil)

type (
	// GroupGroupRelationModel is an interface to be customized, add more methods here,
	// and implement the added methods in customGroupGroupRelationModel.
	GroupGroupRelationModel interface {
		groupGroupRelationModel

		FindAll(in *sqlUtils.GetsReq, resp interface{}) (err error)
		SoftDelete(ctx context.Context, session sqlx.Session, data *GroupGroupRelation) error
	}

	customGroupGroupRelationModel struct {
		*defaultGroupGroupRelationModel
	}
)

// NewGroupGroupRelationModel returns a model for the database table.
func NewGroupGroupRelationModel(conn sqlx.SqlConn, c cache.CacheConf) GroupGroupRelationModel {
	return &customGroupGroupRelationModel{
		defaultGroupGroupRelationModel: newGroupGroupRelationModel(conn, c),
	}
}

func (m *customGroupGroupRelationModel) FindAll(in *sqlUtils.GetsReq, resp interface{}) (err error) {
	queryStr := sqlUtils.NewModelTool().BuildQuery(in, groupGroupRelationListRows, m.table)
	err = m.CachedConn.QueryRowsNoCache(resp, queryStr)
	return
}

func (m *customGroupGroupRelationModel) SoftDelete(ctx context.Context, session sqlx.Session, data *GroupGroupRelation) error {
	data.State = sqlUtils.Del
	return m.Update(ctx, session, data)
}
