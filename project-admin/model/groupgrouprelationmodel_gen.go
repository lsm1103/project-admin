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
	groupGroupRelationFieldNames          = builder.RawFieldNames(&GroupGroupRelation{})
	groupGroupRelationRows                = strings.Join(groupGroupRelationFieldNames, ",")
	groupGroupRelationRowsExpectAutoSet   = strings.Join(stringx.Remove(groupGroupRelationFieldNames, "`update_time`", "`create_at`", "`created_at`", "`create_time`", "`update_at`", "`updated_at`", "`state`"), ",")
	groupGroupRelationRowsWithPlaceHolder = strings.Join(stringx.Remove(groupGroupRelationFieldNames, "`id`", "`update_time`", "`create_at`", "`created_at`", "`create_time`", "`update_at`", "`updated_at`"), "=?,") + "=?"
	groupGroupRelationListRows            = strings.Join(builder.RawFieldNames(&GroupGroupRelation{}), ",")

	cacheGroupGroupRelationIdPrefix                                         = "cache:groupGroupRelation:id:"
	cacheGroupGroupRelationCreateUserMasterGroupIdFromGroupIdRelationPrefix = "cache:groupGroupRelation:createUser:masterGroupId:fromGroupId:relation:"
)

type (
	groupGroupRelationModel interface {
		Insert(ctx context.Context, session sqlx.Session, data *GroupGroupRelation) (sql.Result, error)
		FindOne(ctx context.Context, session sqlx.Session, id int64, resp interface{}) (err error)
		FindOneByCreateUserMasterGroupIdFromGroupIdRelation(ctx context.Context, createUser int64, masterGroupId int64, fromGroupId int64, relation int64) (*GroupGroupRelation, error)
		Update(ctx context.Context, session sqlx.Session, data *GroupGroupRelation) error
		Delete(ctx context.Context, session sqlx.Session, id int64) error
	}

	defaultGroupGroupRelationModel struct {
		sqlc.CachedConn
		table string
	}

	GroupGroupRelation struct {
		Id            int64     `db:"id"`              // 自增主键
		CreateUser    int64     `db:"create_user"`     // 创建者id
		MasterGroupId int64     `db:"master_group_id"` // 主组id
		FromGroupId   int64     `db:"from_group_id"`   // 从组id
		Relation      int64     `db:"relation"`        // 关系，-1禁止(非自己)，1可读(非自己)，2可读写(非自己)
		State         int64     `db:"state"`           // 状态，-2删除，-1禁用，待审核0，启用1
		CreateTime    time.Time `db:"create_time"`     // 创建时间
		UpdateTime    time.Time `db:"update_time"`     // 更新时间
	}
)

func newGroupGroupRelationModel(conn sqlx.SqlConn, c cache.CacheConf) *defaultGroupGroupRelationModel {
	return &defaultGroupGroupRelationModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`group_group_relation`",
	}
}

func (m *defaultGroupGroupRelationModel) Delete(ctx context.Context, session sqlx.Session, id int64) error {
	data := &GroupGroupRelation{}
	err := m.FindOne(ctx, session, id, data)
	if err != nil {
		return err
	}
	groupGroupRelationCreateUserMasterGroupIdFromGroupIdRelationKey := fmt.Sprintf("%s%v:%v:%v:%v", cacheGroupGroupRelationCreateUserMasterGroupIdFromGroupIdRelationPrefix, data.CreateUser, data.MasterGroupId, data.FromGroupId, data.Relation)
	groupGroupRelationIdKey := fmt.Sprintf("%s%v", cacheGroupGroupRelationIdPrefix, id)

	_, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		if session != nil {
			return session.ExecCtx(ctx, query, id)
		}
		return conn.ExecCtx(ctx, query, id)
	}, groupGroupRelationCreateUserMasterGroupIdFromGroupIdRelationKey, groupGroupRelationIdKey)
	return err
}

func (m *defaultGroupGroupRelationModel) FindOne(ctx context.Context, session sqlx.Session, id int64, resp interface{}) (err error) {
	groupGroupRelationIdKey := fmt.Sprintf("%s%v", cacheGroupGroupRelationIdPrefix, id)
	err = m.QueryRowCtx(ctx, resp, groupGroupRelationIdKey, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", groupGroupRelationRows, m.table)
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

func (m *defaultGroupGroupRelationModel) FindOneByCreateUserMasterGroupIdFromGroupIdRelation(ctx context.Context, createUser int64, masterGroupId int64, fromGroupId int64, relation int64) (*GroupGroupRelation, error) {
	groupGroupRelationCreateUserMasterGroupIdFromGroupIdRelationKey := fmt.Sprintf("%s%v:%v:%v:%v", cacheGroupGroupRelationCreateUserMasterGroupIdFromGroupIdRelationPrefix, createUser, masterGroupId, fromGroupId, relation)
	var resp GroupGroupRelation
	err := m.QueryRowIndexCtx(ctx, &resp, groupGroupRelationCreateUserMasterGroupIdFromGroupIdRelationKey, m.formatPrimary, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) (i interface{}, e error) {
		query := fmt.Sprintf("select %s from %s where `create_user` = ? and `master_group_id` = ? and `from_group_id` = ? and `relation` = ? limit 1", groupGroupRelationRows, m.table)
		if err := conn.QueryRowCtx(ctx, &resp, query, createUser, masterGroupId, fromGroupId, relation); err != nil {
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

func (m *defaultGroupGroupRelationModel) Insert(ctx context.Context, session sqlx.Session, data *GroupGroupRelation) (sql.Result, error) {
	groupGroupRelationCreateUserMasterGroupIdFromGroupIdRelationKey := fmt.Sprintf("%s%v:%v:%v:%v", cacheGroupGroupRelationCreateUserMasterGroupIdFromGroupIdRelationPrefix, data.CreateUser, data.MasterGroupId, data.FromGroupId, data.Relation)
	groupGroupRelationIdKey := fmt.Sprintf("%s%v", cacheGroupGroupRelationIdPrefix, data.Id)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?)", m.table, groupGroupRelationRowsExpectAutoSet)
		if session != nil {
			return session.ExecCtx(ctx, query, data.Id, data.CreateUser, data.MasterGroupId, data.FromGroupId, data.Relation)
		}
		return conn.ExecCtx(ctx, query, data.Id, data.CreateUser, data.MasterGroupId, data.FromGroupId, data.Relation)
	}, groupGroupRelationCreateUserMasterGroupIdFromGroupIdRelationKey, groupGroupRelationIdKey)
	return ret, err
}

func (m *defaultGroupGroupRelationModel) Update(ctx context.Context, session sqlx.Session, data *GroupGroupRelation) error {
	err := m.FindOne(ctx, session, data.Id, &GroupGroupRelation{})
	if err != nil {
		return err
	}
	groupGroupRelationCreateUserMasterGroupIdFromGroupIdRelationKey := fmt.Sprintf("%s%v:%v:%v:%v", cacheGroupGroupRelationCreateUserMasterGroupIdFromGroupIdRelationPrefix, data.CreateUser, data.MasterGroupId, data.FromGroupId, data.Relation)
	groupGroupRelationIdKey := fmt.Sprintf("%s%v", cacheGroupGroupRelationIdPrefix, data.Id)
	_, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, strings.Join(sqlUtils.BuildFields(data, sqlUtils.IsEmptyValue), ", "))
		if session != nil {
			return session.Exec(query, data.Id)
		}
		return conn.Exec(query, data.Id)
	}, groupGroupRelationCreateUserMasterGroupIdFromGroupIdRelationKey, groupGroupRelationIdKey)
	return err
}

func (m *defaultGroupGroupRelationModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheGroupGroupRelationIdPrefix, primary)
}

func (m *defaultGroupGroupRelationModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", groupGroupRelationRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultGroupGroupRelationModel) tableName() string {
	return m.table
}
