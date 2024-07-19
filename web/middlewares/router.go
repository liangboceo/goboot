package middlewares

import (
	"github.com/liangboceo/yuanboot/web/context"
	"github.com/liangboceo/yuanboot/web/router"
	"strings"
)

//var ReqFuncMap = make(map[string]func(ctx * yuanboot.HttpContext))

type RouterMiddleware struct {
	RouterBuilder router.IRouterBuilder
}

func NewRouter(builder router.IRouterBuilder) *RouterMiddleware {
	return &RouterMiddleware{RouterBuilder: builder}
}

func (router *RouterMiddleware) Inovke(ctx *context.HttpContext, next func(ctx *context.HttpContext)) {
	var handler func(ctx *context.HttpContext)
	handler = router.RouterBuilder.Search(ctx, strings.Split(ctx.Input.Request.URL.Path, "/")[1:], ctx.Input.RouterData)
	if handler != nil {
		handler(ctx)
	}
	next(ctx)

}
