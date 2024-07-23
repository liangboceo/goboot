package console

const ProjectItem_go_mod = `
module {{.ModelName}}

go 1.16

require (
	github.com/liangboceo/dependencyinjection v1.0.0
	github.com/liangboceo/yuanboot {{.Version}}
)
`
