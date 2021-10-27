package model

import (
	"context"
	"github.com/zooyer/coupon/common"
)

func Init() {
	db := common.DB(context.Background())
	db.AutoMigrate(&Activity{})
	db.AutoMigrate(&Batch{})
}
