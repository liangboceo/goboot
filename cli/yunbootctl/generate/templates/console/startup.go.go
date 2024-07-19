package console

const ProjectItem_startup_go = `
package {{.CurrentModelName}}

import (
	"github.com/liangboceo/yuanboot/abstractions"
	"github.com/liangboceo/yuanboot/abstractions/hosting"
	"github.com/yoyofxteam/dependencyinjection"
)

type AppStartup struct {
}

func Startup() abstractions.IStartup {
	return &AppStartup{}
}

func (s *AppStartup) ConfigureServices(collection *dependencyinjection.ServiceCollection) {
	hosting.AddHostService(collection, NewService)
}
`
