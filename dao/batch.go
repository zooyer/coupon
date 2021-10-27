package dao

import (
	"context"
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/zooyer/coupon/common/consts"
	"github.com/zooyer/coupon/model"
)

type batch struct {
	dao
}

var Batch = batch{
	dao: dao{
		model: new(model.Batch),
		table: consts.TableNameBatch,
	},
}

func (b batch) Get(ctx context.Context, id int64) (*model.Batch, error) {
	var (
		err   error
		batch = &model.Batch{
			BatchID: id,
		}
	)
	if err = b.db(ctx).Where(batch).First(batch).Error; err != nil {
		return nil, err
	}
	return batch, nil
}

func (b batch) Create(ctx context.Context, batch model.Batch) (*model.Batch, error) {
	var err error
	if err = Activity.Transaction(ctx, func(db *gorm.DB) (err error) {
		var activity = &model.Activity{
			ActivityID: batch.ActivityID,
		}
		if err = db.Where(activity).First(activity).Error; err != nil {
			return
		}
		if activity.Used >= activity.Budget {
			return fmt.Errorf("activity %d, used %d >= budget %d", activity.ActivityID, activity.Used, activity.Budget)
		}
		if activity.Used+batch.Budget > activity.Budget {
			return fmt.Errorf(
				"activity %d, used %d + %d = %d > budget %d",
				activity.ActivityID, activity.Used, batch.Budget, activity.Used+batch.Budget, activity.Budget,
			)
		}
		if err = db.Where(activity).Update(map[string]interface{}{"used": activity.Used + batch.Budget}).Error; err != nil {
			return
		}
		activity.Used += batch.Budget
		if err = db.Table(b.table).Create(&batch).Error; err != nil {
			return
		}
		return
	}); err != nil {
		return nil, err
	}

	return &batch, nil
}

func (b batch) Update(ctx context.Context, id int64, update map[string]interface{}) (err error) {
	var batch = &model.Batch{
		BatchID: id,
	}
	if err = b.db(ctx).Where(batch).Update(update).Error; err != nil {
		return
	}
	return
}

func (b batch) Delete(ctx context.Context, id int64) (err error) {
	var batch = &model.Batch{
		BatchID: id,
	}
	if err = b.db(ctx).Where(batch).Delete(batch).Error; err != nil {
		return
	}
	return
}
