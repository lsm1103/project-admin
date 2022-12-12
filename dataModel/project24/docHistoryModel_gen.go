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
	docHistoryFieldNames          = builder.RawFieldNames(&DocHistory{})
	docHistoryRows                = strings.Join(docHistoryFieldNames, ",")
	docHistoryRowsExpectAutoSet   = strings.Join(stringx.Remove(docHistoryFieldNames, "`update_at`", "`updated_at`", "`update_time`", "`create_at`", "`created_at`", "`create_time`", "`state`"), ",")
	docHistoryRowsWithPlaceHolder = strings.Join(stringx.Remove(docHistoryFieldNames, "`id`", "`update_at`", "`updated_at`", "`update_time`", "`create_at`", "`created_at`", "`create_time`"), "=?,") + "=?"
	docHistoryListRows            = strings.Join(builder.RawFieldNames(&DocHistory{}), ",")

	cacheDocHistoryIdPrefix = "cache:docHistory:id:"
)

type (
	docHistoryModel interface {
		Insert(ctx context.Context, session sqlx.Session, data *DocHistory) (sql.Result, error)
		FindOne(ctx context.Context, session sqlx.Session, id int64, resp interface{}) (err error)
		Update(ctx context.Context, session sqlx.Session, data *DocHistory) error
		Delete(ctx context.Context, session sqlx.Session, id int64) error
	}

	defaultDocHistoryModel struct {
		sqlc.CachedConn
		table string
	}

	DocHistory struct {
		Id         int64     `db:"id"`          // 主键
		PreContent string    `db:"pre_content"` // 编辑内容
		CreateUser int64     `db:"create_user"` // 所属用户
		DocId      int64     `db:"doc_id"`      // 文档id
		CreateTime time.Time `db:"create_time"` // 创建时间
	}
)

func newDocHistoryModel(conn sqlx.SqlConn, c cache.CacheConf) *defaultDocHistoryModel {
	return &defaultDocHistoryModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`doc_history`",
	}
}

func (m *defaultDocHistoryModel) Delete(ctx context.Context, session sqlx.Session, id int64) error {
	docHistoryIdKey := fmt.Sprintf("%s%v", cacheDocHistoryIdPrefix, id)

	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		if session != nil {
			return session.ExecCtx(ctx, query, id)
		}
		return conn.ExecCtx(ctx, query, id)
	}, docHistoryIdKey)
	return err
}

func (m *defaultDocHistoryModel) FindOne(ctx context.Context, session sqlx.Session, id int64, resp interface{}) (err error) {
	docHistoryIdKey := fmt.Sprintf("%s%v", cacheDocHistoryIdPrefix, id)
	err = m.QueryRowCtx(ctx, resp, docHistoryIdKey, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", docHistoryRows, m.table)
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

func (m *defaultDocHistoryModel) Insert(ctx context.Context, session sqlx.Session, data *DocHistory) (sql.Result, error) {
	docHistoryIdKey := fmt.Sprintf("%s%v", cacheDocHistoryIdPrefix, data.Id)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?)", m.table, docHistoryRowsExpectAutoSet)
		if session != nil {
			return session.ExecCtx(ctx, query, data.Id, data.PreContent, data.CreateUser, data.DocId)
		}
		return conn.ExecCtx(ctx, query, data.Id, data.PreContent, data.CreateUser, data.DocId)
	}, docHistoryIdKey)
	return ret, err
}

func (m *defaultDocHistoryModel) Update(ctx context.Context, session sqlx.Session, data *DocHistory) error {
	docHistoryIdKey := fmt.Sprintf("%s%v", cacheDocHistoryIdPrefix, data.Id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, strings.Join(sqlUtils.BuildFields(data, sqlUtils.IsEmptyValue), ", "))
		if session != nil {
			return session.Exec(query, data.Id)
		}
		return conn.Exec(query, data.Id)
	}, docHistoryIdKey)
	return err
}

func (m *defaultDocHistoryModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheDocHistoryIdPrefix, primary)
}

func (m *defaultDocHistoryModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", docHistoryRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultDocHistoryModel) tableName() string {
	return m.table
}
