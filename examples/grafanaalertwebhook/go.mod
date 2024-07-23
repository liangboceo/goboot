module grafanaalertwebhook

go 1.21

toolchain go1.21.5

require (
	github.com/liangboceo/yuanboot v0.0.0
	gopkg.in/go-playground/assert.v1 v1.2.1
)

require gopkg.in/yaml.v3 v3.0.1 // indirect

replace github.com/liangboceo/yuanboot => ../../
