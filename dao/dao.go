package dao

import (
	"context"
	"github.com/jinzhu/gorm"
	"github.com/zooyer/coupon/common"
)

type dao struct {
	model interface{}
	table string
}

func (d dao) db(ctx context.Context) *gorm.DB {
	return common.DB(ctx).Table(d.table)
}

func (d dao) Transaction(ctx context.Context, fn func(db *gorm.DB) (err error)) (err error) {
	db := d.db(ctx).Begin()
	defer func() {
		if err != nil {
			db.Rollback()
		} else {
			err = db.Commit().Error
		}
	}()
	return fn(db)
}
