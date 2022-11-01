package dataModel_

import (
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
	// "<no value>"
)

var (
	userGroupFieldNames          = builder.RawFieldNames(&UserGroup{})
	userGroupRows                = strings.Join(userGroupFieldNames, ",")
	userGroupRowsExpectAutoSet   = strings.Join(stringx.Remove(userGroupFieldNames, "`create_time`", "`update_time`"), ",")
	userGroupRowsWithPlaceHolder = strings.Join(stringx.Remove(userGroupFieldNames, "`id`", "`create_time`", "`update_time`"), "=?,") + "=?"
	userGroupListRows            = strings.Join(builder.RawFieldNames(&UserGroup{}), ",")

	cacheUserGroupIdPrefix            = "cache:userGroup:id:"
	cacheUserGroupUserIdGroupIdPrefix = "cache:userGroup:userId:groupId:"
)

type (
	UserGroupModel interface {
		Insert(session sqlx.Session, data *UserGroup) (sql.Result, error)
		FindOne(id int64) (*UserGroup, error)
		FindAll(in *GetsReq) ([]*UserGroup, error)
		FindOneByUserIdGroupId(userId int64, groupId int64) (*UserGroup, error)
		Update(session sqlx.Session, data *UserGroup) error
		Delete(id int64) error
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

func NewUserGroupModel(conn sqlx.SqlConn, c cache.CacheConf) UserGroupModel {
	return &defaultUserGroupModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`user_group`",
	}
}

func (m *defaultUserGroupModel) Insert(session sqlx.Session, data *UserGroup) (sql.Result, error) {
	userGroupIdKey := fmt.Sprintf("%s%v", cacheUserGroupIdPrefix, data.Id)
	userGroupUserIdGroupIdKey := fmt.Sprintf("%s%v:%v", cacheUserGroupUserIdGroupIdPrefix, data.UserId, data.GroupId)
	ret, err := m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?)", m.table, userGroupRowsExpectAutoSet)
		if session != nil {
			return session.Exec(query, data.Id, data.UserId, data.GroupId, data.State)
		}
		return conn.Exec(query, data.Id, data.UserId, data.GroupId, data.State)
	}, userGroupIdKey, userGroupUserIdGroupIdKey)
	return ret, err
}

func (m *defaultUserGroupModel) FindOne(id int64) (*UserGroup, error) {
	userGroupIdKey := fmt.Sprintf("%s%v", cacheUserGroupIdPrefix, id)
	var resp UserGroup
	err := m.QueryRow(&resp, userGroupIdKey, func(conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", userGroupRows, m.table)
		return conn.QueryRow(v, query, id)
	})
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultUserGroupModel) FindOneByUserIdGroupId(userId int64, groupId int64) (*UserGroup, error) {
	userGroupUserIdGroupIdKey := fmt.Sprintf("%s%v:%v", cacheUserGroupUserIdGroupIdPrefix, userId, groupId)
	var resp UserGroup
	err := m.QueryRowIndex(&resp, userGroupUserIdGroupIdKey, m.formatPrimary, func(conn sqlx.SqlConn, v interface{}) (i interface{}, e error) {
		query := fmt.Sprintf("select %s from %s where `user_id` = ? and `group_id` = ? limit 1", userGroupRows, m.table)
		if err := conn.QueryRow(&resp, query, userId, groupId); err != nil {
			return nil, err
		}
		return resp.Id, nil
	}, m.queryPrimary)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultUserGroupModel) Update(session sqlx.Session, data *UserGroup) error {
	userGroupIdKey := fmt.Sprintf("%s%v", cacheUserGroupIdPrefix, data.Id)
	userGroupUserIdGroupIdKey := fmt.Sprintf("%s%v:%v", cacheUserGroupUserIdGroupIdPrefix, data.UserId, data.GroupId)
	_, err := m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, userGroupRowsWithPlaceHolder)
		if session != nil {
			return session.Exec(query, data.UserId, data.GroupId, data.State, data.Id)
		}
		return conn.Exec(query, data.UserId, data.GroupId, data.State, data.Id)
	}, userGroupIdKey, userGroupUserIdGroupIdKey)
	return err
}
func (m *defaultUserGroupModel) Delete(id int64) error {
	data, err := m.FindOne(id)
	if err != nil {
		return err
	}
	userGroupIdKey := fmt.Sprintf("%s%v", cacheUserGroupIdPrefix, id)
	userGroupUserIdGroupIdKey := fmt.Sprintf("%s%v:%v", cacheUserGroupUserIdGroupIdPrefix, data.UserId, data.GroupId)
	_, err = m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.Exec(query, id)
	}, userGroupIdKey, userGroupUserIdGroupIdKey)
	return err
}

// 222333

func (m *defaultUserGroupModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheUserGroupIdPrefix, primary)
}

func (m *defaultUserGroupModel) queryPrimary(conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", userGroupRows, m.table)
	return conn.QueryRow(v, query, primary)
}

func (m *defaultUserGroupModel) FindAll(in *GetsReq) ([]*UserGroup, error) {
	resp := make([]*UserGroup, 0)
	queryStr := NewModelTool().BuildQuery(in, userGroupListRows, m.table)
	err := m.CachedConn.QueryRowsNoCache(&resp, queryStr)
	return resp, err
}
