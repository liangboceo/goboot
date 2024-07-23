package scheduler

import (
	"github.com/liangboceo/dependencyinjection"
	"github.com/liangboceo/yuanboot/abstractions"
	"github.com/liangboceo/yuanboot/console"
)

func NewXxlJobBuilder(config abstractions.IConfiguration) *abstractions.HostBuilder {
	return console.NewHostBuilder().
		UseConfiguration(config).
		ConfigureServices(func(collection *dependencyinjection.ServiceCollection) {
			UseXxlJob(collection)
		})
}
