package redis

import (
	"github.com/liangboceo/yuanboot/abstractions"
	"github.com/liangboceo/yuanboot/abstractions/health"
	"github.com/liangboceo/yuanboot/pkg/cache/redis"
	"github.com/yoyofxteam/dependencyinjection"
)

func init() {
	abstractions.RegisterConfigurationProcessor(
		func(config abstractions.IConfiguration, serviceCollection *dependencyinjection.ServiceCollection) {
			serviceCollection.AddSingletonByImplementsAndName("redis-master", NewRedis, new(abstractions.IDataSource))
			serviceCollection.AddTransientByImplements(NewRedisClient, new(redis.IClient))
			serviceCollection.AddTransientByImplements(NewRedisHealthIndicator, new(health.Indicator))
		})
}
