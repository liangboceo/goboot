package configuration

import "github.com/liangboceo/yuanboot/abstractions"

func LocalConfig(configPath string) *abstractions.Configuration {
	return abstractions.NewConfigurationBuilder().AddEnvironment().AddYamlFile(configPath).Build()
}
