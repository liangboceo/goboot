package main

import (
	"github.com/liangboceo/yuanboot/abstractions/configuration"
	"github.com/liangboceo/yuanboot/pkg/scheduler"
	"github.com/yoyofxteam/dependencyinjection"
)

func main() {
	// -f ./conf/test_conf.yml 指定配置文件 , 默认读取 config_{profile}.yml , -profile [dev,test,prod]
	config := configuration.YAML("config")

	scheduler.NewXxlJobBuilder(config).
		ConfigureServices(func(collection *dependencyinjection.ServiceCollection) {
			scheduler.AddJobs(collection, NewDemoJob)
		}).
		Build().Run()
}
