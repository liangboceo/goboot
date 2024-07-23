package httpclient

import (
	"github.com/liangboceo/dependencyinjection"
	"github.com/liangboceo/yuanboot/abstractions"
)

func init() {
	abstractions.RegisterConfigurationProcessor(func(config abstractions.IConfiguration, serviceCollection *dependencyinjection.ServiceCollection) {
		serviceCollection.AddSingleton(NewFactory)
	})
}
