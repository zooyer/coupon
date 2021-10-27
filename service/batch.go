package service

import (
	"context"
	"github.com/zooyer/coupon/dao"
	"github.com/zooyer/coupon/dto"
	"github.com/zooyer/coupon/model"
	"github.com/zooyer/coupon/rpc/iden"
)

type batch struct{}

var Batch batch

func (b batch) toResp(batch model.Batch) *dto.BatchResp {
	return &dto.BatchResp{
		BatchID: batch.BatchID,
		BatchCreateReq: dto.BatchCreateReq{
			ActivityID: batch.ActivityID,
			Name:       batch.Name,
			BeginTime:  batch.BeginTime,
			EndTime:    batch.EndTime,
			Budget:     batch.Budget,
			Remark:     batch.Remark,
		},
		Balance: batch.Budget - batch.Used,
	}
}

func (b batch) Detail(ctx context.Context, id int64) (resp *dto.BatchResp, err error) {
	batch, err := dao.Batch.Get(ctx, id)
	if err != nil {
		return
	}

	return b.toResp(*batch), nil
}

func (b batch) Create(ctx context.Context, req dto.BatchCreateReq) (*dto.BatchResp, error) {
	id, err := iden.GenID()
	if err != nil {
		return nil, err
	}
	var batch = &model.Batch{
		ActivityID: req.ActivityID,
		BatchID:    id,
		Name:       req.Name,
		BeginTime:  req.BeginTime,
		EndTime:    req.EndTime,
		Budget:     req.Budget,
		Used:       0,
		Remark:     req.Remark,
	}
	if batch, err = dao.Batch.Create(ctx, *batch); err != nil {
		return nil, err
	}

	return b.toResp(*batch), nil
}

func (b batch) Update(ctx context.Context, req dto.BatchUpdateReq) (resp *dto.BatchResp, err error) {
	var update = make(map[string]interface{})

	if req.Name != "" {
		update["name"] = req.Name
	}
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
		if err = dao.Batch.Update(ctx, req.BatchID, update); err != nil {
			return
		}
	}

	if resp, err = b.Detail(ctx, req.BatchID); err != nil {
		return
	}

	return

}

func (b batch) Delete(ctx context.Context, req dto.BatchDeleteReq) (resp *dto.BatchResp, err error) {
	if resp, err = b.Detail(ctx, req.BatchID); err != nil {
		return
	}

	if err = dao.Batch.Delete(ctx, req.BatchID); err != nil {
		return
	}

	return
}

func (b batch) Get(ctx context.Context) {

}

func (b batch) List(ctx context.Context) {

}
