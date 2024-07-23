package abstractions

import "github.com/liangboceo/dependencyinjection"

type IStartup interface {
	ConfigureServices(collection *dependencyinjection.ServiceCollection)
}
