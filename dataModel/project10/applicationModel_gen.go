// Code generated by goctl. DO NOT EDIT!

package project10

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
)

var (
	applicationFieldNames          = builder.RawFieldNames(&Application{})
	applicationRows                = strings.Join(applicationFieldNames, ",")
	applicationRowsExpectAutoSet   = strings.Join(stringx.Remove(applicationFieldNames, "`updated_at`", "`update_time`", "`create_at`", "`created_at`", "`create_time`", "`update_at`"), ",")
	applicationRowsWithPlaceHolder = strings.Join(stringx.Remove(applicationFieldNames, "`id`", "`updated_at`", "`update_time`", "`create_at`", "`created_at`", "`create_time`", "`update_at`"), "=?,") + "=?"

	cacheApplicationIdPrefix                        = "cache:application:id:"
	cacheApplicationCreateUserEnNameProjectIdPrefix = "cache:application:createUser:enName:projectId:"
)

type (
	applicationModel interface {
		Insert(ctx context.Context, data *Application) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*Application, error)
		FindOneByCreateUserEnNameProjectId(ctx context.Context, createUser int64, enName string, projectId string) (*Application, error)
		Update(ctx context.Context, data *Application) error
		Delete(ctx context.Context, id int64) error
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

func (m *defaultApplicationModel) Delete(ctx context.Context, id int64) error {
	data, err := m.FindOne(ctx, id)
	if err != nil {
		return err
	}

	applicationCreateUserEnNameProjectIdKey := fmt.Sprintf("%s%v:%v:%v", cacheApplicationCreateUserEnNameProjectIdPrefix, data.CreateUser, data.EnName, data.ProjectId)
	applicationIdKey := fmt.Sprintf("%s%v", cacheApplicationIdPrefix, id)
	_, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.ExecCtx(ctx, query, id)
	}, applicationCreateUserEnNameProjectIdKey, applicationIdKey)
	return err
}

func (m *defaultApplicationModel) FindOne(ctx context.Context, id int64) (*Application, error) {
	applicationIdKey := fmt.Sprintf("%s%v", cacheApplicationIdPrefix, id)
	var resp Application
	err := m.QueryRowCtx(ctx, &resp, applicationIdKey, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", applicationRows, m.table)
		return conn.QueryRowCtx(ctx, v, query, id)
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

func (m *defaultApplicationModel) FindOneByCreateUserEnNameProjectId(ctx context.Context, createUser int64, enName string, projectId string) (*Application, error) {
	applicationCreateUserEnNameProjectIdKey := fmt.Sprintf("%s%v:%v:%v", cacheApplicationCreateUserEnNameProjectIdPrefix, createUser, enName, projectId)
	var resp Application
	err := m.QueryRowIndexCtx(ctx, &resp, applicationCreateUserEnNameProjectIdKey, m.formatPrimary, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) (i interface{}, e error) {
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
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultApplicationModel) Insert(ctx context.Context, data *Application) (sql.Result, error) {
	applicationCreateUserEnNameProjectIdKey := fmt.Sprintf("%s%v:%v:%v", cacheApplicationCreateUserEnNameProjectIdPrefix, data.CreateUser, data.EnName, data.ProjectId)
	applicationIdKey := fmt.Sprintf("%s%v", cacheApplicationIdPrefix, data.Id)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", m.table, applicationRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.Id, data.ZnName, data.EnName, data.Ico, data.Info, data.CreateUser, data.DemandIds, data.DocIds, data.JoinUsers, data.JoinGroups, data.ProjectId, data.Remark, data.Rank)
	}, applicationCreateUserEnNameProjectIdKey, applicationIdKey)
	return ret, err
}

func (m *defaultApplicationModel) Update(ctx context.Context, newData *Application) error {
	data, err := m.FindOne(ctx, newData.Id)
	if err != nil {
		return err
	}

	applicationCreateUserEnNameProjectIdKey := fmt.Sprintf("%s%v:%v:%v", cacheApplicationCreateUserEnNameProjectIdPrefix, data.CreateUser, data.EnName, data.ProjectId)
	applicationIdKey := fmt.Sprintf("%s%v", cacheApplicationIdPrefix, data.Id)
	_, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, applicationRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, newData.ZnName, newData.EnName, newData.Ico, newData.Info, newData.CreateUser, newData.DemandIds, newData.DocIds, newData.JoinUsers, newData.JoinGroups, newData.ProjectId, newData.Remark, newData.Rank, newData.State, newData.Id)
	}, applicationCreateUserEnNameProjectIdKey, applicationIdKey)
	return err
}

func (m *defaultApplicationModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheApplicationIdPrefix, primary)
}

func (m *defaultApplicationModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", applicationRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultApplicationModel) tableName() string {
	return m.table
}
