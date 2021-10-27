package config

import (
	"time"

	"github.com/zooyer/micro/conf"
)

var config *conf.Config

var env string // debug/test/sim/prod

var Gin struct {
	Mode string
	Addr string
}

var Coupon struct {
	Addr string
}

var DB struct {
	Dialect string
	Args    string
}

var Redis struct {
	Addr     string
	Password string
	Database int
}

var binding = map[string]interface{}{
	"env":   &env,
	"gin":   &Gin,
	"db":    &DB,
	"redis": &Redis,
	"count": &Coupon,
}

func Init(file string) {
	var err error

	if config, err = conf.Init(file, time.Second*10); err != nil {
		panic(err)
	}

	var bind = func(key string, v interface{}) {
		if err = config.Bind(key, v, true); err != nil {
			panic(err)
		}
	}

	for key, addr := range binding {
		bind(key, addr)
	}

	return
}
