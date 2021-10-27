package dto

import "github.com/zooyer/jsons"

type BatchResp struct {
	BatchID int64 `json:"batch_id"` // 批次ID
	BatchCreateReq
	Balance int64 `json:"balance"` // 批次余额
}

type BatchDetailReq struct {
	BatchID int64 `form:"batch_id" binding:"required"` // 批次ID
}

type BatchCreateReq struct {
	ActivityID int64  `json:"activity_id" binding:"required"` // 活动ID
	Name       string `json:"name"`                           // 批次名称
	BeginTime  string `json:"begin_time" binding:"required"`  // 开始时间
	EndTime    string `json:"end_time" binding:"required"`    // 截止时间
	Budget     int64  `json:"budget" binding:"required"`      // 批次预算

	Type     string      `json:"type"`     // 批次类型
	Amount   int64       `json:"amount"`   // 批次面额
	Discount int64       `json:"discount"` // 批次折扣
	Capping  int64       `json:"capping"`  // 批次封顶
	Cond     jsons.Array `json:"cond"`     // 批次条件
	Rule     jsons.Value `json:"rule"`     // 批次规则

	Remark string `json:"remark"` // 批次备注
}

type BatchUpdateReq struct {
	BatchID int64  `json:"batch_id"` // 批次ID
	Name    string `json:"name"`     // 批次名称
	EndTime string `json:"end_time"` // 截止时间
	Budget  int64  `json:"budget"`   // 批次预算
	Remark  string `json:"remark"`   // 批次备注
}

type BatchDeleteReq struct {
	BatchID int64 `json:"batch_id" binding:"required"` // 批次ID
}
