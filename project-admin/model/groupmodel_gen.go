// Code generated by goctl. DO NOT EDIT!

package model

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
	groupFieldNames          = builder.RawFieldNames(&Group{})
	groupRows                = strings.Join(groupFieldNames, ",")
	groupRowsExpectAutoSet   = strings.Join(stringx.Remove(groupFieldNames, "`updated_at`", "`update_time`", "`create_at`", "`created_at`", "`create_time`", "`update_at`", "`state`"), ",")
	groupRowsWithPlaceHolder = strings.Join(stringx.Remove(groupFieldNames, "`id`", "`updated_at`", "`update_time`", "`create_at`", "`created_at`", "`create_time`", "`update_at`"), "=?,") + "=?"
	groupListRows            = strings.Join(builder.RawFieldNames(&Group{}), ",")

	cacheGroupIdPrefix             = "cache:group:id:"
	cacheGroupCreateUserNamePrefix = "cache:group:createUser:name:"
)

type (
	groupModel interface {
		Insert(ctx context.Context, session sqlx.Session, data *Group) (sql.Result, error)
		FindOne(ctx context.Context, session sqlx.Session, id int64, resp interface{}) (err error)
		FindOneByCreateUserName(ctx context.Context, createUser int64, name string) (*Group, error)
		Update(ctx context.Context, session sqlx.Session, data *Group) error
		Delete(ctx context.Context, session sqlx.Session, id int64) error
	}

	defaultGroupModel struct {
		sqlc.CachedConn
		table string
	}

	Group struct {
		Id         int64     `db:"id"`          // 自增主键
		Name       string    `db:"name"`        // 组名称
		CreateUser int64     `db:"create_user"` // 创建者id
		Ico        string    `db:"ico"`         // 组图标
		Remark     string    `db:"remark"`      // 备注
		ParentId   int64     `db:"parent_id"`   // 父级id
		GroupType  int64     `db:"group_type"`  // 类型: 1部门、2用户组、3群组、4圈子、5话题
		Rank       int64     `db:"rank"`        // 排序
		State      int64     `db:"state"`       // 状态，-2删除，-1禁用，待审核0，启用1
		CreateTime time.Time `db:"create_time"` // 创建时间
		UpdateTime time.Time `db:"update_time"` // 更新时间
	}
)

func newGroupModel(conn sqlx.SqlConn, c cache.CacheConf) *defaultGroupModel {
	return &defaultGroupModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`group`",
	}
}

func (m *defaultGroupModel) Delete(ctx context.Context, session sqlx.Session, id int64) error {
	data := &Group{}
	err := m.FindOne(ctx, session, id, data)
	if err != nil {
		return err
	}
	groupCreateUserNameKey := fmt.Sprintf("%s%v:%v", cacheGroupCreateUserNamePrefix, data.CreateUser, data.Name)
	groupIdKey := fmt.Sprintf("%s%v", cacheGroupIdPrefix, id)

	_, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		if session != nil {
			return session.ExecCtx(ctx, query, id)
		}
		return conn.ExecCtx(ctx, query, id)
	}, groupCreateUserNameKey, groupIdKey)
	return err
}

func (m *defaultGroupModel) FindOne(ctx context.Context, session sqlx.Session, id int64, resp interface{}) (err error) {
	groupIdKey := fmt.Sprintf("%s%v", cacheGroupIdPrefix, id)
	err = m.QueryRowCtx(ctx, resp, groupIdKey, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", groupRows, m.table)
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

func (m *defaultGroupModel) FindOneByCreateUserName(ctx context.Context, createUser int64, name string) (*Group, error) {
	groupCreateUserNameKey := fmt.Sprintf("%s%v:%v", cacheGroupCreateUserNamePrefix, createUser, name)
	var resp Group
	err := m.QueryRowIndexCtx(ctx, &resp, groupCreateUserNameKey, m.formatPrimary, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) (i interface{}, e error) {
		query := fmt.Sprintf("select %s from %s where `create_user` = ? and `name` = ? limit 1", groupRows, m.table)
		if err := conn.QueryRowCtx(ctx, &resp, query, createUser, name); err != nil {
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

func (m *defaultGroupModel) Insert(ctx context.Context, session sqlx.Session, data *Group) (sql.Result, error) {
	groupCreateUserNameKey := fmt.Sprintf("%s%v:%v", cacheGroupCreateUserNamePrefix, data.CreateUser, data.Name)
	groupIdKey := fmt.Sprintf("%s%v", cacheGroupIdPrefix, data.Id)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?)", m.table, groupRowsExpectAutoSet)
		if session != nil {
			return session.ExecCtx(ctx, query, data.Id, data.Name, data.CreateUser, data.Ico, data.Remark, data.ParentId, data.GroupType, data.Rank)
		}
		return conn.ExecCtx(ctx, query, data.Id, data.Name, data.CreateUser, data.Ico, data.Remark, data.ParentId, data.GroupType, data.Rank)
	}, groupCreateUserNameKey, groupIdKey)
	return ret, err
}

func (m *defaultGroupModel) Update(ctx context.Context, session sqlx.Session, data *Group) error {
	err := m.FindOne(ctx, session, data.Id, &Group{})
	if err != nil {
		return err
	}
	groupCreateUserNameKey := fmt.Sprintf("%s%v:%v", cacheGroupCreateUserNamePrefix, data.CreateUser, data.Name)
	groupIdKey := fmt.Sprintf("%s%v", cacheGroupIdPrefix, data.Id)
	_, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, strings.Join(sqlUtils.BuildFields(data, sqlUtils.IsEmptyValue), ", "))
		if session != nil {
			return session.Exec(query, data.Id)
		}
		return conn.Exec(query, data.Id)
	}, groupCreateUserNameKey, groupIdKey)
	return err
}

func (m *defaultGroupModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheGroupIdPrefix, primary)
}

func (m *defaultGroupModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", groupRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultGroupModel) tableName() string {
	return m.table
}
