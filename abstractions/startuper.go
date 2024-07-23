package abstractions

type IStartup interface {
	ConfigureServices(collection *dependencyinjection.ServiceCollection)
}
