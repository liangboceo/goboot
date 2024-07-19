package httpclient

import (
	"github.com/liangboceo/yuanboot/abstractions"
	"github.com/yoyofxteam/dependencyinjection"
)

func init() {
	abstractions.RegisterConfigurationProcessor(func(config abstractions.IConfiguration, serviceCollection *dependencyinjection.ServiceCollection) {
		serviceCollection.AddSingleton(NewFactory)
	})
}
