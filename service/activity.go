package service

import (
	"context"
	"github.com/zooyer/coupon/rpc/iden"

	"github.com/zooyer/coupon/dao"
	"github.com/zooyer/coupon/dto"
	"github.com/zooyer/coupon/model"
)

type activity struct{}

var Activity activity

func (a activity) toResp(activity model.Activity) *dto.ActivityResp {
	return &dto.ActivityResp{
		ActivityID: activity.ActivityID,
		Balance:    activity.Budget - activity.Used,
		ActivityCreateReq: dto.ActivityCreateReq{
			Name:      activity.Name,
			BeginTime: activity.BeginTime,
			EndTime:   activity.EndTime,
			Budget:    activity.Budget,
			Remark:    activity.Remark,
		},
	}
}

func (a activity) Detail(ctx context.Context, id int64) (resp *dto.ActivityResp, err error) {
	activity, err := dao.Activity.Get(ctx, id)
	if err != nil {
		return
	}

	return a.toResp(*activity), nil
}

func (a activity) Create(ctx context.Context, req dto.ActivityCreateReq) (resp *dto.ActivityResp, err error) {
	id, err := iden.GenID()
	if err != nil {
		return
	}

	var activity = &model.Activity{
		ActivityID: id,
		Name:       req.Name,
		BeginTime:  req.BeginTime,
		EndTime:    req.EndTime,
		Budget:     req.Budget,
		Remark:     req.Remark,
	}

	if activity, err = dao.Activity.Create(ctx, *activity); err != nil {
		return
	}

	return a.toResp(*activity), nil
}

func (a activity) Update(ctx context.Context, req dto.ActivityUpdateReq) (resp *dto.ActivityResp, err error) {
	var update = make(map[string]interface{})

	if req.Budget > 0 {
		update["budget"] = req.Budget
	}
	if req.EndTime != "" {
		update["end_time"] = req.EndTime
	}
	if req.Remark != "" {
		update["remark"] = req.Remark
	}

	if len(update) > 0 {
		if err = dao.Activity.Update(ctx, req.ActivityID, update); err != nil {
			return
		}
	}

	if resp, err = a.Detail(ctx, req.ActivityID); err != nil {
		return
	}

	return
}

func (a activity) Delete(ctx context.Context, req dto.ActivityDeleteReq) (resp *dto.ActivityResp, err error) {
	if resp, err = a.Detail(ctx, req.ActivityID); err != nil {
		return
	}

	if err = dao.Activity.Delete(ctx, req.ActivityID); err != nil {
		return
	}

	return
}
