package main

import (
	"github.com/liangboceo/yuanboot/abstractions"
	"github.com/liangboceo/yuanboot/web"
	"github.com/liangboceo/yuanboot/web/endpoints"
	"github.com/liangboceo/yuanboot/web/router"
)

func main() {
	configuration := abstractions.NewConfigurationBuilder().AddYamlFile("config").Build()
	// --profile=prod , to read , config.yml
	web.NewWebHostBuilder().
		UseConfiguration(configuration).
		Configure(func(app *web.ApplicationBuilder) {
			app.UseEndpoints(func(router router.IRouterBuilder) {
				router.POST("/alert", PostAlert)
				endpoints.UsePrometheus(router)
			})
		}).Build().Run()
}
