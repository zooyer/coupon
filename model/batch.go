package model

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/zooyer/coupon/common/consts"
	"sync"
)

type Batch struct {
	Model
	ActivityID   int64       `gorm:"column:activity_id"`   // 活动ID
	ActivityName string      `gorm:"column:activity_name"` // 活动名称
	BatchID      int64       `gorm:"column:batch_id"`      // 批次ID
	BatchName    string      `gorm:"column:batch_name"`    // 批次名称
	BeginTime    string      `gorm:"column:begin_time"`    // 开始时间
	EndTime      string      `gorm:"column:end_time"`      // 结束时间
	BudgetTotal  int64       `gorm:"column:budget_total"`  // 活动预算
	BudgetUsed   int64       `gorm:"column:budget_used"`   // 预算使用
	BindType     int         `gorm:"column:bind_type"`     // 绑定类型
	BatchRule    interface{} `gorm:"column:batch_rule"`    // 批次规则
	BatchPrice   interface{} `gorm:"column:batch_price"`   // 批次定价
	BatchCapping int         `gorm:"column:batch_capping"` // 批次封顶
	WhiteList    int         `gorm:"column:white_list"`    // 白名单
	Remark       string      `gorm:"column:remark"`        // 批次备注
}

func (Batch) TableName() string {
	return consts.TableNameBatch
}

type RulerArgs struct {
	Key  string `json:"key"`
	Args int    `json:"args"`
}

type Ruler interface {
	Verify(ctx context.Context, args RulerArgs) (valid bool, reason string)
}

type RulerFunc func(data []byte) (ruler Ruler, err error)

type RuleEngine interface {
}

type Pricer interface {
}

type PriceEngine interface {
}

var rulerBatchIndex = [][]string{
	{"city", "county"},
	{"abc", "tag", "ice"},
}

var rulerIndex = make(map[string]int)

var rulers = make(map[string]RulerFunc)

func init() {
	for i, keys := range rulerBatchIndex {
		for _, key := range keys {
			rulerIndex[key] = i
		}
	}
}

const demo = `{"city": {"city": [1,2,3], "county": [11,22,33]}}`

type CityRule struct {
	City   []int `json:"city"`
	County []int `json:"county"`
}

func (c CityRule) OK() bool {
	return false
}

type batchArgs struct {
	Key  string          `json:"key"`
	Rule json.RawMessage `json:"rule"`
}

func init() {
	var rules = make(map[string]json.RawMessage)
	if err := json.Unmarshal([]byte(demo), &rules); err != nil {
		panic(err)
	}

	var batch = make([][]batchArgs, len(rulerBatchIndex))
	var other []batchArgs

	for key, rule := range rules {
		var args = batchArgs{
			Key:  key,
			Rule: rule,
		}

		if index, exists := rulerIndex[key]; exists {
			batch[index] = append(batch[index], args)
			continue
		}

		other = append(other, args)
	}

	if len(other) > 0 {
		batch = append(batch, other)
	}

	for _, batch := range batch {
		if len(batch) == 0 {
			continue
		}

		var wg sync.WaitGroup
		for _, args := range batch {
			wg.Add(1)
			go func(args batchArgs) {
				defer wg.Done()
				var fn = rulers[args.Key]
				if fn == nil {
					//fn = default
				}
				ruler, err := fn(args.Rule)
				if err != nil {
					panic(err)
					return
				}
				ok, reason := ruler.Verify(context.Background(), RulerArgs{
					Key:  args.Key,
					Args: 0,
				})
				fmt.Println(args.Key, "is", ok, "reason:", reason)
			}(args)
		}
	}
}
