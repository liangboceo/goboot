package hosting

import (
	"github.com/liangboceo/yuanboot/abstractions"
	"github.com/yoyofxteam/dependencyinjection"
)

func AddHostService(collection *dependencyinjection.ServiceCollection, serviceCtor interface{}) {
	collection.AddSingletonByImplements(serviceCtor, new(abstractions.IHostService))
}
