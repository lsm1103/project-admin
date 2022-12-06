package model

import (
	"context"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"

	"project-admin/common/sqlUtils"
)

var _ UserGroupModel = (*customUserGroupModel)(nil)

type (
	// UserGroupModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUserGroupModel.
	UserGroupModel interface {
		userGroupModel

		FindAll(in *sqlUtils.GetsReq, resp interface{}) (err error)
		SoftDelete(ctx context.Context, session sqlx.Session, data *UserGroup) error
	}

	customUserGroupModel struct {
		*defaultUserGroupModel
	}
)

// NewUserGroupModel returns a model for the database table.
func NewUserGroupModel(conn sqlx.SqlConn, c cache.CacheConf) UserGroupModel {
	return &customUserGroupModel{
		defaultUserGroupModel: newUserGroupModel(conn, c),
	}
}

func (m *customUserGroupModel) FindAll(in *sqlUtils.GetsReq, resp interface{}) (err error) {
	queryStr := sqlUtils.NewModelTool().BuildQuery(in, userGroupListRows, m.table)
	err = m.CachedConn.QueryRowsNoCache(resp, queryStr)
	return
}

func (m *customUserGroupModel) SoftDelete(ctx context.Context, session sqlx.Session, data *UserGroup) error {
	data.State = sqlUtils.Del
	err := m.Update(ctx, session, data)
	return err
}
