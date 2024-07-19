package main

import (
	"github.com/liangboceo/yuanboot/abstractions"
	"github.com/liangboceo/yuanboot/abstractions/hosting"
	"github.com/liangboceo/yuanboot/console"
	"github.com/liangboceo/yuanboot/pkg/servicediscovery/nacos"
	"github.com/yoyofxteam/dependencyinjection"
)

func main() {
	configuration := abstractions.NewConfigurationBuilder().
		AddEnvironment().
		AddYamlFile("config").Build()

	console.NewHostBuilder().
		UseConfiguration(configuration).
		ConfigureServices(func(collection *dependencyinjection.ServiceCollection) {
			hosting.AddHostService(collection, NewClientService)
			collection.AddSingleton(NewHelloworldApi)
			//register sd for nacos
			nacos.UseServiceDiscovery(collection)
		}).
		Build().Run()
}
