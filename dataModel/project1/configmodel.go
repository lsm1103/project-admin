package project1

import (
	"context"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"

	"project-admin/common/sqlUtils"
)

var _ ConfigModel = (*customConfigModel)(nil)

type (
	// ConfigModel is an interface to be customized, add more methods here,
	// and implement the added methods in customConfigModel.
	ConfigModel interface {
		configModel

		FindAll(in *sqlUtils.GetsReq, resp interface{}) (err error)
		SoftDelete(ctx context.Context, session sqlx.Session, data *Config) error
	}

	customConfigModel struct {
		*defaultConfigModel
	}
)

// NewConfigModel returns a model for the database table.
func NewConfigModel(conn sqlx.SqlConn, c cache.CacheConf) ConfigModel {
	return &customConfigModel{
		defaultConfigModel: newConfigModel(conn, c),
	}
}

func (m *customConfigModel) FindAll(in *sqlUtils.GetsReq, resp interface{}) (err error) {
	queryStr := sqlUtils.NewModelTool().BuildQuery(in, configListRows, m.table)
	err = m.CachedConn.QueryRowsNoCache(resp, queryStr)
	return
}

func (m *customConfigModel) SoftDelete(ctx context.Context, session sqlx.Session, data *Config) error {
	data.State = sqlUtils.Del
	return m.Update(ctx, session, data)
}
