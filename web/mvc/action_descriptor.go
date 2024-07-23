package mvc

import "github.com/liangboceo/reflectx"

type ActionDescriptor struct {
	ActionName       string
	ActionMethod     string
	MethodInfo       reflectx.MethodInfo
	IsAttributeRoute bool
	Route            *RouteAttribute
}
