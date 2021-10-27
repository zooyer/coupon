package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/zooyer/coupon/common/errno"
	"github.com/zooyer/coupon/dto"
	"net/http"
)

type controller struct{}

type Controller controller

func (c controller) Bind(ctx *gin.Context, v interface{}) (err error) {
	if err = ctx.Bind(v); err != nil {
		return errno.New(errno.InvalidRequest, err)
	}
	return
}

func (c controller) Response(ctx *gin.Context, err error, data interface{}) {
	var resp dto.Response
	if err != nil {
		eno, ok := err.(errno.Error)
		if !ok {
			eno = errno.New(errno.UnknownError, err)
		}
		// 记录错误日志，上报metric
		eno.Record(ctx).Metric()
		// 直接返回错误码
		resp.Code = eno.Errno()
	} else {
		resp.Data = data
	}
	resp.Msg = errno.Msg(resp.Code)
	ctx.JSON(http.StatusOK, resp)
}

func (c controller) Mount(router gin.IRouter, method string, path string, handler ...gin.HandlerFunc) {
	router.Handle(method, path, handler...)
}
