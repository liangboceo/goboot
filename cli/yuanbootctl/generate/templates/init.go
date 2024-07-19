package templates

import (
	"yuanbootctl/generate/templates/console"
	"yuanbootctl/generate/templates/grpc"
	"yuanbootctl/generate/templates/mvc"
	"yuanbootctl/generate/templates/webapi"
	"yuanbootctl/generate/templates/xxl_job"
)

func init() {
	registerProject("console", console.Project)
	registerProject("webapi", webapi.Project)
	registerProject("mvc", mvc.Project)
	registerProject("grpc", grpc.Project)
	registerProject("xxl-job", xxl_job.Project)
}
