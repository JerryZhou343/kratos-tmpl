package conf

import (
	"github.com/bilibili/kratos/pkg/cache/redis"
	"github.com/bilibili/kratos/pkg/conf/paladin"
	"github.com/bilibili/kratos/pkg/database/sql"
	"github.com/bilibili/kratos/pkg/log"
)

// Config .
type Config struct {
	Redis *redis.Config
	SQL   *sql.Config
}

func NewConf() (conf *Config, err error) {
	var (
		sqlct paladin.TOML
	)
	conf = &Config{}

	if err = paladin.Get("db.toml").Unmarshal(&sqlct); err != nil {
		log.Error("%+v", err)
		return
	}

	sql := sql.Config{}
	err = sqlct.Get("Client").UnmarshalTOML(&sql)
	if err != nil {
		log.Error("%+v", err)
		return
	}
	conf.SQL = &sql
	return
}
