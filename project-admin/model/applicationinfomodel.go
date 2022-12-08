package model

import (
	"context"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"

	"project-admin/common/sqlUtils"
)

var _ ApplicationInfoModel = (*customApplicationInfoModel)(nil)

type (
	// ApplicationInfoModel is an interface to be customized, add more methods here,
	// and implement the added methods in customApplicationInfoModel.
	ApplicationInfoModel interface {
		applicationInfoModel

		FindAll(in *sqlUtils.GetsReq, resp interface{}) (err error)
		SoftDelete(ctx context.Context, session sqlx.Session, data *ApplicationInfo) error
	}

	customApplicationInfoModel struct {
		*defaultApplicationInfoModel
	}
)

// NewApplicationInfoModel returns a model for the database table.
func NewApplicationInfoModel(conn sqlx.SqlConn, c cache.CacheConf) ApplicationInfoModel {
	return &customApplicationInfoModel{
		defaultApplicationInfoModel: newApplicationInfoModel(conn, c),
	}
}

func (m *customApplicationInfoModel) FindAll(in *sqlUtils.GetsReq, resp interface{}) (err error) {
	queryStr := sqlUtils.NewModelTool().BuildQuery(in, applicationInfoListRows, m.table)
	err = m.CachedConn.QueryRowsNoCache(resp, queryStr)
	return
}

func (m *customApplicationInfoModel) SoftDelete(ctx context.Context, session sqlx.Session, data *ApplicationInfo) error {
	data.State = sqlUtils.Del
	err := m.Update(ctx, session, data)
	return err
}
