package hosting

import (
	"github.com/liangboceo/dependencyinjection"
	"github.com/liangboceo/yuanboot/abstractions"
)

func AddHostService(collection *dependencyinjection.ServiceCollection, serviceCtor interface{}) {
	collection.AddSingletonByImplements(serviceCtor, new(abstractions.IHostService))
}
