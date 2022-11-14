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
	userGroupFieldNames          = builder.RawFieldNames(&UserGroup{})
	userGroupRows                = strings.Join(userGroupFieldNames, ",")
	userGroupRowsExpectAutoSet   = strings.Join(stringx.Remove(userGroupFieldNames, "`create_at`", "`created_at`", "`create_time`", "`update_at`", "`updated_at`", "`update_time`", "`state`"), ",")
	userGroupRowsWithPlaceHolder = strings.Join(stringx.Remove(userGroupFieldNames, "`id`", "`create_at`", "`created_at`", "`create_time`", "`update_at`", "`updated_at`", "`update_time`"), "=?,") + "=?"
	userGroupListRows            = strings.Join(builder.RawFieldNames(&UserGroup{}), ",")

	cacheUserGroupIdPrefix            = "cache:userGroup:id:"
	cacheUserGroupUserIdGroupIdPrefix = "cache:userGroup:userId:groupId:"
)

type (
	userGroupModel interface {
		Insert(ctx context.Context, session sqlx.Session, data *UserGroup) (sql.Result, error)
		FindOne(ctx context.Context, session sqlx.Session, id int64, resp interface{}) (err error)
		FindOneByUserIdGroupId(ctx context.Context, userId int64, groupId int64) (*UserGroup, error)
		Update(ctx context.Context, session sqlx.Session, data *UserGroup) error
		Delete(ctx context.Context, session sqlx.Session, id int64) error
	}

	defaultUserGroupModel struct {
		sqlc.CachedConn
		table string
	}

	UserGroup struct {
		Id         int64     `db:"id"`          // 自增主键
		UserId     int64     `db:"user_id"`     // 用户id
		GroupId    int64     `db:"group_id"`    // 组id
		State      int64     `db:"state"`       // 状态，-2删除，-1禁用，待审核0，启用1
		CreateTime time.Time `db:"create_time"` // 创建时间
		UpdateTime time.Time `db:"update_time"` // 更新时间
	}
)

func newUserGroupModel(conn sqlx.SqlConn, c cache.CacheConf) *defaultUserGroupModel {
	return &defaultUserGroupModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`user_group`",
	}
}

func (m *defaultUserGroupModel) Delete(ctx context.Context, session sqlx.Session, id int64) error {
	data := &UserGroup{}
	err := m.FindOne(ctx, session, id, data)
	if err != nil {
		return err
	}
	userGroupIdKey := fmt.Sprintf("%s%v", cacheUserGroupIdPrefix, id)
	userGroupUserIdGroupIdKey := fmt.Sprintf("%s%v:%v", cacheUserGroupUserIdGroupIdPrefix, data.UserId, data.GroupId)

	_, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		if session != nil {
			return session.ExecCtx(ctx, query, id)
		}
		return conn.ExecCtx(ctx, query, id)
	}, userGroupIdKey, userGroupUserIdGroupIdKey)
	return err
}

func (m *defaultUserGroupModel) FindOne(ctx context.Context, session sqlx.Session, id int64, resp interface{}) (err error) {
	userGroupIdKey := fmt.Sprintf("%s%v", cacheUserGroupIdPrefix, id)
	err = m.QueryRowCtx(ctx, resp, userGroupIdKey, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", userGroupRows, m.table)
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

func (m *defaultUserGroupModel) FindOneByUserIdGroupId(ctx context.Context, userId int64, groupId int64) (*UserGroup, error) {
	userGroupUserIdGroupIdKey := fmt.Sprintf("%s%v:%v", cacheUserGroupUserIdGroupIdPrefix, userId, groupId)
	var resp UserGroup
	err := m.QueryRowIndexCtx(ctx, &resp, userGroupUserIdGroupIdKey, m.formatPrimary, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) (i interface{}, e error) {
		query := fmt.Sprintf("select %s from %s where `user_id` = ? and `group_id` = ? limit 1", userGroupRows, m.table)
		if err := conn.QueryRowCtx(ctx, &resp, query, userId, groupId); err != nil {
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

func (m *defaultUserGroupModel) Insert(ctx context.Context, session sqlx.Session, data *UserGroup) (sql.Result, error) {
	userGroupIdKey := fmt.Sprintf("%s%v", cacheUserGroupIdPrefix, data.Id)
	userGroupUserIdGroupIdKey := fmt.Sprintf("%s%v:%v", cacheUserGroupUserIdGroupIdPrefix, data.UserId, data.GroupId)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?)", m.table, userGroupRowsExpectAutoSet)
		if session != nil {
			return session.ExecCtx(ctx, query, data.Id, data.UserId, data.GroupId)
		}
		return conn.ExecCtx(ctx, query, data.Id, data.UserId, data.GroupId)
	}, userGroupIdKey, userGroupUserIdGroupIdKey)
	return ret, err
}

func (m *defaultUserGroupModel) Update(ctx context.Context, session sqlx.Session, data *UserGroup) error {
	err := m.FindOne(ctx, session, data.Id, &UserGroup{})
	if err != nil {
		return err
	}
	userGroupIdKey := fmt.Sprintf("%s%v", cacheUserGroupIdPrefix, data.Id)
	userGroupUserIdGroupIdKey := fmt.Sprintf("%s%v:%v", cacheUserGroupUserIdGroupIdPrefix, data.UserId, data.GroupId)

	_, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, strings.Join(sqlUtils.BuildFields(data, sqlUtils.IsEmptyValue), ", "))
		if session != nil {
			return session.Exec(query, data.Id)
		}
		return conn.Exec(query, data.Id)
	}, userGroupIdKey, userGroupUserIdGroupIdKey)

	return err
}

func (m *defaultUserGroupModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheUserGroupIdPrefix, primary)
}

func (m *defaultUserGroupModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", userGroupRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultUserGroupModel) tableName() string {
	return m.table
}
