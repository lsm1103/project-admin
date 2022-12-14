// Code generated by goctl. DO NOT EDIT!

package project3

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
	configFieldNames          = builder.RawFieldNames(&Config{})
	configRows                = strings.Join(configFieldNames, ",")
	configRowsExpectAutoSet   = strings.Join(stringx.Remove(configFieldNames, "`updated_at`", "`update_time`", "`create_at`", "`created_at`", "`create_time`", "`update_at`", "`state`"), ",")
	configRowsWithPlaceHolder = strings.Join(stringx.Remove(configFieldNames, "`id`", "`updated_at`", "`update_time`", "`create_at`", "`created_at`", "`create_time`", "`update_at`"), "=?,") + "=?"
	configListRows            = strings.Join(builder.RawFieldNames(&Config{}), ",")

	cacheConfigIdPrefix        = "cache:config:id:"
	cacheConfigUserIdKeyPrefix = "cache:config:userId:key:"
)

type (
	configModel interface {
		Insert(ctx context.Context, session sqlx.Session, data *Config) (sql.Result, error)
		FindOne(ctx context.Context, session sqlx.Session, id int64, resp interface{}) (err error)
		FindOneByUserIdKey(ctx context.Context, userId int64, key int64) (*Config, error)
		Update(ctx context.Context, session sqlx.Session, data *Config) error
		Delete(ctx context.Context, session sqlx.Session, id int64) error
	}

	defaultConfigModel struct {
		sqlc.CachedConn
		table string
	}

	Config struct {
		Id         int64     `db:"id"`          // 主键
		UserId     int64     `db:"user_id"`     // 用户id
		Key        int64     `db:"key"`         // key
		Value      int64     `db:"value"`       // value
		State      int64     `db:"state"`       // 状态，-2删除，-1禁用，待审核0，启用1
		CreateTime time.Time `db:"create_time"` // 创建时间
		UpdateTime time.Time `db:"update_time"` // 更新时间
	}
)

func newConfigModel(conn sqlx.SqlConn, c cache.CacheConf) *defaultConfigModel {
	return &defaultConfigModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`config`",
	}
}

func (m *defaultConfigModel) Delete(ctx context.Context, session sqlx.Session, id int64) error {
	data := &Config{}
	err := m.FindOne(ctx, session, id, data)
	if err != nil {
		return err
	}
	configIdKey := fmt.Sprintf("%s%v", cacheConfigIdPrefix, id)
	configUserIdKeyKey := fmt.Sprintf("%s%v:%v", cacheConfigUserIdKeyPrefix, data.UserId, data.Key)

	_, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		if session != nil {
			return session.ExecCtx(ctx, query, id)
		}
		return conn.ExecCtx(ctx, query, id)
	}, configIdKey, configUserIdKeyKey)
	return err
}

func (m *defaultConfigModel) FindOne(ctx context.Context, session sqlx.Session, id int64, resp interface{}) (err error) {
	configIdKey := fmt.Sprintf("%s%v", cacheConfigIdPrefix, id)
	err = m.QueryRowCtx(ctx, resp, configIdKey, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", configRows, m.table)
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

func (m *defaultConfigModel) FindOneByUserIdKey(ctx context.Context, userId int64, key int64) (*Config, error) {
	configUserIdKeyKey := fmt.Sprintf("%s%v:%v", cacheConfigUserIdKeyPrefix, userId, key)
	var resp Config
	err := m.QueryRowIndexCtx(ctx, &resp, configUserIdKeyKey, m.formatPrimary, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) (i interface{}, e error) {
		query := fmt.Sprintf("select %s from %s where `user_id` = ? and `key` = ? limit 1", configRows, m.table)
		if err := conn.QueryRowCtx(ctx, &resp, query, userId, key); err != nil {
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

func (m *defaultConfigModel) Insert(ctx context.Context, session sqlx.Session, data *Config) (sql.Result, error) {
	configIdKey := fmt.Sprintf("%s%v", cacheConfigIdPrefix, data.Id)
	configUserIdKeyKey := fmt.Sprintf("%s%v:%v", cacheConfigUserIdKeyPrefix, data.UserId, data.Key)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?)", m.table, configRowsExpectAutoSet)
		if session != nil {
			return session.ExecCtx(ctx, query, data.Id, data.UserId, data.Key, data.Value)
		}
		return conn.ExecCtx(ctx, query, data.Id, data.UserId, data.Key, data.Value)
	}, configIdKey, configUserIdKeyKey)
	return ret, err
}

func (m *defaultConfigModel) Update(ctx context.Context, session sqlx.Session, data *Config) error {
	err := m.FindOne(ctx, session, data.Id, &Config{})
	if err != nil {
		return err
	}
	configIdKey := fmt.Sprintf("%s%v", cacheConfigIdPrefix, data.Id)
	configUserIdKeyKey := fmt.Sprintf("%s%v:%v", cacheConfigUserIdKeyPrefix, data.UserId, data.Key)

	_, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, strings.Join(sqlUtils.BuildFields(data, sqlUtils.IsEmptyValue), ", "))
		if session != nil {
			return session.Exec(query, data.Id)
		}
		return conn.Exec(query, data.Id)
	}, configIdKey, configUserIdKeyKey)

	return err
}

func (m *defaultConfigModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheConfigIdPrefix, primary)
}

func (m *defaultConfigModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", configRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultConfigModel) tableName() string {
	return m.table
}
