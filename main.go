package main

import (
	"flag"
	"github.com/gin-gonic/gin"
	"github.com/zooyer/coupon/common"
	"github.com/zooyer/coupon/common/config"
	"github.com/zooyer/coupon/controller"
	"github.com/zooyer/coupon/model"
)

var filename = flag.String("config", "conf/conf.prod.hna.yaml", "config file.")

func init() {
	flag.Parse()

	config.Init(*filename)

	common.InitDB()
	common.InitRedis()

	model.Init()
}

func main() {
	engine := gin.New()

	v1 := engine.Group("/coupon/v1")
	activity := v1.Group("/activity")
	{
		activity.GET("/detail", controller.Activity.Detail)  // 活动详情
		activity.POST("/create", controller.Activity.Create) // 创建活动
		activity.POST("/update", controller.Activity.Update) // 更新活动
		activity.POST("/delete", controller.Activity.Delete) // 删除活动
	}
	batch := v1.Group("/batch")
	{
		batch.GET("/detail", controller.Batch.Detail)  // 批次详情
		batch.POST("/create", controller.Batch.Create) // 创建批次
		batch.POST("/update", controller.Batch.Update) // 更新批次
		batch.POST("/delete", controller.Batch.Delete) // 删除批次
	}
	coupon := v1.Group("/coupon")
	{
		coupon.GET("/detail", nil)
		coupon.POST("/bind", nil)
		coupon.POST("/unbind", nil)
		coupon.POST("/freeze", nil)
		coupon.POST("/unfreeze", nil)
		coupon.POST("/speed", nil)
	}

	engine.Run()
}
