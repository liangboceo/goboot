package mysql

import (
	"github.com/liangboceo/yuanboot/abstractions"
	"github.com/liangboceo/yuanboot/abstractions/health"
	"github.com/yoyofxteam/dependencyinjection"
)

func init() {
	abstractions.RegisterConfigurationProcessor(
		func(config abstractions.IConfiguration, serviceCollection *dependencyinjection.ServiceCollection) {
			serviceCollection.AddSingletonByImplementsAndName("mysql-master", NewMysqlDataSource, new(abstractions.IDataSource))
			serviceCollection.AddSingleton(NewGormDb)
			serviceCollection.AddTransientByImplements(NewMysqlHealthIndicator, new(health.Indicator))
		})

}
