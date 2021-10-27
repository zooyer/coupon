package dao

import (
	"context"
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/zooyer/coupon/common/consts"
	"github.com/zooyer/coupon/model"
)

type activity struct {
	dao
}

var Activity = activity{
	dao: dao{
		model: new(model.Activity),
		table: consts.TableNameActivity,
	},
}

func (a activity) Get(ctx context.Context, id int64) (*model.Activity, error) {
	var (
		err      error
		activity = &model.Activity{
			ActivityID: id,
		}
	)
	if err = a.db(ctx).Where(activity).First(activity).Error; err != nil {
		return nil, err
	}
	return activity, nil
}

func (a activity) Create(ctx context.Context, activity model.Activity) (*model.Activity, error) {
	var err error
	if err = a.db(ctx).Create(&activity).Error; err != nil {
		return nil, err
	}
	return &activity, nil
}

func (a activity) Update(ctx context.Context, id int64, update map[string]interface{}) (err error) {
	var activity = &model.Activity{
		ActivityID: id,
	}
	if err = a.db(ctx).Where(activity).Update(update).Error; err != nil {
		return
	}
	return
}

func (a activity) Delete(ctx context.Context, id int64) (err error) {
	var activity = &model.Activity{
		ActivityID: id,
	}
	if err = a.db(ctx).Where(activity).Delete(activity).Error; err != nil {
		return
	}
	return
}

func (a activity) AddBatch(ctx context.Context, id, used int64) (activity *model.Activity, err error) {
	if err = a.Transaction(ctx, func(db *gorm.DB) (err error) {
		activity = &model.Activity{
			ActivityID: id,
		}
		if err = a.db(ctx).Where(activity).First(activity).Error; err != nil {
			return
		}
		if activity.Used >= activity.Budget {
			return fmt.Errorf("activity %d, used %d >= budget %d", id, activity.Used, activity.Budget)
		}
		if activity.Used+used > activity.Budget {
			return fmt.Errorf("activity %d, used %d + %d = %d > budget %d", id, activity.Used, used, activity.Used+used, activity.Budget)
		}
		if err = a.db(ctx).Where(activity).Update(map[string]interface{}{"used": activity.Used + used}).Error; err != nil {
			return
		}
		activity.Used += used



		return
	}); err != nil {
		return nil, err
	}
	return
}
