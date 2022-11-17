package project4

import (
	"context"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"

	"project-admin/common/sqlUtils"
)

var _ ProjectModel = (*customProjectModel)(nil)

type (
	// ProjectModel is an interface to be customized, add more methods here,
	// and implement the added methods in customProjectModel.
	ProjectModel interface {
		projectModel

		FindAll(in *sqlUtils.GetsReq, resp interface{}) (err error)
		SoftDelete(ctx context.Context, session sqlx.Session, data *Project) error
	}

	customProjectModel struct {
		*defaultProjectModel
	}
)

// NewProjectModel returns a model for the database table.
func NewProjectModel(conn sqlx.SqlConn, c cache.CacheConf) ProjectModel {
	return &customProjectModel{
		defaultProjectModel: newProjectModel(conn, c),
	}
}

func (m *customProjectModel) FindAll(in *sqlUtils.GetsReq, resp interface{}) (err error) {
	queryStr := sqlUtils.NewModelTool().BuildQuery(in, projectListRows, m.table)
	err = m.CachedConn.QueryRowsNoCache(resp, queryStr)
	return
}

func (m *customProjectModel) SoftDelete(ctx context.Context, session sqlx.Session, data *Project) error {
	data.State = sqlUtils.Del
	return m.Update(ctx, session, data)
}
