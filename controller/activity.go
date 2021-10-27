package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/zooyer/coupon/dto"
	"github.com/zooyer/coupon/service"
)

type activity struct {
	controller
}

var Activity activity

func (a activity) Detail(ctx *gin.Context) {
	var (
		err  error
		req  dto.ActivityDetailReq
		resp *dto.ActivityResp
	)

	defer func() { a.Response(ctx, err, resp) }()

	if err = a.Bind(ctx, &req); err != nil {
		return
	}

	if resp, err = service.Activity.Detail(ctx, req.ActivityID); err != nil {
		return
	}

	return
}

func (a activity) Create(ctx *gin.Context) {
	var (
		err  error
		req  dto.ActivityCreateReq
		resp *dto.ActivityResp
	)

	defer func() { a.Response(ctx, err, resp) }()

	if err = a.Bind(ctx, &req); err != nil {
		return
	}

	if resp, err = service.Activity.Create(ctx, req); err != nil {
		return
	}
}

func (a activity) Update(ctx *gin.Context) {
	var (
		err  error
		req  dto.ActivityUpdateReq
		resp *dto.ActivityResp
	)

	defer func() { a.Response(ctx, err, resp) }()

	if err = a.Bind(ctx, &req); err != nil {
		return
	}

	if resp, err = service.Activity.Update(ctx, req); err != nil {
		return
	}
}

func (a activity) Delete(ctx *gin.Context) {
	var (
		err  error
		req  dto.ActivityDeleteReq
		resp *dto.ActivityResp
	)

	defer func() { a.Response(ctx, err, resp) }()

	if err = a.Bind(ctx, &req); err != nil {
		return
	}

	if resp, err = service.Activity.Delete(ctx, req); err != nil {
		return
	}
}
