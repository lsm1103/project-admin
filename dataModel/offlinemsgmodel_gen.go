// Code generated by goctl. DO NOT EDIT!

package dataModel

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
	offlineMsgFieldNames          = builder.RawFieldNames(&OfflineMsg{})
	offlineMsgRows                = strings.Join(offlineMsgFieldNames, ",")
	offlineMsgRowsExpectAutoSet   = strings.Join(stringx.Remove(offlineMsgFieldNames, "`create_at`", "`created_at`", "`create_time`", "`update_at`", "`updated_at`", "`update_time`", "`state`"), ",")
	offlineMsgRowsWithPlaceHolder = strings.Join(stringx.Remove(offlineMsgFieldNames, "`id`", "`create_at`", "`created_at`", "`create_time`", "`update_at`", "`updated_at`", "`update_time`"), "=?,") + "=?"
	offlineMsgListRows            = strings.Join(builder.RawFieldNames(&OfflineMsg{}), ",")

	cacheOfflineMsgIdPrefix                               = "cache:offlineMsg:id:"
	cacheOfflineMsgUserIdDeviceIdObjectTypeObjectIdPrefix = "cache:offlineMsg:userId:deviceId:objectType:objectId:"
)

type (
	offlineMsgModel interface {
		Insert(ctx context.Context, session sqlx.Session, data *OfflineMsg) (sql.Result, error)
		FindOne(ctx context.Context, session sqlx.Session, id int64, resp interface{}) (err error)
		FindOneByUserIdDeviceIdObjectTypeObjectId(ctx context.Context, userId int64, deviceId string, objectType int64, objectId int64) (*OfflineMsg, error)
		Update(ctx context.Context, session sqlx.Session, data *OfflineMsg) error
		Delete(ctx context.Context, session sqlx.Session, id int64) error
	}

	defaultOfflineMsgModel struct {
		sqlc.CachedConn
		table string
	}

	OfflineMsg struct {
		Id         int64     `db:"id"`           // 自增主键
		UserId     int64     `db:"user_id"`      // 用户id
		DeviceId   string    `db:"device_id"`    // 设备id
		ObjectType int64     `db:"object_type"`  // 对象类型,1:friend；2：群组
		ObjectId   int64     `db:"object_id"`    // 对象id, friendId/groupId
		LastAckSeq int64     `db:"last_ack_seq"` // 最后确认序列号
		NewestSeq  int64     `db:"newest_seq"`   // 最新的消息序列号
		State      int64     `db:"state"`        // 消息状态：-1撤回，0未处理，1已读
		CreateTime time.Time `db:"create_time"`  // 创建时间
		UpdateTime time.Time `db:"update_time"`  // 更新时间
	}
)

func newOfflineMsgModel(conn sqlx.SqlConn, c cache.CacheConf) *defaultOfflineMsgModel {
	return &defaultOfflineMsgModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`offline_msg`",
	}
}

func (m *defaultOfflineMsgModel) Delete(ctx context.Context, session sqlx.Session, id int64) error {
	data := &OfflineMsg{}
	err := m.FindOne(ctx, session, id, data)
	if err != nil {
		return err
	}
	offlineMsgIdKey := fmt.Sprintf("%s%v", cacheOfflineMsgIdPrefix, id)
	offlineMsgUserIdDeviceIdObjectTypeObjectIdKey := fmt.Sprintf("%s%v:%v:%v:%v", cacheOfflineMsgUserIdDeviceIdObjectTypeObjectIdPrefix, data.UserId, data.DeviceId, data.ObjectType, data.ObjectId)
	_, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		if session != nil {
			return session.ExecCtx(ctx, query, id)
		}
		return conn.ExecCtx(ctx, query, id)
	}, offlineMsgIdKey, offlineMsgUserIdDeviceIdObjectTypeObjectIdKey)
	return err
}

func (m *defaultOfflineMsgModel) FindOne(ctx context.Context, session sqlx.Session, id int64, resp interface{}) (err error) {
	offlineMsgIdKey := fmt.Sprintf("%s%v", cacheOfflineMsgIdPrefix, id)
	err = m.QueryRowCtx(ctx, resp, offlineMsgIdKey, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", offlineMsgRows, m.table)
		if session != nil {
			return session.QueryRowCtx(ctx, v, query, id)
		}
		return conn.QueryRowCtx(ctx, v, query, id)
	})
	switch err {
	case nil:
		return nil
	case sqlc.ErrNotFound:
		return ErrNotFound
	default:
		return err
	}
}

func (m *defaultOfflineMsgModel) FindOneByUserIdDeviceIdObjectTypeObjectId(ctx context.Context, userId int64, deviceId string, objectType int64, objectId int64) (*OfflineMsg, error) {
	offlineMsgUserIdDeviceIdObjectTypeObjectIdKey := fmt.Sprintf("%s%v:%v:%v:%v", cacheOfflineMsgUserIdDeviceIdObjectTypeObjectIdPrefix, userId, deviceId, objectType, objectId)
	var resp OfflineMsg
	err := m.QueryRowIndexCtx(ctx, &resp, offlineMsgUserIdDeviceIdObjectTypeObjectIdKey, m.formatPrimary, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) (i interface{}, e error) {
		query := fmt.Sprintf("select %s from %s where `user_id` = ? and `device_id` = ? and `object_type` = ? and `object_id` = ? limit 1", offlineMsgRows, m.table)
		if err := conn.QueryRowCtx(ctx, &resp, query, userId, deviceId, objectType, objectId); err != nil {
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

func (m *defaultOfflineMsgModel) Insert(ctx context.Context, session sqlx.Session, data *OfflineMsg) (sql.Result, error) {
	offlineMsgIdKey := fmt.Sprintf("%s%v", cacheOfflineMsgIdPrefix, data.Id)
	offlineMsgUserIdDeviceIdObjectTypeObjectIdKey := fmt.Sprintf("%s%v:%v:%v:%v", cacheOfflineMsgUserIdDeviceIdObjectTypeObjectIdPrefix, data.UserId, data.DeviceId, data.ObjectType, data.ObjectId)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?)", m.table, offlineMsgRowsExpectAutoSet)
		if session != nil {
			return session.ExecCtx(ctx, query, data.Id, data.UserId, data.DeviceId, data.ObjectType, data.ObjectId, data.LastAckSeq, data.NewestSeq)
		}
		return conn.ExecCtx(ctx, query, data.Id, data.UserId, data.DeviceId, data.ObjectType, data.ObjectId, data.LastAckSeq, data.NewestSeq)
	}, offlineMsgIdKey, offlineMsgUserIdDeviceIdObjectTypeObjectIdKey)
	return ret, err
}

func (m *defaultOfflineMsgModel) Update(ctx context.Context, session sqlx.Session, newData *OfflineMsg) error {
	data := &OfflineMsg{}
	err := m.FindOne(ctx, session, newData.Id, data)
	if err != nil {
		return err
	}

	offlineMsgIdKey := fmt.Sprintf("%s%v", cacheOfflineMsgIdPrefix, data.Id)
	offlineMsgUserIdDeviceIdObjectTypeObjectIdKey := fmt.Sprintf("%s%v:%v:%v:%v", cacheOfflineMsgUserIdDeviceIdObjectTypeObjectIdPrefix, data.UserId, data.DeviceId, data.ObjectType, data.ObjectId)
	_, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, offlineMsgRowsWithPlaceHolder)
		if session != nil {
			return session.ExecCtx(ctx, query, newData.UserId, newData.DeviceId, newData.ObjectType, newData.ObjectId, newData.LastAckSeq, newData.NewestSeq, newData.State, newData.Id)
		}
		return conn.ExecCtx(ctx, query, newData.UserId, newData.DeviceId, newData.ObjectType, newData.ObjectId, newData.LastAckSeq, newData.NewestSeq, newData.State, newData.Id)
	}, offlineMsgIdKey, offlineMsgUserIdDeviceIdObjectTypeObjectIdKey)
	return err
}

func (m *defaultOfflineMsgModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheOfflineMsgIdPrefix, primary)
}

func (m *defaultOfflineMsgModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", offlineMsgRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultOfflineMsgModel) tableName() string {
	return m.table
}
