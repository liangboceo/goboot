module websockethub

go 1.18

require (
	github.com/fasthttp/websocket v1.5.7
	github.com/liangboceo/yuanboot v0.0.0
	github.com/liangboceo/dependencyinjection v1.0.0
)

require gopkg.in/yaml.v3 v3.0.1 // indirect

replace github.com/liangboceo/yuanboot => ../../
