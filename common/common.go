package common

import (
	"context"

	"github.com/garyburd/redigo/redis"
	"github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/mattn/go-sqlite3"
	"github.com/zooyer/coupon/common/config"
)

var (
	_ mysql.MySQLDriver
	_ sqlite3.SQLiteDriver
)

var db *gorm.DB

var rds redis.Conn

func InitDB() {
	var err error
	if db, err = gorm.Open(config.DB.Dialect, config.DB.Args); err != nil {
		panic(err)
	}
}

func InitRedis() {
	var err error
	if rds, err = redis.Dial("tcp", config.Redis.Addr); err != nil {
		panic(err)
	}
}

func DB(ctx context.Context) *gorm.DB {
	if config.IsDebug() {
		return db.Set("ctx", ctx).Debug()
	}
	return db.Set("ctx", ctx).Debug()
}

func Redis(ctx context.Context) redis.Conn {
	return rds
}
