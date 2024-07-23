package grpc

import "github.com/liangboceo/dependencyinjection"

func AddService(collection *dependencyinjection.ServiceCollection, serviceCtor interface{}) {
	collection.AddSingleton(serviceCtor)
}
