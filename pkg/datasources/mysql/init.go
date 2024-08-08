package mysql

import (
	"github.com/liangboceo/dependencyinjection"
	"github.com/liangboceo/yuanboot/abstractions"
	"github.com/liangboceo/yuanboot/abstractions/health"
	"gorm.io/gorm"
)

func init() {
	abstractions.RegisterConfigurationProcessor(
		func(config abstractions.IConfiguration, serviceCollection *dependencyinjection.ServiceCollection) {
			serviceCollection.AddSingletonByImplementsAndName("mysql-master", NewMysqlDataSource, new(abstractions.IDataSource))
			serviceCollection.AddSingletonByImplements(NewGormDb, new(gorm.DB))
			serviceCollection.AddTransientByImplements(NewMysqlHealthIndicator, new(health.Indicator))
		})

}
