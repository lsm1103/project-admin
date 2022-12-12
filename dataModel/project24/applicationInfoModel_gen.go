// Code generated by goctl. DO NOT EDIT!

package project24

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"

	"project-admin/common/sqlUtils"
)

var (
	applicationInfoFieldNames          = builder.RawFieldNames(&ApplicationInfo{})
	applicationInfoRows                = strings.Join(applicationInfoFieldNames, ",")
	applicationInfoRowsExpectAutoSet   = strings.Join(stringx.Remove(applicationInfoFieldNames, "`update_at`", "`updated_at`", "`update_time`", "`create_at`", "`created_at`", "`create_time`", "`state`"), ",")
	applicationInfoRowsWithPlaceHolder = strings.Join(stringx.Remove(applicationInfoFieldNames, "`id`", "`update_at`", "`updated_at`", "`update_time`", "`create_at`", "`created_at`", "`create_time`"), "=?,") + "=?"
	applicationInfoListRows            = strings.Join(builder.RawFieldNames(&ApplicationInfo{}), ",")

	cacheApplicationInfoIdPrefix                             = "cache:applicationInfo:id:"
	cacheApplicationInfoCreateUserApplicationIdVersionPrefix = "cache:applicationInfo:createUser:applicationId:version:"
)

type (
	applicationInfoModel interface {
		Insert(ctx context.Context, session sqlx.Session, data *ApplicationInfo) (sql.Result, error)
		FindOne(ctx context.Context, session sqlx.Session, id int64, resp interface{}) (err error)
		FindOneByCreateUserApplicationIdVersion(ctx context.Context, createUser int64, applicationId int64, version string) (*ApplicationInfo, error)
		Update(ctx context.Context, session sqlx.Session, data *ApplicationInfo) error
		Delete(ctx context.Context, session sqlx.Session, id int64) error
	}

	defaultApplicationInfoModel struct {
		sqlc.CachedConn
		table string
	}

	ApplicationInfo struct {
		Id            int64     `db:"id"`             // 主键
		CreateUser    int64     `db:"create_user"`    // 所属用户
		ApplicationId int64     `db:"application_id"` // 应用id
		Version       string    `db:"version"`        // 版本号
		ConfigIds     string    `db:"config_ids"`     // 配置ids ","号分隔
		State         int64     `db:"state"`          // 状态，-2删除，-1禁用，待审核0，启用1
		CreateTime    time.Time `db:"create_time"`    // 创建时间
		UpdateTime    time.Time `db:"update_time"`    // 更新时间
	}
)

func newApplicationInfoModel(conn sqlx.SqlConn, c cache.CacheConf) *defaultApplicationInfoModel {
	return &defaultApplicationInfoModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`application_info`",
	}
}

func (m *defaultApplicationInfoModel) Delete(ctx context.Context, session sqlx.Session, id int64) error {
	data := &ApplicationInfo{}
	err := m.FindOne(ctx, session, id, data)
	if err != nil {
		return err
	}
	applicationInfoCreateUserApplicationIdVersionKey := fmt.Sprintf("%s%v:%v:%v", cacheApplicationInfoCreateUserApplicationIdVersionPrefix, data.CreateUser, data.ApplicationId, data.Version)
	applicationInfoIdKey := fmt.Sprintf("%s%v", cacheApplicationInfoIdPrefix, id)

	_, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		if session != nil {
			return session.ExecCtx(ctx, query, id)
		}
		return conn.ExecCtx(ctx, query, id)
	}, applicationInfoCreateUserApplicationIdVersionKey, applicationInfoIdKey)
	return err
}

func (m *defaultApplicationInfoModel) FindOne(ctx context.Context, session sqlx.Session, id int64, resp interface{}) (err error) {
	applicationInfoIdKey := fmt.Sprintf("%s%v", cacheApplicationInfoIdPrefix, id)
	err = m.QueryRowCtx(ctx, resp, applicationInfoIdKey, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", applicationInfoRows, m.table)
		if session != nil {
			return session.QueryRowCtx(ctx, v, query, id)
		}
		return conn.QueryRowCtx(ctx, v, query, id)
	})
	switch err {
	case nil:
		return nil
	case sqlc.ErrNotFound:
		return sqlUtils.ErrNotFound
	default:
		return err
	}
}

func (m *defaultApplicationInfoModel) FindOneByCreateUserApplicationIdVersion(ctx context.Context, createUser int64, applicationId int64, version string) (*ApplicationInfo, error) {
	applicationInfoCreateUserApplicationIdVersionKey := fmt.Sprintf("%s%v:%v:%v", cacheApplicationInfoCreateUserApplicationIdVersionPrefix, createUser, applicationId, version)
	var resp ApplicationInfo
	err := m.QueryRowIndexCtx(ctx, &resp, applicationInfoCreateUserApplicationIdVersionKey, m.formatPrimary, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) (i interface{}, e error) {
		query := fmt.Sprintf("select %s from %s where `create_user` = ? and `application_id` = ? and `version` = ? limit 1", applicationInfoRows, m.table)
		if err := conn.QueryRowCtx(ctx, &resp, query, createUser, applicationId, version); err != nil {
			return nil, err
		}
		return resp.Id, nil
	}, m.queryPrimary)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, sqlUtils.ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultApplicationInfoModel) Insert(ctx context.Context, session sqlx.Session, data *ApplicationInfo) (sql.Result, error) {
	applicationInfoCreateUserApplicationIdVersionKey := fmt.Sprintf("%s%v:%v:%v", cacheApplicationInfoCreateUserApplicationIdVersionPrefix, data.CreateUser, data.ApplicationId, data.Version)
	applicationInfoIdKey := fmt.Sprintf("%s%v", cacheApplicationInfoIdPrefix, data.Id)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?)", m.table, applicationInfoRowsExpectAutoSet)
		if session != nil {
			return session.ExecCtx(ctx, query, data.Id, data.CreateUser, data.ApplicationId, data.Version, data.ConfigIds)
		}
		return conn.ExecCtx(ctx, query, data.Id, data.CreateUser, data.ApplicationId, data.Version, data.ConfigIds)
	}, applicationInfoCreateUserApplicationIdVersionKey, applicationInfoIdKey)
	return ret, err
}

func (m *defaultApplicationInfoModel) Update(ctx context.Context, session sqlx.Session, data *ApplicationInfo) error {
	err := m.FindOne(ctx, session, data.Id, &ApplicationInfo{})
	if err != nil {
		return err
	}
	applicationInfoCreateUserApplicationIdVersionKey := fmt.Sprintf("%s%v:%v:%v", cacheApplicationInfoCreateUserApplicationIdVersionPrefix, data.CreateUser, data.ApplicationId, data.Version)
	applicationInfoIdKey := fmt.Sprintf("%s%v", cacheApplicationInfoIdPrefix, data.Id)
	_, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, strings.Join(sqlUtils.BuildFields(data, sqlUtils.IsEmptyValue), ", "))
		if session != nil {
			return session.Exec(query, data.Id)
		}
		return conn.Exec(query, data.Id)
	}, applicationInfoCreateUserApplicationIdVersionKey, applicationInfoIdKey)
	return err
}

func (m *defaultApplicationInfoModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheApplicationInfoIdPrefix, primary)
}

func (m *defaultApplicationInfoModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", applicationInfoRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultApplicationInfoModel) tableName() string {
	return m.table
}
