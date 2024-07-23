package abstractions

import (
	"github.com/liangboceo/dependencyinjection"
	"github.com/liangboceo/yuanboot/abstractions/hostenv"
)

type HostBuilderContext struct {
	RequestDelegate        interface{}
	ApplicationCycle       *ApplicationLife
	HostingEnvironment     *HostEnvironment
	Configuration          IConfiguration
	HostConfiguration      *hostenv.HostConfig
	ApplicationServicesDef *dependencyinjection.ServiceCollection
	ApplicationServices    dependencyinjection.IServiceProvider
	HostServices           dependencyinjection.IServiceProvider
}
