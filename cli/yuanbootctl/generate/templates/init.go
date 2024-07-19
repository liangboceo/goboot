package templates

import (
	"github.com/liangboceo/yuanboot/cli/yuanbootctl/generate/templates/console"
	"github.com/liangboceo/yuanboot/cli/yuanbootctl/generate/templates/grpc"
	"github.com/liangboceo/yuanboot/cli/yuanbootctl/generate/templates/mvc"
	"github.com/liangboceo/yuanboot/cli/yuanbootctl/generate/templates/webapi"
	"github.com/liangboceo/yuanboot/cli/yuanbootctl/generate/templates/xxl_job"
)

func init() {
	registerProject("console", console.Project)
	registerProject("webapi", webapi.Project)
	registerProject("mvc", mvc.Project)
	registerProject("grpc", grpc.Project)
	registerProject("xxl-job", xxl_job.Project)
}
