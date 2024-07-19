package configuration

import "github.com/liangboceo/yuanboot/abstractions"

type OptionsSnapshot[T any] struct {
	config      abstractions.IConfiguration
	sectionName string
	value       T
}

func (options OptionsSnapshot[T]) CurrentValue() T {
	var configObject T
	options.config.GetConfigObject(options.sectionName, &configObject)
	return configObject
}
