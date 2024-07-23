package mvc

import "github.com/liangboceo/dependencyinjection"

func ActivateController(serviceProvider dependencyinjection.IServiceProvider, controllerName string) (IController, error) {
	var controller IController
	err := serviceProvider.GetServiceByName(&controller, controllerName)
	return controller, err
}
