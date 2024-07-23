package consul

import (
	"github.com/liangboceo/dependencyinjection"
	"github.com/liangboceo/yuanboot/abstractions/servicediscovery"
	sd "github.com/liangboceo/yuanboot/pkg/servicediscovery"
)

func UseServiceDiscovery(serviceCollection *dependencyinjection.ServiceCollection) {
	sd.UseGeneralServiceDiscovery(serviceCollection)
	serviceCollection.AddSingletonByImplements(NewServerDiscoveryWithDI, new(servicediscovery.IServiceDiscovery))
	//serviceCollection.AddSingletonByImplements(sd.NewClient, new(servicediscovery.IServiceDiscoveryClient))

}
