package model

import "github.com/zooyer/coupon/common/consts"

type Activity struct {
	Model
	ActivityID int64  `gorm:"column:activity_id"`
	Name       string `gorm:"column:name"`       // 活动名称
	BeginTime  string `gorm:"column:begin_time"` // 开始时间
	EndTime    string `gorm:"column:end_time"`   // 结束时间
	Budget     int64  `gorm:"column:budget"`     // 活动预算
	Used       int64  `gorm:"column:used"`       // 预算使用
	Remark     string `gorm:"column:remark"`     // 活动备注
}

func (Activity) TableName() string {
	return consts.TableNameActivity
}
