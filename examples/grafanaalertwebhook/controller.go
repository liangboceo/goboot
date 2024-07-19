package main

import (
	"github.com/liangboceo/yuanboot/abstractions"
	"github.com/liangboceo/yuanboot/web/context"
	"grafanaalertwebhook/wechatrequests"
)

func PostAlert(ctx *context.HttpContext) {
	var request wechatrequests.GrafanaAlertRequest
	_ = ctx.Bind(&request)
	var config abstractions.IConfiguration
	_ = ctx.RequiredServices.GetService(&config)

	ctx.JSON(200, context.H{
		"Message": wechatrequests.SendTxtMessage(request, config),
	})
}
