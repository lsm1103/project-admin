package config

import (
    {{.authImport}}
	"github.com/zeromicro/go-zero/core/stores/cache"
)

type Config struct {
	rest.RestConf
	{{.auth}}
	{{.jwtTrans}}

	DB struct {
		DataSource string
	}
	Cache cache.CacheConf
}
