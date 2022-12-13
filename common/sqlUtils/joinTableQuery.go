package sqlUtils

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

type (
	JoinTableQuery interface {
		FindUserJoinUserAuthsByKey(keyType string, key string) (*UserJoinUserAuths, error)
		FindApplicationJoinConfig(ctx context.Context, create_user, application_id int64, version string, resp interface{}) (err error)
	}
	defaultJoinTableQuery struct {
		sqlc.CachedConn
	}

	UserJoinUserAuths struct {
		UserId         int64  `db:"user_id"`
		Nickname       string `db:"nickname"`
		RealName       string `db:"realName"`
		Password       string `db:"password"`
		LoginSalt      string `db:"login_salt"`
		RegisterDevice string `db:"register_device"`
		Sex            int64  `db:"sex"`
		Ico            string `db:"ico"`
		IdentityType   string `db:"identity_type"`
		Identifier     string `db:"identifier"`
		UStatus        int64  `db:"uStatus"`
		//CreateTime     time.Time `db:"create_time"`
		//UpdateTime     time.Time `db:"update_time"`
	}
)

func NewJoinTableQuery(conn sqlx.SqlConn, c cache.CacheConf) JoinTableQuery {
	return &defaultJoinTableQuery{
		CachedConn: sqlc.NewConn(conn, c),
	}
}

func (m *defaultJoinTableQuery) FindUserJoinUserAuthsByKey(keyType string, key string) (*UserJoinUserAuths, error) {
	var resp UserJoinUserAuths
	UserJoinUserAuthsRows := ""
	for _, str := range builder.RawFieldNames(&resp) {
		if str != "`uStatus`" {
			if stringx.Contains([]string{"`user_id`", "`uStatus`", "`identity_type`", "`identifier`"}, str) {
				UserJoinUserAuthsRows += fmt.Sprintf("ua.%s,", str[1:len(str)-1])
			} else {
				UserJoinUserAuthsRows += fmt.Sprintf("u.%s,", str[1:len(str)-1])
			}
		}
	}
	//UserJoinUserAuthsRows := strings.Join(stringx.Remove(builder.RawFieldNames(&resp), "`user_id`", "`uStatus`", "`identity_type`", "`identifier`"), ",")
	//UserJoinUserAuthsRows = strings.ReplaceAll(UserJoinUserAuthsRows, ",`", ",`u.")
	query := fmt.Sprintf("select %s,u.status uStatus from user u inner join user_auths ua on u.id=ua.user_id where `%s` = ? limit 1", UserJoinUserAuthsRows[:len(UserJoinUserAuthsRows)-1], keyType)
	err := m.QueryRowNoCache(&resp, query, key)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultJoinTableQuery) FindApplicationJoinConfig(ctx context.Context, create_user, application_id int64, version string, resp interface{}) (err error) {
	applicationJoinConfigKey := fmt.Sprintf("cache:userId:appId:version:%s:%s:%s", create_user, application_id, version)
	err = m.QueryRowCtx(ctx, resp, applicationJoinConfigKey, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
		query := "select * from config where FIND_IN_SET(`key`, (select `config_ids` from `application_info` where create_user = ? and application_id = ? and version = '?') ) and `state` > 0 and `user_id` = ?;"
		return conn.QueryRowCtx(ctx, resp, query, create_user, application_id, version, create_user)
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
