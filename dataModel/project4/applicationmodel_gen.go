// Code generated by goctl. DO NOT EDIT!

package project4

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
	applicationFieldNames          = builder.RawFieldNames(&Application{})
	applicationRows                = strings.Join(applicationFieldNames, ",")
	applicationRowsExpectAutoSet   = strings.Join(stringx.Remove(applicationFieldNames, "`create_at`", "`created_at`", "`create_time`", "`update_at`", "`updated_at`", "`update_time`", "`state`"), ",")
	applicationRowsWithPlaceHolder = strings.Join(stringx.Remove(applicationFieldNames, "`id`", "`create_at`", "`created_at`", "`create_time`", "`update_at`", "`updated_at`", "`update_time`"), "=?,") + "=?"
	applicationListRows            = strings.Join(builder.RawFieldNames(&Application{}), ",")

	cacheGoZeroApplicationIdPrefix                        = "cache:goZero:application:id:"
	cacheGoZeroApplicationCreateUserEnNameProjectIdPrefix = "cache:goZero:application:createUser:enName:projectId:"
)

type (
	applicationModel interface {
		Insert(ctx context.Context, session sqlx.Session, data *Application) (sql.Result, error)
		FindOne(ctx context.Context, session sqlx.Session, id int64, resp interface{}) (err error)
		FindOneByCreateUserEnNameProjectId(ctx context.Context, createUser int64, enName string, projectId string) (*Application, error)
		Update(ctx context.Context, session sqlx.Session, data *Application) error
		Delete(ctx context.Context, session sqlx.Session, id int64) error
	}

	defaultApplicationModel struct {
		sqlc.CachedConn
		table string
	}

	Application struct {
		Id         int64     `db:"id"`          // 主键
		ZnName     string    `db:"zn_name"`     // 中文名称
		EnName     string    `db:"en_name"`     // 英文名称，相当于程序名称
		Ico        string    `db:"ico"`         // 图标
		Info       string    `db:"info"`        // 简介
		CreateUser int64     `db:"create_user"` // 创建者id
		DemandIds  int64     `db:"demand_ids"`  // 需求组ids
		DocIds     string    `db:"doc_ids"`     // 文档组ids
		JoinUsers  string    `db:"join_users"`  // 参与者ids
		JoinGroups string    `db:"join_groups"` // 参与组ids
		ProjectId  string    `db:"project_id"`  // 所属项目id
		Remark     string    `db:"remark"`      // 备注
		Rank       int64     `db:"rank"`        // 排序
		State      int64     `db:"state"`       // 状态，-2删除，-1禁用，待审核0，启用1
		CreateTime time.Time `db:"create_time"` // 创建时间
		UpdateTime time.Time `db:"update_time"` // 更新时间
	}
)

func newApplicationModel(conn sqlx.SqlConn, c cache.CacheConf) *defaultApplicationModel {
	return &defaultApplicationModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`application`",
	}
}

func (m *defaultApplicationModel) Delete(ctx context.Context, session sqlx.Session, id int64) error {
	data := &Application{}
	err := m.FindOne(ctx, session, id, data)
	if err != nil {
		return err
	}
	goZeroApplicationCreateUserEnNameProjectIdKey := fmt.Sprintf("%s%v:%v:%v", cacheGoZeroApplicationCreateUserEnNameProjectIdPrefix, data.CreateUser, data.EnName, data.ProjectId)
	goZeroApplicationIdKey := fmt.Sprintf("%s%v", cacheGoZeroApplicationIdPrefix, id)

	_, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		if session != nil {
			return session.ExecCtx(ctx, query, id)
		}
		return conn.ExecCtx(ctx, query, id)
	}, goZeroApplicationCreateUserEnNameProjectIdKey, goZeroApplicationIdKey)
	return err
}

func (m *defaultApplicationModel) FindOne(ctx context.Context, session sqlx.Session, id int64, resp interface{}) (err error) {
	goZeroApplicationIdKey := fmt.Sprintf("%s%v", cacheGoZeroApplicationIdPrefix, id)
	err = m.QueryRowCtx(ctx, resp, goZeroApplicationIdKey, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", applicationRows, m.table)
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

func (m *defaultApplicationModel) FindOneByCreateUserEnNameProjectId(ctx context.Context, createUser int64, enName string, projectId string) (*Application, error) {
	goZeroApplicationCreateUserEnNameProjectIdKey := fmt.Sprintf("%s%v:%v:%v", cacheGoZeroApplicationCreateUserEnNameProjectIdPrefix, createUser, enName, projectId)
	var resp Application
	err := m.QueryRowIndexCtx(ctx, &resp, goZeroApplicationCreateUserEnNameProjectIdKey, m.formatPrimary, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) (i interface{}, e error) {
		query := fmt.Sprintf("select %s from %s where `create_user` = ? and `en_name` = ? and `project_id` = ? limit 1", applicationRows, m.table)
		if err := conn.QueryRowCtx(ctx, &resp, query, createUser, enName, projectId); err != nil {
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

func (m *defaultApplicationModel) Insert(ctx context.Context, session sqlx.Session, data *Application) (sql.Result, error) {
	goZeroApplicationCreateUserEnNameProjectIdKey := fmt.Sprintf("%s%v:%v:%v", cacheGoZeroApplicationCreateUserEnNameProjectIdPrefix, data.CreateUser, data.EnName, data.ProjectId)
	goZeroApplicationIdKey := fmt.Sprintf("%s%v", cacheGoZeroApplicationIdPrefix, data.Id)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", m.table, applicationRowsExpectAutoSet)
		if session != nil {
			return session.ExecCtx(ctx, query, data.Id, data.ZnName, data.EnName, data.Ico, data.Info, data.CreateUser, data.DemandIds, data.DocIds, data.JoinUsers, data.JoinGroups, data.ProjectId, data.Remark, data.Rank)
		}
		return conn.ExecCtx(ctx, query, data.Id, data.ZnName, data.EnName, data.Ico, data.Info, data.CreateUser, data.DemandIds, data.DocIds, data.JoinUsers, data.JoinGroups, data.ProjectId, data.Remark, data.Rank)
	}, goZeroApplicationCreateUserEnNameProjectIdKey, goZeroApplicationIdKey)
	return ret, err
}

func (m *defaultApplicationModel) Update(ctx context.Context, session sqlx.Session, data *Application) error {
	err := m.FindOne(ctx, session, data.Id, &Application{})
	if err != nil {
		return err
	}
	goZeroApplicationCreateUserEnNameProjectIdKey := fmt.Sprintf("%s%v:%v:%v", cacheGoZeroApplicationCreateUserEnNameProjectIdPrefix, data.CreateUser, data.EnName, data.ProjectId)
	goZeroApplicationIdKey := fmt.Sprintf("%s%v", cacheGoZeroApplicationIdPrefix, data.Id)

	_, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, strings.Join(sqlUtils.BuildFields(data, sqlUtils.IsEmptyValue), ", "))
		if session != nil {
			return session.Exec(query, data.Id)
		}
		return conn.Exec(query, data.Id)
	}, goZeroApplicationCreateUserEnNameProjectIdKey, goZeroApplicationIdKey)

	return err
}

func (m *defaultApplicationModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheGoZeroApplicationIdPrefix, primary)
}

func (m *defaultApplicationModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", applicationRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultApplicationModel) tableName() string {
	return m.table
}
