// Code generated by goctl. DO NOT EDIT!

package project2

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
	applicationConfigFieldNames          = builder.RawFieldNames(&ApplicationConfig{})
	applicationConfigRows                = strings.Join(applicationConfigFieldNames, ",")
	applicationConfigRowsExpectAutoSet   = strings.Join(stringx.Remove(applicationConfigFieldNames, "`create_time`", "`update_at`", "`updated_at`", "`update_time`", "`create_at`", "`created_at`", "`state`"), ",")
	applicationConfigRowsWithPlaceHolder = strings.Join(stringx.Remove(applicationConfigFieldNames, "`id`", "`create_time`", "`update_at`", "`updated_at`", "`update_time`", "`create_at`", "`created_at`"), "=?,") + "=?"
	applicationConfigListRows            = strings.Join(builder.RawFieldNames(&ApplicationConfig{}), ",")

	cacheApplicationConfigIdPrefix                              = "cache:applicationConfig:id:"
	cacheApplicationConfigCreateUserApplicationIdConfigIdPrefix = "cache:applicationConfig:createUser:applicationId:configId:"
)

type (
	applicationConfigModel interface {
		Insert(ctx context.Context, session sqlx.Session, data *ApplicationConfig) (sql.Result, error)
		FindOne(ctx context.Context, session sqlx.Session, id int64, resp interface{}) (err error)
		FindOneByCreateUserApplicationIdConfigId(ctx context.Context, createUser int64, applicationId int64, configId int64) (*ApplicationConfig, error)
		Update(ctx context.Context, session sqlx.Session, data *ApplicationConfig) error
		Delete(ctx context.Context, session sqlx.Session, id int64) error
	}

	defaultApplicationConfigModel struct {
		sqlc.CachedConn
		table string
	}

	ApplicationConfig struct {
		Id            int64     `db:"id"`             // 主键
		CreateUser    int64     `db:"create_user"`    // 所属用户
		ApplicationId int64     `db:"application_id"` // 应用id
		ConfigId      int64     `db:"config_id"`      // 配置id
		State         int64     `db:"state"`          // 状态，-2删除，-1禁用，待审核0，启用1
		CreateTime    time.Time `db:"create_time"`    // 创建时间
		UpdateTime    time.Time `db:"update_time"`    // 更新时间
	}
)

func newApplicationConfigModel(conn sqlx.SqlConn, c cache.CacheConf) *defaultApplicationConfigModel {
	return &defaultApplicationConfigModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`application_config`",
	}
}

func (m *defaultApplicationConfigModel) Delete(ctx context.Context, session sqlx.Session, id int64) error {
	data := &ApplicationConfig{}
	err := m.FindOne(ctx, session, id, data)
	if err != nil {
		return err
	}
	applicationConfigCreateUserApplicationIdConfigIdKey := fmt.Sprintf("%s%v:%v:%v", cacheApplicationConfigCreateUserApplicationIdConfigIdPrefix, data.CreateUser, data.ApplicationId, data.ConfigId)
	applicationConfigIdKey := fmt.Sprintf("%s%v", cacheApplicationConfigIdPrefix, id)

	_, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		if session != nil {
			return session.ExecCtx(ctx, query, id)
		}
		return conn.ExecCtx(ctx, query, id)
	}, applicationConfigCreateUserApplicationIdConfigIdKey, applicationConfigIdKey)
	return err
}

func (m *defaultApplicationConfigModel) FindOne(ctx context.Context, session sqlx.Session, id int64, resp interface{}) (err error) {
	applicationConfigIdKey := fmt.Sprintf("%s%v", cacheApplicationConfigIdPrefix, id)
	err = m.QueryRowCtx(ctx, resp, applicationConfigIdKey, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", applicationConfigRows, m.table)
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

func (m *defaultApplicationConfigModel) FindOneByCreateUserApplicationIdConfigId(ctx context.Context, createUser int64, applicationId int64, configId int64) (*ApplicationConfig, error) {
	applicationConfigCreateUserApplicationIdConfigIdKey := fmt.Sprintf("%s%v:%v:%v", cacheApplicationConfigCreateUserApplicationIdConfigIdPrefix, createUser, applicationId, configId)
	var resp ApplicationConfig
	err := m.QueryRowIndexCtx(ctx, &resp, applicationConfigCreateUserApplicationIdConfigIdKey, m.formatPrimary, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) (i interface{}, e error) {
		query := fmt.Sprintf("select %s from %s where `create_user` = ? and `application_id` = ? and `config_id` = ? limit 1", applicationConfigRows, m.table)
		if err := conn.QueryRowCtx(ctx, &resp, query, createUser, applicationId, configId); err != nil {
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

func (m *defaultApplicationConfigModel) Insert(ctx context.Context, session sqlx.Session, data *ApplicationConfig) (sql.Result, error) {
	applicationConfigCreateUserApplicationIdConfigIdKey := fmt.Sprintf("%s%v:%v:%v", cacheApplicationConfigCreateUserApplicationIdConfigIdPrefix, data.CreateUser, data.ApplicationId, data.ConfigId)
	applicationConfigIdKey := fmt.Sprintf("%s%v", cacheApplicationConfigIdPrefix, data.Id)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?)", m.table, applicationConfigRowsExpectAutoSet)
		if session != nil {
			return session.ExecCtx(ctx, query, data.Id, data.CreateUser, data.ApplicationId, data.ConfigId)
		}
		return conn.ExecCtx(ctx, query, data.Id, data.CreateUser, data.ApplicationId, data.ConfigId)
	}, applicationConfigCreateUserApplicationIdConfigIdKey, applicationConfigIdKey)
	return ret, err
}

func (m *defaultApplicationConfigModel) Update(ctx context.Context, session sqlx.Session, data *ApplicationConfig) error {
	err := m.FindOne(ctx, session, data.Id, &ApplicationConfig{})
	if err != nil {
		return err
	}
	applicationConfigCreateUserApplicationIdConfigIdKey := fmt.Sprintf("%s%v:%v:%v", cacheApplicationConfigCreateUserApplicationIdConfigIdPrefix, data.CreateUser, data.ApplicationId, data.ConfigId)
	applicationConfigIdKey := fmt.Sprintf("%s%v", cacheApplicationConfigIdPrefix, data.Id)

	_, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, strings.Join(sqlUtils.BuildFields(data, sqlUtils.IsEmptyValue), ", "))
		if session != nil {
			return session.Exec(query, data.Id)
		}
		return conn.Exec(query, data.Id)
	}, applicationConfigCreateUserApplicationIdConfigIdKey, applicationConfigIdKey)

	return err
}

func (m *defaultApplicationConfigModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheApplicationConfigIdPrefix, primary)
}

func (m *defaultApplicationConfigModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", applicationConfigRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultApplicationConfigModel) tableName() string {
	return m.table
}
