package dataModel

import (
	"context"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"

	"project-admin/common/sqlUtils"
)

var _ FriendModel = (*customFriendModel)(nil)

type (
	// FriendModel is an interface to be customized, add more methods here,
	// and implement the added methods in customFriendModel.
	FriendModel interface {
		friendModel

		FindAll(in *sqlUtils.GetsReq, resp interface{}) (err error)
		SoftDelete(ctx context.Context, session sqlx.Session, data *Friend) error
	}

	customFriendModel struct {
		*defaultFriendModel
	}
)

// NewFriendModel returns a model for the database table.
func NewFriendModel(conn sqlx.SqlConn, c cache.CacheConf) FriendModel {
	return &customFriendModel{
		defaultFriendModel: newFriendModel(conn, c),
	}
}

func (m *customFriendModel) FindAll(in *sqlUtils.GetsReq, resp interface{}) (err error) {
	queryStr := sqlUtils.NewModelTool().BuildQuery(in, friendListRows, m.table)
	err = m.CachedConn.QueryRowsNoCache(resp, queryStr)
	return
}

func (m *customFriendModel) SoftDelete(ctx context.Context, session sqlx.Session, data *Friend) error {
	data.State = sqlUtils.Del
	return m.Update(ctx, session, data)
}
