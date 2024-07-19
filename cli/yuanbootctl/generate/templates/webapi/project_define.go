package webapi

import "github.com/liangboceo/yuanboot/cli/yuanbootctl/generate/projects"

var Project = projects.NewEmptyProject("webapi", "Web API Application").With(func(root *projects.ProjectItem) {
	root.AddFileWithContent("main.go", Main_tel)
	root.AddFileWithContent("go.mod", Mod_tel)
})
