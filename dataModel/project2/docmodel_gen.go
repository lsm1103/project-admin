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
	docFieldNames          = builder.RawFieldNames(&Doc{})
	docRows                = strings.Join(docFieldNames, ",")
	docRowsExpectAutoSet   = strings.Join(stringx.Remove(docFieldNames, "`create_at`", "`created_at`", "`create_time`", "`update_at`", "`updated_at`", "`update_time`", "`state`"), ",")
	docRowsWithPlaceHolder = strings.Join(stringx.Remove(docFieldNames, "`id`", "`create_at`", "`created_at`", "`create_time`", "`update_at`", "`updated_at`", "`update_time`"), "=?,") + "=?"
	docListRows            = strings.Join(builder.RawFieldNames(&Doc{}), ",")

	cacheDocIdPrefix             = "cache:doc:id:"
	cacheDocNameCreateUserPrefix = "cache:doc:name:createUser:"
)

type (
	docModel interface {
		Insert(ctx context.Context, session sqlx.Session, data *Doc) (sql.Result, error)
		FindOne(ctx context.Context, session sqlx.Session, id int64, resp interface{}) (err error)
		FindOneByNameCreateUser(ctx context.Context, name string, createUser int64) (*Doc, error)
		Update(ctx context.Context, session sqlx.Session, data *Doc) error
		Delete(ctx context.Context, session sqlx.Session, id int64) error
	}

	defaultDocModel struct {
		sqlc.CachedConn
		table string
	}

	Doc struct {
		Id           int64     `db:"id"`            // 主键
		Name         string    `db:"name"`          // 文档标题
		CreateUser   int64     `db:"create_user"`   // 所属用户
		PreContent   string    `db:"pre_content"`   // 编辑内容
		Content      string    `db:"content"`       // 文档内容
		ParentDoc    int64     `db:"parent_doc"`    // 上级文档
		GroupId      int64     `db:"group_id"`      // 所属文档组
		Sort         int64     `db:"sort"`          // 排序
		EditorMode   int64     `db:"editor_mode"`   // 编辑器模式,1表示Editormd编辑器，2表示Vditor编辑器，3表示iceEditor编辑器
		OpenChildren int64     `db:"open_children"` // 展开下级目录
		ShowChildren int64     `db:"show_children"` // 显示下级文档
		State        int64     `db:"state"`         // 文档状态，-2删除，-1禁用，待审核-草稿0，启用1
		CreateTime   time.Time `db:"create_time"`   // 创建时间
		UpdateTime   time.Time `db:"update_time"`   // 更新时间
	}
)

func newDocModel(conn sqlx.SqlConn, c cache.CacheConf) *defaultDocModel {
	return &defaultDocModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`doc`",
	}
}

func (m *defaultDocModel) Delete(ctx context.Context, session sqlx.Session, id int64) error {
	data := &Doc{}
	err := m.FindOne(ctx, session, id, data)
	if err != nil {
		return err
	}
	docIdKey := fmt.Sprintf("%s%v", cacheDocIdPrefix, id)
	docNameCreateUserKey := fmt.Sprintf("%s%v:%v", cacheDocNameCreateUserPrefix, data.Name, data.CreateUser)

	_, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		if session != nil {
			return session.ExecCtx(ctx, query, id)
		}
		return conn.ExecCtx(ctx, query, id)
	}, docIdKey, docNameCreateUserKey)
	return err
}

func (m *defaultDocModel) FindOne(ctx context.Context, session sqlx.Session, id int64, resp interface{}) (err error) {
	docIdKey := fmt.Sprintf("%s%v", cacheDocIdPrefix, id)
	err = m.QueryRowCtx(ctx, resp, docIdKey, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", docRows, m.table)
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

func (m *defaultDocModel) FindOneByNameCreateUser(ctx context.Context, name string, createUser int64) (*Doc, error) {
	docNameCreateUserKey := fmt.Sprintf("%s%v:%v", cacheDocNameCreateUserPrefix, name, createUser)
	var resp Doc
	err := m.QueryRowIndexCtx(ctx, &resp, docNameCreateUserKey, m.formatPrimary, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) (i interface{}, e error) {
		query := fmt.Sprintf("select %s from %s where `name` = ? and `create_user` = ? limit 1", docRows, m.table)
		if err := conn.QueryRowCtx(ctx, &resp, query, name, createUser); err != nil {
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

func (m *defaultDocModel) Insert(ctx context.Context, session sqlx.Session, data *Doc) (sql.Result, error) {
	docIdKey := fmt.Sprintf("%s%v", cacheDocIdPrefix, data.Id)
	docNameCreateUserKey := fmt.Sprintf("%s%v:%v", cacheDocNameCreateUserPrefix, data.Name, data.CreateUser)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", m.table, docRowsExpectAutoSet)
		if session != nil {
			return session.ExecCtx(ctx, query, data.Id, data.Name, data.CreateUser, data.PreContent, data.Content, data.ParentDoc, data.GroupId, data.Sort, data.EditorMode, data.OpenChildren, data.ShowChildren)
		}
		return conn.ExecCtx(ctx, query, data.Id, data.Name, data.CreateUser, data.PreContent, data.Content, data.ParentDoc, data.GroupId, data.Sort, data.EditorMode, data.OpenChildren, data.ShowChildren)
	}, docIdKey, docNameCreateUserKey)
	return ret, err
}

func (m *defaultDocModel) Update(ctx context.Context, session sqlx.Session, data *Doc) error {
	err := m.FindOne(ctx, session, data.Id, &Doc{})
	if err != nil {
		return err
	}
	docIdKey := fmt.Sprintf("%s%v", cacheDocIdPrefix, data.Id)
	docNameCreateUserKey := fmt.Sprintf("%s%v:%v", cacheDocNameCreateUserPrefix, data.Name, data.CreateUser)

	_, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, strings.Join(sqlUtils.BuildFields(data, sqlUtils.IsEmptyValue), ", "))
		if session != nil {
			return session.Exec(query, data.Id)
		}
		return conn.Exec(query, data.Id)
	}, docIdKey, docNameCreateUserKey)

	return err
}

func (m *defaultDocModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheDocIdPrefix, primary)
}

func (m *defaultDocModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", docRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultDocModel) tableName() string {
	return m.table
}
