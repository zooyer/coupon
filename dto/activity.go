package dto

type ActivityDetailReq struct {
	ActivityID int64 `form:"activity_id" binding:"required"` // 活动ID
}

type ActivityCreateReq struct {
	Name      string `json:"name" binding:"required"`       // 活动名称
	BeginTime string `json:"begin_time" binding:"required"` // 开始时间
	EndTime   string `json:"end_time" binding:"required"`   // 结束时间
	Budget    int64  `json:"budget" binding:"required"`     // 活动预算
	Remark    string `json:"remark"`                        // 活动备注
}

type ActivityUpdateReq struct {
	ActivityID int64  `json:"activity_id" binding:"required"` // 活动ID
	Budget     int64  `json:"budget"`                         // 活动预算
	EndTime    string `json:"end_time"`                       // 结束时间
	Remark     string `json:"remark"`                         // 活动备注
}

type ActivityDeleteReq struct {
	ActivityID int64 `json:"activity_id" binding:"required"` // 活动ID
}

type ActivityResp struct {
	ActivityID int64 `json:"activity_id"` // 活动ID
	ActivityCreateReq
	Balance int64 `json:"balance"` // 活动余额
}
