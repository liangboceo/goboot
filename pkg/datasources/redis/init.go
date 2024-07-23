package redis

import (
	"github.com/liangboceo/dependencyinjection"
	"github.com/liangboceo/yuanboot/abstractions"
	"github.com/liangboceo/yuanboot/abstractions/health"
	"github.com/liangboceo/yuanboot/pkg/cache/redis"
)

func init() {
	abstractions.RegisterConfigurationProcessor(
		func(config abstractions.IConfiguration, serviceCollection *dependencyinjection.ServiceCollection) {
			serviceCollection.AddSingletonByImplementsAndName("redis-master", NewRedis, new(abstractions.IDataSource))
			serviceCollection.AddTransientByImplements(NewRedisClient, new(redis.IClient))
			serviceCollection.AddTransientByImplements(NewRedisHealthIndicator, new(health.Indicator))
		})
}
