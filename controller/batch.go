package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/zooyer/coupon/dto"
	"github.com/zooyer/coupon/service"
)

type batch struct {
	controller
}

var Batch batch

func (b batch) Detail(ctx *gin.Context) {
	var (
		err  error
		req  dto.BatchDetailReq
		resp *dto.BatchResp
	)

	defer func() { b.Response(ctx, err, resp) }()

	if err = b.Bind(ctx, &req); err != nil {
		return
	}

	if resp, err = service.Batch.Detail(ctx, req.BatchID); err != nil {
		return
	}
}

func (b batch) Create(ctx *gin.Context) {
	var (
		err  error
		req  dto.BatchCreateReq
		resp *dto.BatchResp
	)

	defer func() { b.Response(ctx, err, resp) }()

	if err = b.Bind(ctx, &req); err != nil {
		return
	}

	if resp, err = service.Batch.Create(ctx, req); err != nil {
		return
	}
}

func (b batch) Update(ctx *gin.Context) {
	var (
		err  error
		req  dto.BatchUpdateReq
		resp *dto.BatchResp
	)

	defer func() { b.Response(ctx, err, resp) }()

	if err = b.Bind(ctx, &req); err != nil {
		return
	}

	if resp, err = service.Batch.Update(ctx, req); err != nil {
		return
	}
}

func (b batch) Delete(ctx *gin.Context) {
	var (
		err  error
		req  dto.BatchDeleteReq
		resp *dto.BatchResp
	)

	defer func() { b.Response(ctx, err, resp) }()

	if err = b.Bind(ctx, &req); err != nil {
		return
	}

	if resp, err = service.Batch.Delete(ctx, req); err != nil {
		return
	}
}
