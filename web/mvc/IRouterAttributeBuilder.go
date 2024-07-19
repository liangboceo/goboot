package mvc

import (
	"github.com/liangboceo/yuanboot/web/context"
)

type IRouterAttributeBuilder interface {
	Match(ctx *context.HttpContext, pathComponents []string) (string, bool)
	Insert(method, path string, handler func(ctx *context.HttpContext))
}
