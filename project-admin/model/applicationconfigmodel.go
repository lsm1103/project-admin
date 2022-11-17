package model

import (
	"context"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"

	"project-admin/common/sqlUtils"
)

var _ ApplicationConfigModel = (*customApplicationConfigModel)(nil)

type (
	// ApplicationConfigModel is an interface to be customized, add more methods here,
	// and implement the added methods in customApplicationConfigModel.
	ApplicationConfigModel interface {
		applicationConfigModel

		FindAll(in *sqlUtils.GetsReq, resp interface{}) (err error)
		SoftDelete(ctx context.Context, session sqlx.Session, data *ApplicationConfig) error
	}

	customApplicationConfigModel struct {
		*defaultApplicationConfigModel
	}
)

// NewApplicationConfigModel returns a model for the database table.
func NewApplicationConfigModel(conn sqlx.SqlConn, c cache.CacheConf) ApplicationConfigModel {
	return &customApplicationConfigModel{
		defaultApplicationConfigModel: newApplicationConfigModel(conn, c),
	}
}

func (m *customApplicationConfigModel) FindAll(in *sqlUtils.GetsReq, resp interface{}) (err error) {
	queryStr := sqlUtils.NewModelTool().BuildQuery(in, applicationConfigListRows, m.table)
	err = m.CachedConn.QueryRowsNoCache(resp, queryStr)
	return
}

func (m *customApplicationConfigModel) SoftDelete(ctx context.Context, session sqlx.Session, data *ApplicationConfig) error {
	data.State = sqlUtils.Del
	return m.Update(ctx, session, data)
}
