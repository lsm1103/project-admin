package dataModel_

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ GroupGroupRelationModel = (*customGroupGroupRelationModel)(nil)

type (
	// GroupGroupRelationModel is an interface to be customized, add more methods here,
	// and implement the added methods in customGroupGroupRelationModel.
	GroupGroupRelationModel interface {
		groupGroupRelationModel
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
